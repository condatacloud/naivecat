package tools

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type IFile interface {
	// 读取文件内的所有行
	ReadLines(filePath string) ([]string, error)
	// 判断给定的路径/文件是否存在
	Exists(path string) bool
	// 判断给定的路径是否是文件夹
	IsDir(path string) bool
	// 判断给定的路径是否是文件
	IsFile(path string) bool
	// 创建文件夹
	Mkdir(dest string) error
	// 写内容到文件，没有文件会自动创建
	Write(text string, filePath string) error
	// 写二进制到文件，没有文件会自动创建
	WriteBin(bytes []byte, filePath string) error
}

type file struct{}

var File IFile = &file{}

func (f *file) ReadLines(filePath string) ([]string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func (f *file) Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func (f *file) ExistsLink(path string) bool {
	if _, err := os.Lstat(path); err == nil {
		return true
	}
	return false
}

func (f *file) IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func (f *file) IsFile(path string) bool {
	return !f.IsDir(path)
}

func (f *file) Mkdir(dest string) error {
	//分割path目录
	dest = strings.Replace(dest, "\\", "/", -1)
	destSplitPathDirs := strings.Split(dest, "/")
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs) {
			destSplitPath = destSplitPath + dir + "/"
			b := f.Exists(destSplitPath)
			if !b {
				err := os.Mkdir(destSplitPath, 0775)
				if err != nil {
					return fmt.Errorf("failed to create %s directory: %s", destSplitPath, err.Error())
				}
			}
		}
	}
	return nil
}

func (f *file) Write(text string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed to open %s file: %s", filePath, err.Error())
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	if _, err := write.WriteString(text); err != nil {
		return fmt.Errorf("failed to write %s file: %s", filePath, err.Error())
	}
	if err := write.Flush(); err != nil {
		return fmt.Errorf("failed to write %s file: %s", filePath, err.Error())
	}
	return nil
}

func (f *file) WriteBin(bytes []byte, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed to open %s file: %s", filePath, err.Error())
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	if _, err := write.Write(bytes); err != nil {
		return fmt.Errorf("failed to write %s file: %s", filePath, err.Error())
	}
	if err := write.Flush(); err != nil {
		return fmt.Errorf("failed to write %s file: %s", filePath, err.Error())
	}
	return nil
}
