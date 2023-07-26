//go:build darwin

package cmd

import (
	"os/exec"
)

// 普通的cmd 客户端
func setAttr(cmd *exec.Cmd) {
}
func killProcess(cmd *exec.Cmd) {
	cmd.Process.Kill()
}
