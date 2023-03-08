package service

import (
	"fmt"
	"log"
	"naivecat/model"
	"naivecat/resource"
	"naivecat/tools"
	"os"
	"runtime"
	"time"

	"github.com/gotopkg/mslnk/pkg/mslnk"
)

type ISetup interface {
	// 判断是否已经安装
	Done() bool
	// 安装
	Install()
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
	done := false
	if runtime.GOOS == "linux" {
		lnk := home + "/.local/share/applications/naivecat.desktop"
		if tools.File.Exists(lnk) {
			done = true
		}
	} else if runtime.GOOS == "windows" {
		lnk := home + "/Desktop/Naivecat.lnk"
		if tools.File.Exists(lnk) {
			done = true
		}
	} else {
		log.Fatal("不支持该系统", runtime.GOOS)
	}
	return done
}

func (s *setupService) Install() {
	home, err := tools.HomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if runtime.GOOS == "linux" {
		// 拷贝可执行文件
		data := s.readExecFile()
		execFileDir := home + "/.naivecat"
		if !tools.File.Exists(execFileDir) {
			tools.File.Mkdir(execFileDir)
		}
		execFilePath := execFileDir + "/naivecat"
		if err := os.WriteFile(execFilePath, data, 0755); err != nil {
			log.Fatal(err)
		}
		// 拷贝图标
		iconFilePath := execFileDir + "/naivecat.png"
		tools.File.WriteBin(resource.AppIconPngBytes, iconFilePath)

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
	} else if runtime.GOOS == "windows" {
		// 拷贝可执行文件
		data := s.readExecFile()
		execFileDir := home + "/.naivecat"
		if !tools.File.Exists(execFileDir) {
			tools.File.Mkdir(execFileDir)
		}
		execFilePath := execFileDir + "/naivecat.exe"
		if err := os.WriteFile(execFilePath, data, 0755); err != nil {
			log.Fatal(err)
		}
		// 拷贝图标
		iconFilePath := execFileDir + "/naivecat.png"
		tools.File.WriteBin(resource.AppIconPngBytes, iconFilePath)
		// 创建快捷方式
		lnk := home + "/Desktop/Naivecat.lnk"
		if err := mslnk.LinkFile(execFilePath, lnk); err != nil {
			log.Fatal(err)
		}
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
	callback(0.3)
	s.Install()
	callback(1)
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
