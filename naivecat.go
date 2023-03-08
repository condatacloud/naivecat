package main

import (
	"flag"
	"fmt"
	"naivecat/model"
	"naivecat/resource"
	"naivecat/service"
	"naivecat/ui"
	"naivecat/ui/recipe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

// https://github.com/Nikeweke/pomadorik
// https://blog.meetwhy.com/blog/some-experience-in-the-use-of-fyne.html
// https://www.cnblogs.com/wustjq/p/16525001.html

var winWidth float32 = 980
var winHeight float32 = 720
var appIcon *fyne.StaticResource
var wnd fyne.Window
var appName = "Naivecat"

func init() {
	appIcon = fyne.NewStaticResource("app_icon", resource.AppIconPngBytes)
}

func cmdFlags() bool {
	var command string
	flag.StringVar(&command, "command", "", "BuildTime")
	//解析命令行参数
	flag.Parse()

	switch command {
	case "BuildTime":
		fmt.Println(model.BuildTime)
		return true
	}
	return false
}

func setup() bool {
	// 检查是否安装过，没有的话执行安装步骤
	if !service.SetupService.Done() {
		service.SetupService.Install()
	} else {
		if service.SetupService.NeedUpgrade() {
			ui.NewUpgradeUI()
			return true
		}
	}
	return false
}

func main() {
	// 命令解析命令，如果返回true就代表退出程序
	if cmdFlags() {
		return
	}
	// 安装的一些检查，如果返回true就代表退出程序
	if setup() {
		return
	}
	// 加载配置文件
	ui.GConfig.LoadConfig()
	// 初始化Naive的执行环境，拷贝执行文件到系统内
	service.NaiveService.InitEnv()
	// 创建一个app
	app := app.New()
	ui.App = app
	// 设置主题、字体
	app.Settings().SetTheme(&recipe.DarkTheme{})
	// 创建一个窗口
	wnd = app.NewWindow(appName)
	ui.Wnd = wnd
	// 修改窗口大小
	wnd.Resize(fyne.NewSize(winWidth, winHeight))
	// 设置图标
	wnd.SetIcon(appIcon)
	// 设置系统托盘
	if desk, ok := app.(desktop.App); ok {
		setupSystray(desk)
	}
	// 主窗口关闭时设置自动隐藏
	wnd.SetCloseIntercept(func() { wnd.Hide() })
	// 构建UI
	wnd.SetContent(ui.NewUI())
	// 更新UI的一些数据，否则页面显示时数据是空的
	ui.UpdateUIData()
	// 启动
	wnd.Show()
	// UI显示之后需要做的一些事情
	ui.AfterUIShowExec()
	// 堵塞主线程
	app.Run()
}

// 系统托盘
func setupSystray(desk desktop.App) {
	desk.SetSystemTrayIcon(appIcon)

	menu := fyne.NewMenu(
		appName,
		fyne.NewMenuItem("Show", wnd.Show),
	)

	desk.SetSystemTrayMenu(menu)
}
