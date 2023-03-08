package service

import (
	"os/exec"

	_ "embed"
)

//go:embed naive/linux/naive
var NaiveBytes []byte

func hideCommandWindow(cmd *exec.Cmd) {

}
