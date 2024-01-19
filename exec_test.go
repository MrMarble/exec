package exec_test

import (
	"testing"
	"time"

	"github.com/mrmarble/exec"
)

func TestExec(t *testing.T) {
	t.Parallel()
	t.Run("Output", func(t *testing.T) {
		out, err := exec.Command("/bin/sh", "-c", "echo 'stdout'").Output()
		if err != nil {
			t.Error(err)
		}
		if string(out) != "stdout\n" {
			t.Errorf("Expected 'stdout\\n', got '%s'", out)
		}
	})

	t.Run("Long running", func(t *testing.T) {
		start := time.Now()
		out, err := exec.Command("./testdata/long_running.sh").Output()
		if err != nil {
			t.Error(err)
		}
		elapsed := time.Since(start)
		if string(out) != "outputA\noutputB\n" {
			t.Errorf("Expected 'outputA\\noutputB\\n', got '%s'", out)
		}
		if elapsed > 1*time.Second {
			t.Errorf("Expected execution time < 1s, got %s", elapsed.String())
		}
	})

	t.Run("CombinedOutput", func(t *testing.T) {
		out, err := exec.Command("/bin/sh", "-c", "echo 'stdout'; echo 'stderr' 1>&2").CombinedOutput()
		if err != nil {
			t.Error(err)
		}
		if string(out) != "stdout\nstderr\n" {
			t.Errorf("Expected 'stdout\\nstderr\\n', got '%s'", out)
		}
	})
}
