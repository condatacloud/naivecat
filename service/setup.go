package service

import (
	"fmt"
	"log"
	"naivecat/model"
	"naivecat/resource"
	"naivecat/tools"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gotopkg/mslnk/pkg/mslnk"
)

type ISetup interface {
	// 判断是否已经安装
	Done() bool
	// 安装
	Install(callback func(float32))
	// 是否需要升级
	NeedUpgrade() bool
	// 升级
	Upgrade(callback func(float32))
}

type setupService struct {
}

var SetupService ISetup = &setupService{}

func (s *setupService) Done() bool {
	home, err := tools.HomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if runtime.GOOS == "linux" {
		lnk := home + "/.local/share/applications/naivecat.desktop"
		if !tools.File.Exists(lnk) {
			return false
		}

		naiveFilePath := home + "/.naivecat/naive"
		if !tools.File.Exists(naiveFilePath) {
			return false
		}

	} else if runtime.GOOS == "windows" {
		lnk := home + "/Desktop/Naivecat.lnk"
		if !tools.File.Exists(lnk) {
			return false
		}

		naiveFilePath := home + "/.naivecat/naive.exe"
		if !tools.File.Exists(naiveFilePath) {
			return false
		}
	} else {
		log.Fatal("不支持该系统", runtime.GOOS)
	}
	return true
}

func (s *setupService) Install(callback func(float32)) {
	home, err := tools.HomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if runtime.GOOS == "linux" {
		// 拷贝可执行文件
		data := s.readExecFile()
		callback(0.2)
		execFileDir := home + "/.naivecat"
		if !tools.File.Exists(execFileDir) {
			tools.File.Mkdir(execFileDir)
		}
		callback(0.3)
		execFilePath := execFileDir + "/naivecat"
		if err := os.WriteFile(execFilePath, data, 0755); err != nil {
			log.Fatal(err)
		}
		callback(0.5)
		// 拷贝图标
		iconFilePath := execFileDir + "/naivecat.png"
		tools.File.WriteBin(resource.AppIconPngBytes, iconFilePath)
		callback(0.65)

		// 创建启动图标
		lnk := home + "/.local/share/applications/naivecat.desktop"
		lnkContent := "[Desktop Entry]\n"
		lnkContent += "Name=Naivecat\n"
		lnkContent += "Comment=Network proxy\n"
		lnkContent += fmt.Sprintf("Exec=%s\n", execFilePath)
		lnkContent += "Terminal=false\n"
		lnkContent += fmt.Sprintf("Icon=%s\n", iconFilePath)
		lnkContent += "Type=Application\n"
		lnkContent += "Categories=Network;\n"
		lnkContent += "SingleMainWindow=true\n"
		if err := os.WriteFile(lnk, []byte(lnkContent), 0755); err != nil {
			log.Fatal(err)
		}
		callback(0.7)

		naiveFilePath := home + "/.naivecat/naive"
		if err := tools.File.WriteBin(NaiveBytes, naiveFilePath); err != nil {
			panic(err)
		}
		callback(0.9)
		// 增加执行权限
		cmd := exec.Command("chmod", "+x", naiveFilePath)
		_, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}
		callback(1)
	} else if runtime.GOOS == "windows" {
		// 拷贝可执行文件
		data := s.readExecFile()
		callback(0.2)
		execFileDir := home + "/.naivecat"
		if !tools.File.Exists(execFileDir) {
			tools.File.Mkdir(execFileDir)
		}
		callback(0.3)
		execFilePath := execFileDir + "/naivecat.exe"
		if err := os.WriteFile(execFilePath, data, 0755); err != nil {
			log.Fatal(err)
		}
		callback(0.5)
		// 拷贝图标
		iconFilePath := execFileDir + "/naivecat.png"
		tools.File.WriteBin(resource.AppIconPngBytes, iconFilePath)
		callback(0.65)
		// 创建快捷方式
		lnk := home + "/Desktop/Naivecat.lnk"
		if err := mslnk.LinkFile(execFilePath, lnk); err != nil {
			log.Fatal(err)
		}
		callback(0.8)

		naiveFilePath := home + "/.naivecat/naive.exe"
		if err := tools.File.WriteBin(NaiveBytes, naiveFilePath); err != nil {
			panic(err)
		}
		callback(1)
	} else {
		log.Fatal("不支持该系统", runtime.GOOS)
	}
}

func (s *setupService) NeedUpgrade() bool {
	home, err := tools.HomeDir()
	if err != nil {
		log.Fatal(err)
	}
	buildTime, _ := time.Parse("2006-01-02 15:04:05", model.BuildTime)
	execFilePath := ""
	if runtime.GOOS == "linux" {
		execFilePath = home + "/.naivecat/naivecat"
	} else if runtime.GOOS == "windows" {
		execFilePath = home + "/.naivecat/naivecat.exe"
	} else {
		log.Fatal("不支持该系统", runtime.GOOS)
	}

	result, err := tools.Exec("", execFilePath, "-command", "BuildTime")
	if err != nil {
		log.Fatal(err)
	}
	result = tools.Strings.TrimN(result)
	result = result[len(result)-19:]
	oldBuildTime, _ := time.Parse("2006-01-02 15:04:05", result)
	return buildTime.After(oldBuildTime)
}

func (s *setupService) Upgrade(callback func(float32)) {
	s.Install(callback)
}

func (s *setupService) readExecFile() []byte {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	data, err := os.ReadFile(ex)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
