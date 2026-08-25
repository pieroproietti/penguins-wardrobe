package utils

import (
	"testing"
	"time"
)

func TestSpinnerLifeCycle(t *testing.T) {
	sp := NewSpinner("Testing spinner...")
	if sp == nil {
		t.Fatal("expected non-nil spinner")
	}

	sp.Start()
	time.Sleep(50 * time.Millisecond)
	sp.UpdateText("Updated spinner text...")
	time.Sleep(50 * time.Millisecond)
	sp.Success("Spinner completed successfully")
}

func TestSpinnerWarnFailInfo(t *testing.T) {
	spWarn := NewSpinner("Testing warn...")
	spWarn.Start()
	spWarn.Warn("Warning notice")

	spFail := NewSpinner("Testing fail...")
	spFail.Start()
	spFail.Fail("Failure notice")

	spInfo := NewSpinner("Testing info...")
	spInfo.Start()
	spInfo.Info("Info notice")
}
