package service

import (
	"context"
	"errors"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

type IPing interface {
	Ping(ip string) (int64, error)
	Pings(ips []string) ([]int64, error)
}

type pingService struct {
	timeout int
}

var PingService IPing = &pingService{
	timeout: 4,
}

func (s *pingService) Ping(ip string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*time.Duration(s.timeout)))
	defer cancel()

	pinger, err := probing.NewPinger(ip)
	if err != nil {
		return 0, err
	}
	pinger.Count = 2

	execResult := make(chan bool)
	go func(execResult chan<- bool) {
		err = pinger.Run() // Blocks until finished.
		execResult <- true
	}(execResult)

	// 等待结果
	select {
	case <-ctx.Done():
		return 0, errors.New("超时")
	case <-execResult:
		if err != nil {
			return 0, errors.New("错误")
		}
		stats := pinger.Statistics()
		return stats.AvgRtt.Milliseconds(), nil
	}
}

func (s *pingService) Pings(ips []string) ([]int64, error) {
	return nil, nil
}
