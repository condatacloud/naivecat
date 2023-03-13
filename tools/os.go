package tools

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

// HomeDir 获得用户的主目录
// https://www.w3cschool.cn/cuhkj/cuhkj-azch266d.html
func HomeDir() (string, error) {
	user, err := user.Current()
	if err == nil {
		if runtime.GOOS == "linux" {
			// 排除sudo的影响
			var ruser = os.Getenv("SUDO_USER")
			if ruser == "" {
				return "/home/" + user.Username, nil
			}
			return "/home/" + ruser, nil
		}
		return user.HomeDir, nil
	}
	return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
}

func Exec(wd string, name string, subname string, args ...string) (string, error) {
	args = append([]string{subname}, args...)

	cmd := exec.Command(name, args...)
	if wd != "" {
		cmd.Dir = wd
	}
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		cmdInfo := name + " " + strings.Join(args, " ")
		return string(bytes), fmt.Errorf("failed to execute %s command: %s", cmdInfo, err.Error())
	}

	return string(bytes), nil
}
