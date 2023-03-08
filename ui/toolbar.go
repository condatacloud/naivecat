package ui

import (
	"fmt"
	"naivecat/resource"
	"naivecat/service"
	"naivecat/tools"
	"naivecat/ui/controls"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ToolbarUI struct {
	startBtn     *widget.ToolbarAction
	startIconRes *fyne.StaticResource
	stopIconRes  *fyne.StaticResource
	topToolbar   *widget.Toolbar
}

var toolbarUI = &ToolbarUI{
	startIconRes: fyne.NewStaticResource("start", resource.IconStart),
	stopIconRes:  fyne.NewStaticResource("stop", resource.IconStop),
}

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
	u.startBtn = widget.NewToolbarAction(u.startIconRes, u.start)

	u.topToolbar = widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(fyne.NewStaticResource("internet", resource.IconInternet), u.ping),
		u.startBtn,
		widget.NewToolbarAction(fyne.NewStaticResource("shared", resource.IconShare), func() {
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
		u.startBtn.SetIcon(u.startIconRes)
		u.topToolbar.Refresh()
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
			}
		}()
	}

	go func() {
		if err := service.NaiveService.Start(); err != nil {
			controls.Msgbox("错误", fmt.Sprintf("运行该链接发生错误\n%s", err.Error()), Wnd)
		}
	}()

	time.Sleep(600 * time.Millisecond)
	u.startBtn.SetIcon(u.stopIconRes)
	u.topToolbar.Refresh()
}

func (u *ToolbarUI) stop() {
	if service.NaiveService.IsRunning() {
		service.NaiveService.Close()
		consoleUI.clear()
	}

	if service.ProxyService.IsRunning() {
		if err := service.ProxyService.Close(); err != nil {
			controls.Msgbox("错误", err.Error(), Wnd)
		}
	}
}
