package model

var (
	GitHash   string
	BuildTime string
	GoVersion string
	GitBranch string
)

const (
	THEME_LIGHT = "Light"
	THEME_DARK  = "Dark"
)

var (
	THEMES = []string{THEME_LIGHT, THEME_DARK}
)
