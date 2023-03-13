package model

import (
	"errors"
	"fmt"
	"image"
	"strings"

	"naivecat/tools"

	qrcode "github.com/skip2/go-qrcode"
)

type Link struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
	Username string `json:"username"`
	Password string `json:"password"`
	Padding  bool   `json:"padding"`
	Ping     string `json:"ping"`
}

var LinkProtocols = []string{"https", "quic"}

type Links []*Link

func (l *Links) ToNameList() []string {
	res := []string{}
	for _, v := range *l {
		res = append(res, v.Name)
	}
	return res
}

func (l *Links) ToPingList() []string {
	res := []string{}
	for _, v := range *l {
		res = append(res, v.Ping)
	}
	return res
}

func (l *Link) NewDefaultLink() {
	l.ID = 0
	l.Name = "cat"
	l.Host = "www.google.com"
	l.Port = "443"
	l.Protocol = "https"
	l.Password = "000000"
	l.Username = "google"
	l.Padding = false
	l.Ping = ""
}

func (l *Link) Copy(n *Link) {
	l.ID = n.ID
	l.Name = n.Name
	l.Host = n.Host
	l.Port = n.Port
	l.Protocol = n.Protocol
	l.Username = n.Username
	l.Password = n.Password
	l.Padding = n.Padding
	l.Ping = n.Ping
}

func (l *Link) FromString(txt string) error {
	i1 := strings.Index(txt, "+")
	if i1 == -1 {
		return errors.New("协议错误")
	}
	i2 := strings.Index(txt, ":")
	if i2 == -1 {
		return errors.New("协议错误")
	}
	l.Protocol = txt[i1+1 : i2]

	ntxt := txt[i2+1:]

	i1 = strings.Index(ntxt, "//")
	if i1 == -1 {
		return errors.New("协议错误")
	}

	i2 = strings.Index(ntxt, ":")
	if i2 == -1 {
		return errors.New("协议错误")
	}

	l.Username = ntxt[i1+2 : i2]

	ntxt = ntxt[i2+1:]

	i2 = strings.Index(ntxt, "@")
	if i2 == -1 {
		return errors.New("协议错误")
	}

	l.Password = ntxt[:i2]

	ntxt = ntxt[i2+1:]

	i2 = strings.Index(ntxt, ":")
	if i2 == -1 {
		return errors.New("协议错误")
	}
	l.Host = ntxt[:i2]

	ntxt = ntxt[i2+1:]

	i2 = strings.Index(ntxt, "?")
	if i2 == -1 {
		return errors.New("协议错误")
	}
	l.Port = ntxt[:i2]

	ntxt = ntxt[i2+1:]

	i1 = strings.Index(ntxt, "=")
	if i1 == -1 {
		return errors.New("协议错误")
	}

	i2 = strings.Index(ntxt, "#")
	if i2 == -1 {
		return errors.New("协议错误")
	}

	paddingStr := ntxt[i1+1 : i2]
	if paddingStr == "true" {
		l.Padding = true
	}

	l.Name = ntxt[i2+1:]

	return nil
}

func (l *Link) ToNaiveConfig(host, port string) *NaiveConfig {
	listen := fmt.Sprintf("socks://%s:%s", host, port)
	proxy := fmt.Sprintf("%s://%s:%s@%s:%s?padding=%v", l.Protocol, l.Username, l.Password, l.Host, l.Port, l.Padding)
	return &NaiveConfig{
		Listen: listen,
		Proxy:  proxy,
		Log:    "",
	}
}

func (l *Link) ToText() string {
	// naive+https://username:password@host:port?padding=false#name
	return fmt.Sprintf(
		"naive+%s://%s:%s@%s:%s?padding=%v#%s",
		l.Protocol,
		l.Username,
		l.Password,
		l.Host,
		l.Port,
		l.Padding,
		l.Name,
	)
}

func (l *Link) ToQCode() (image.Image, error) {
	info := l.ToText()
	png, err := qrcode.Encode(info, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return tools.Image.Png2Image(png)
}
