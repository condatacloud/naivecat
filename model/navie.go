package model

type NaiveConfig struct {
	Listen string `json:"listen"`
	Proxy  string `json:"proxy"`
	Log    string `json:"log"`
}
