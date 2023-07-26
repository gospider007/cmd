//go:build linux

package cmd

import (
	"os/exec"
	"syscall"
)

// 普通的cmd 客户端
func setAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
}
func killProcess(cmd *exec.Cmd) {
	cmd.Process.Kill()
	syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL) // Kill the process and its children
}
