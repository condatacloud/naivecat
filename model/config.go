package model

import (
	"naivecat/tools"
)

type Config struct {
	AutoLink    bool   `json:"autoLink"`
	DefaultLink int    `json:"defaultLink"`
	Theme       string `json:"theme"` // Light | Dark
	Links       Links  `json:"links"`
	Host        string `json:"host"`
	Socks       struct {
		Port string `json:"port"`
	} `json:"socks"`
	Http struct {
		Enable bool   `json:"enable"`
		Port   string `json:"port"`
	} `json:"http"`
	EnableLog bool `json:"enableLog"`
}

func (c *Config) Deserialize(filePath string) error {
	return tools.Deserialize(c, filePath)
}

func (c *Config) Serialize(filePath string) error {
	return tools.Serialize(c, filePath)
}

func (c *Config) LoadConfig() {
	home, err := tools.HomeDir()
	if err != nil {
		panic(err)
	}
	folder := home + "/.naivecat"
	if !tools.File.Exists(folder) {
		tools.File.Mkdir(folder)
	}

	filePath := folder + "/config.json"
	if !tools.File.Exists(filePath) {
		c.AutoLink = false
		c.Theme = THEME_DARK
		if err := tools.Serialize(c, filePath); err != nil {
			panic(err)
		}
	}
	if err := tools.Deserialize(c, filePath); err != nil {
		panic(err)
	}

	c.checkConfig()
}

func (c *Config) Update() {
	home, err := tools.HomeDir()
	if err != nil {
		panic(err)
	}
	folder := home + "/.naivecat"
	if !tools.File.Exists(folder) {
		tools.File.Mkdir(folder)
	}

	// id重新排序
	for i, v := range c.Links {
		v.ID = i
	}

	c.checkConfig()

	filePath := folder + "/config.json"
	if err := tools.Serialize(c, filePath); err != nil {
		panic(err)
	}
}

func (c *Config) checkConfig() {
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Socks.Port == "" {
		c.Socks.Port = "1080"
	}

	if c.Http.Enable {
		if c.Http.Port == "" {
			c.Http.Port = "8000"
		}
	}
	if len(c.Links) == 0 {
		defaultLink := &Link{}
		defaultLink.NewDefaultLink()
		c.Links = append(c.Links, defaultLink)
	}
}
