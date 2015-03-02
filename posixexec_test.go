package posixexec

import (
	"os/exec"
	"testing"
)

func TestZero(t *testing.T) {
	cmd := exec.Command("/bin/sh", "-c", "exit 0")
	exitStatus, err := Run(cmd)

	if exitStatus != 0 {
		t.Errorf("exit status should be 0")
	}

	if err != nil {
		t.Errorf("err should be nil")
	}
}

func TestNonZero(t *testing.T) {
	cmd := exec.Command("/bin/sh", "-c", "exit 1")
	exitStatus, err := Run(cmd)

	if exitStatus != 1 {
		t.Errorf("exit status should be 1")
	}

	if err != nil {
		t.Errorf("err should be nil")
	}
}
