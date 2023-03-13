package service

import (
	"bufio"
	"context"
	"io"
	"naivecat/model"
	"naivecat/tools"
	"os/exec"
	"runtime"
	"time"
)

type NaiveLogCallback func(line string)

type INaive interface {
	GetVersion() string
	Start() error
	Close()
	UpdateConfig(config *model.NaiveConfig) error
	SetLogCallback(callback NaiveLogCallback)
	IsRunning() bool
}

type naiveService struct {
	naiveFilePath string
	naiveConfPath string
	logCallback   NaiveLogCallback
	cancel        context.CancelFunc
	running       bool
}

var NaiveService INaive = newNaiveService()

func newNaiveService() *naiveService {

	home, err := tools.HomeDir()
	if err != nil {
		panic(err)
	}

	ns := &naiveService{}

	ns.naiveFilePath = home + "/.naivecat/naive"
	if runtime.GOOS == "windows" {
		ns.naiveFilePath = home + "/.naivecat/naive.exe"
	}

	ns.naiveConfPath = home + "/.naivecat/naive_config.json"

	return ns
}

func (s *naiveService) GetVersion() string {
	cmd := exec.Command(s.naiveFilePath, "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return string(out)
}

func (s *naiveService) Start() error {
	s.running = true
	defer func() {
		if err := recover(); err != nil {
			model.Log.Error(err)
		}
		s.running = false
	}()
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	cmd := exec.CommandContext(ctx, s.naiveFilePath, s.naiveConfPath)
	hideCommandWindow(cmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		return err
	}

	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, _, err := reader.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		if s.logCallback != nil {
			s.logCallback(string(line))
		}
	}
	err = cmd.Wait()
	if s.cancel != nil {
		s.cancel()
		s.cancel = nil
	}

	if err != nil && !s.running {
		return nil
	}
	return err
}

func (s *naiveService) Close() {
	s.running = false
	if s.cancel != nil {
		s.cancel()
		s.cancel = nil
	}
	s.logCallback = nil
	time.Sleep(1 * time.Second) // 防止没有完全停止
}

func (s *naiveService) UpdateConfig(config *model.NaiveConfig) error {
	return tools.Serialize(config, s.naiveConfPath)
}

func (s *naiveService) SetLogCallback(callback NaiveLogCallback) {
	s.logCallback = callback
}

func (s *naiveService) IsRunning() bool {
	return s.running
}
