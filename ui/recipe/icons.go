package recipe

import (
	"naivecat/resource"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const (
	IconNameStart   fyne.ThemeIconName = "start"
	IconNameStop    fyne.ThemeIconName = "stop"
	IconNameNetwork fyne.ThemeIconName = "network"
	IconNameShared  fyne.ThemeIconName = "shared"
	IconNameLink    fyne.ThemeIconName = "link"
)

var (
	Icons = map[fyne.ThemeIconName]fyne.Resource{
		IconNameStart:   theme.NewThemedResource(fyne.NewStaticResource("start", resource.IconStart)),
		IconNameStop:    theme.NewThemedResource(fyne.NewStaticResource("stop", resource.IconStop)),
		IconNameNetwork: theme.NewThemedResource(fyne.NewStaticResource("network", resource.IconNetwork)),
		IconNameShared:  theme.NewThemedResource(fyne.NewStaticResource("shared", resource.IconShared)),
		IconNameLink:    theme.NewThemedResource(fyne.NewStaticResource("link", resource.IconLink)),
	}
)
