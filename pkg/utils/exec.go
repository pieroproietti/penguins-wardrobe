package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func ensureRootPath() {
	if os.Geteuid() != 0 {
		return
	}

	const rootPath = "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

	if os.Getenv("PATH") != rootPath {
		os.Setenv("PATH", rootPath)
	}
}

// Exec esegue un comando sh e mostra l'output in tempo reale sul terminale.
func Exec(command string) error {
	ensureRootPath()

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ExecQuiet esegue un comando senza mostrare nulla
func ExecQuiet(command string) error {
	ensureRootPath()

	cmd := exec.Command("sh", "-c", command)
	return cmd.Run()
}

// ExecLog esegue un comando sh e redirige stdout e stderr nel file di log specificato.
func ExecLog(command string, logFilePath string) error {
	return ExecLogMonitor(command, logFilePath, nil)
}

// ExecLogMonitor esegue un comando sh, scrive tutto nel logfile e notifica onLine riga per riga
func ExecLogMonitor(command string, logFilePath string, onLine func(line string)) error {
	ensureRootPath()

	if err := os.MkdirAll(filepath.Dir(logFilePath), 0755); err != nil {
		// In caso di errore directory, prosegue
	}

	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return ExecQuiet(command)
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	f.WriteString(fmt.Sprintf("[%s] EXEC: %s\n", timestamp, command))

	cmd := exec.Command("sh", "-c", command)

	r, w, err := os.Pipe()
	if err != nil {
		cmd.Stdout = f
		cmd.Stderr = f
		return cmd.Run()
	}
	cmd.Stdout = w
	cmd.Stderr = w

	if err := cmd.Start(); err != nil {
		w.Close()
		r.Close()
		return err
	}
	w.Close()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		f.WriteString(line + "\n")
		if onLine != nil {
			onLine(line)
		}
	}
	r.Close()

	return cmd.Wait()
}

// ExecCapture esegue un comando e restituisce l'output come stringa
func ExecCapture(command string) (string, error) {
	ensureRootPath()

	var out bytes.Buffer
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = &out
	return out.String(), cmd.Run()
}

// ExecCaptureCombined esegue un comando e restituisce sia stdout che stderr integrati come stringa
func ExecCaptureCombined(command string) (string, error) {
	ensureRootPath()

	var out bytes.Buffer
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}
