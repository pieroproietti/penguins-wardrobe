// Copyright 2026 Piero Proietti <piero.proietti@gmail.com>.
// All rights reserved.

package utils

import (
	"fmt"
	"strings"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// GetTerminalSize returns terminal rows and columns, or (24, 80) as default fallback.
func GetTerminalSize() (int, int) {
	ws := &winsize{}
	ret, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))
	if int(ret) == 0 && ws.Row > 0 && ws.Col > 0 {
		return int(ws.Row), int(ws.Col)
	}
	return 24, 80
}

type SplitScreen struct {
	mu           sync.Mutex
	active       bool
	totalRows    int
	totalCols    int
	headerRows   int
	maxSteps     int
	topHeight    int // row number of current action / spinner
	scrollStart  int // first row of bottom scrolling region
	completed    []string
	currentAction string
	startTime    time.Time
	stopChan     chan struct{}
}

var globalSplitScreen *SplitScreen
var splitMu sync.Mutex

// GetSplitScreen returns the active SplitScreen instance if any
func GetSplitScreen() *SplitScreen {
	splitMu.Lock()
	defer splitMu.Unlock()
	return globalSplitScreen
}

// StartSplitScreen initializes the horizontal split screen on terminal
func StartSplitScreen(icon, title, subtitle string) *SplitScreen {
	splitMu.Lock()
	defer splitMu.Unlock()

	if !isTerminal() {
		return nil
	}

	rows, cols := GetTerminalSize()
	if rows < 14 {
		// Terminal too small for meaningful split screen
		return nil
	}

	headerRows := 4
	if subtitle == "" {
		headerRows = 3
	}

	maxSteps := 4
	if rows >= 30 {
		maxSteps = 8
	} else if rows >= 24 {
		maxSteps = 5
	}

	topHeight := headerRows + maxSteps + 1
	sepRow := topHeight + 1
	scrollStart := sepRow + 1

	if scrollStart >= rows-3 {
		maxSteps = 3
		topHeight = headerRows + maxSteps + 1
		sepRow = topHeight + 1
		scrollStart = sepRow + 1
	}

	ss := &SplitScreen{
		active:       true,
		totalRows:    rows,
		totalCols:    cols,
		headerRows:   headerRows,
		maxSteps:     maxSteps,
		topHeight:    topHeight,
		scrollStart:  scrollStart,
		completed:    make([]string, 0),
		currentAction: "",
		startTime:    time.Now(),
		stopChan:     make(chan struct{}),
	}

	// Clear entire terminal and move to (1,1)
	fmt.Print("\033[2J\033[1;1H")

	// Draw top header
	divWidth := cols
	if divWidth > 72 {
		divWidth = 72
	}
	divider := strings.Repeat("=", divWidth)
	
	fmt.Printf("%s%s%s\n", colorize(ColorCyan), divider, colorize(ColorReset))
	fmt.Printf("  %s %s%s%s\n", icon, colorize(ColorBold+ColorWhite), title, colorize(ColorReset))
	if subtitle != "" {
		fmt.Printf("  %s%s%s\n", colorize(ColorDim), subtitle, colorize(ColorReset))
	}
	fmt.Printf("%s%s%s\n", colorize(ColorCyan), divider, colorize(ColorReset))

	// Draw separator line at sepRow
	ss.drawSeparator(sepRow, cols)

	// Set DECSTBM scrolling region for the bottom area
	fmt.Printf("\033[%d;%dr", scrollStart, rows)

	// Position cursor in the bottom scrolling region
	fmt.Printf("\033[%d;1H", scrollStart)

	// Start background spinner updater for the top status line
	go ss.spinnerLoop()

	globalSplitScreen = ss
	return ss
}

