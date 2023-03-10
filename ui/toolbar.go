package ui

import (
	"fmt"
	"naivecat/model"
	"naivecat/service"
	"naivecat/tools"
	"naivecat/ui/controls"
	"naivecat/ui/recipe"
	"strconv"
	"time"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ToolbarUI struct {
	startBtn   *widget.ToolbarAction
	topToolbar *widget.Toolbar
	activeLink *controls.ToolbarIconLbl
}

var toolbarUI = &ToolbarUI{}

func (u *ToolbarUI) Update() {}

func (u *ToolbarUI) NewFooterUI() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			linkNewUI.NewUI()
		}),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			linkEditUI.NewUI()
		}),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			linkDeleteUI.NewUI()
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			linkImportUI.NewUI()
		}),
	)
	return toolbar
}

func (u *ToolbarUI) NewTopUI() *widget.Toolbar {
	u.startBtn = widget.NewToolbarAction(recipe.Icons[recipe.IconNameStart], u.start)

	u.activeLink = controls.NewToolbarIconLbl(recipe.Icons[recipe.IconNameLink], "")

	u.topToolbar = widget.NewToolbar(
		u.activeLink,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(recipe.Icons[recipe.IconNameNetwork], u.ping),
		u.startBtn,
		widget.NewToolbarAction(recipe.Icons[recipe.IconNameShared], func() {
			linkPannelUI.qrcode2Clipboard()
		}),
	)

	return u.topToolbar
}

func (u *ToolbarUI) ping() {
	if len(GConfig.Links) == 0 {
		return
	}

	go func() {
		link := GConfig.Links[GConfig.DefaultLink]
		tm, err := service.PingService.Ping(link.Host)
		if err != nil {
			link.Ping = err.Error()
		} else {
			link.Ping = strconv.FormatInt(tm, 10) + "ms"
		}
		GConfig.Update()
		linkTableUI.Update()
	}()
}

func (u *ToolbarUI) start() {
	if service.NaiveService.IsRunning() {
		u.stop()
		u.setStartIcon()
		return
	}

	if tools.Net.ScanPort("tcp", GConfig.Host, GConfig.Socks.Port) {
		controls.Msgbox("错误", fmt.Sprintf("socks端口%s已经被监听", GConfig.Socks.Port), Wnd)
		return
	}

	if GConfig.Http.Enable {
		if tools.Net.ScanPort("tcp", GConfig.Host, GConfig.Http.Port) {
			controls.Msgbox("错误", fmt.Sprintf("http端口%s已经被监听", GConfig.Http.Port), Wnd)
			return
		}
	}

	consoleUI.clear()
	link := GConfig.Links[GConfig.DefaultLink]
	naiveConfig := link.ToNaiveConfig(GConfig.Host, GConfig.Socks.Port)
	// 先更新配置，再执行
	if err := service.NaiveService.UpdateConfig(naiveConfig); err != nil {
		controls.Msgbox("错误", fmt.Sprintf("更新配置文件发生错误\n%s", err.Error()), Wnd)
		return
	}

	if GConfig.EnableLog {
		callback := func(line string) {
			consoleUI.append(line)
		}
		service.NaiveService.SetLogCallback(callback)
	}

	if GConfig.Http.Enable {
		proxyHttp := &service.ForwardInfo{
			SrcIP:   GConfig.Host,
			SrcPort: GConfig.Socks.Port,
			DstIP:   GConfig.Host,
			DstPort: GConfig.Http.Port,
		}
		go func() {
			if err := service.ProxyService.Start(proxyHttp); err != nil {
				controls.Msgbox("错误", err.Error(), Wnd)
				model.Log.Error(err.Error())
				u.setStartIcon()
			}
		}()
	}

	go func() {
		if err := service.NaiveService.Start(); err != nil {
			controls.Msgbox("错误", fmt.Sprintf("运行该链接发生错误\n%s", err.Error()), Wnd)
			model.Log.Error(err.Error())
			u.setStartIcon()
		}
	}()

	u.activeLink.SetText(link.Name)
	time.Sleep(600 * time.Millisecond)
	u.setStopIcon()
}

func (u *ToolbarUI) stop() {
	if service.NaiveService.IsRunning() {
		service.NaiveService.Close()
		consoleUI.clear()
		u.activeLink.SetText("")
	}

	if service.ProxyService.IsRunning() {
		if err := service.ProxyService.Close(); err != nil {
			controls.Msgbox("错误", err.Error(), Wnd)
		}
	}
}

func (u *ToolbarUI) setStartIcon() {
	u.startBtn.SetIcon(recipe.Icons[recipe.IconNameStart])
	u.topToolbar.Refresh()
}

func (u *ToolbarUI) setStopIcon() {
	u.startBtn.SetIcon(recipe.Icons[recipe.IconNameStop])
	u.topToolbar.Refresh()
}
