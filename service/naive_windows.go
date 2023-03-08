package service

import (
	_ "embed"
	"os/exec"
	"syscall"
)

//go:embed naive/win/naive.exe
var NaiveBytes []byte

func hideCommandWindow(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