func (ss *SplitScreen) drawSeparator(row, cols int) {
	tag := " CONSOLE INTERATTIVA "
	totalLen := cols
	if totalLen > 72 {
		totalLen = 72
	}
	leftLen := 4
	rightLen := totalLen - leftLen - len(tag)
	if rightLen < 4 {
		rightLen = 4
	}

	sepLine := fmt.Sprintf("%s%s%s%s%s%s%s",
		colorize(ColorCyan),
		strings.Repeat("─", leftLen),
		colorize(ColorBold+ColorWhite),
		tag,
		colorize(ColorReset+ColorCyan),
		strings.Repeat("─", rightLen),
		colorize(ColorReset),
	)

	// Save cursor, draw separator at row, restore cursor
	fmt.Printf("\0337\033[%d;1H\033[2K%s\0338", row, sepLine)
}

func (ss *SplitScreen) spinnerLoop() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	idx := 0

	for {
		select {
		case <-ss.stopChan:
			return
		case <-ticker.C:
			ss.mu.Lock()
			if !ss.active {
				ss.mu.Unlock()
				return
			}
			action := ss.currentAction
			if action == "" {
				ss.mu.Unlock()
				continue
			}

			frame := asciiSpinnerFrames[idx%len(asciiSpinnerFrames)]
			idx++
			elapsed := time.Since(ss.startTime)
			row := ss.topHeight
			ss.mu.Unlock()

			mins := int(elapsed.Minutes())
			secs := int(elapsed.Seconds()) % 60
			timeStr := fmt.Sprintf("[%02d:%02d]", mins, secs)

			line := fmt.Sprintf("  %s[%s%s%s]%s %s %s%s%s",
				colorize(ColorCyan),
				colorize(ColorBold+ColorWhite), frame, colorize(ColorReset+ColorCyan),
				colorize(ColorReset),
				action,
				colorize(ColorDim), timeStr, colorize(ColorReset))

			// DEC Save cursor, move to spinner row, clear line, write text, DEC Restore cursor
			fmt.Printf("\0337\033[%d;1H\033[2K%s\0338", row, line)
		}
	}
}

// AddStep appends a completed step to the top pane and refreshes the status view
func (ss *SplitScreen) AddStep(step string) {
	if ss == nil {
		return
	}
	ss.mu.Lock()
	defer ss.mu.Unlock()

	ss.completed = append(ss.completed, step)
	ss.currentAction = ""
	ss.redrawStatusLocked()
}

// SetAction sets the current action description shown on the animated spinner line
func (ss *SplitScreen) SetAction(format string, a ...interface{}) {
	if ss == nil {
		return
	}
	ss.mu.Lock()
	defer ss.mu.Unlock()

	ss.currentAction = fmt.Sprintf(format, a...)
}

func (ss *SplitScreen) redrawStatusLocked() {
	startRow := ss.headerRows + 1
	maxVisible := ss.maxSteps

	// Determine which completed steps to display
	var visible []string
	if len(ss.completed) <= maxVisible {
		visible = ss.completed
	} else {
		visible = ss.completed[len(ss.completed)-maxVisible:]
	}

	// Buffer all redraw operations
	var sb strings.Builder
	sb.WriteString("\0337") // save cursor

	for i := 0; i < maxVisible; i++ {
		row := startRow + i
		sb.WriteString(fmt.Sprintf("\033[%d;1H\033[2K", row))
		if i < len(visible) {
			sb.WriteString("  " + visible[i])
		}
	}

	// Clear current action row
	sb.WriteString(fmt.Sprintf("\033[%d;1H\033[2K", ss.topHeight))
	sb.WriteString("\0338") // restore cursor

	fmt.Print(sb.String())
}

// Close finishes the split screen, restores scrolling margins and moves cursor to bottom
func (ss *SplitScreen) Close() {
	if ss == nil {
		return
	}
	ss.mu.Lock()
	if !ss.active {
		ss.mu.Unlock()
		return
	}
	ss.active = false
	close(ss.stopChan)
	ss.mu.Unlock()

	splitMu.Lock()
	globalSplitScreen = nil
	splitMu.Unlock()

	// Reset DECSTBM scrolling region to full screen
	fmt.Print("\033[r")
	// Move cursor to bottom row and output newline
	fmt.Printf("\033[%d;1H\n", ss.totalRows)
	// Show cursor
	fmt.Print("\033[?25h")
}
