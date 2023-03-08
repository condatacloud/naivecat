package tools

import (
	"bytes"
	"regexp"
	"strings"
)

type IStrings interface {
	// 判断字符串尾部是否是某字符串
	HasSuffix(s, suffix string) bool
	// 去除newline换行符号
	TrimN(str string) string
	// 去字符串中的字符
	Strip(str string, c string) string
	// 重复一定次数某一个字符串变成新的字符串
	RepeatStr(s string, n int) string
	// 对json数据进行严格的修正，会去除注释
	StrictJSON(json []byte) ([]byte, error)
	// 去除Json的尾随逗号
	TrimJsonTrailComma(json []byte) ([]byte, error)
	// 去除字符串中的终端颜色 转义码
	StripANSI(str string) string
	// 字符串首字母大写
	FirstUpper(s string) string
	// 去除字符串尾部斜杠
	TrimTrailSlash(s string) string
}

type cstrings struct{}

var Strings IStrings = &cstrings{}

func (c *cstrings) HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func (c *cstrings) TrimN(str string) string {
	return strings.Replace(str, "\n", "", -1)
}

func (c *cstrings) Strip(str string, char string) string {
	return strings.Replace(str, char, "", -1)
}

func (c *cstrings) RepeatStr(s string, n int) string {
	b := bytes.Buffer{}
	for i := 0; i < n; i++ {
		b.WriteString(s)
	}
	return b.String()
}

func (c *cstrings) StrictJSON(json []byte) ([]byte, error) {
	// v1:
	// 去除以 // 开头的注释 和 /*...*/ 的注释
	// re := regexp.MustCompile("//.*?\n|/\\*.*?\\*/")

	// v2:
	// |左边为\s代表匹配任何空白字符，包括空格、制表符、换页符，也就是必须是
	// 空格+//，防止http://被误伤，但是这也会导致定格的注释不能去除
	// re := regexp.MustCompile("\\s//.*?\n|/\\*.*?\\*/")

	// v3:
	// 增加 ^//.*?\n 解决定格写的注释
	re := regexp.MustCompile("\\s//.*?\n|/\\*.*?\\*/|^//.*?\n")
	newBytes := re.ReplaceAll(json, nil)
	return c.TrimJsonTrailComma(newBytes)
}

func (c *cstrings) TrimJsonTrailComma(json []byte) ([]byte, error) {
	re := regexp.MustCompile(`,((\s*?[\}\]])|(\s*?$))`)
	newBytes := re.ReplaceAllFunc(json, func(b []byte) []byte {
		for i, v := range b {
			if v == ',' {
				b[i] = ' '
			}
		}
		return b
	})
	return newBytes, nil
}

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansi)

func (c *cstrings) StripANSI(str string) string {
	return re.ReplaceAllString(str, "")
}

func (c *cstrings) FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// 去除字符串尾部斜杠
func (c *cstrings) TrimTrailSlash(s string) string {
	if strings.HasSuffix(s, "/") {
		return s[:len(s)-1]
	}
	if strings.HasSuffix(s, "\\") {
		return s[:len(s)-2]
	}
	return s
}
