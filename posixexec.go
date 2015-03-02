package posixexec

import (
	"os/exec"
	"syscall"
)

// Run runs *exec.Cmd and returns exit-status
func Run(cmd *exec.Cmd) (int, error) {
	exitStatus := -1
	var exiterr error

	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if s, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				exitStatus = s.ExitStatus()
			}
		}
	} else {
		exitStatus = 0
	}

	return exitStatus, exiterr
}
