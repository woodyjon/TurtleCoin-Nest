// +build windows

package walletdmanager

import (
	"os/exec"
	"syscall"
)

func hideCmdWindowIfNeeded(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
