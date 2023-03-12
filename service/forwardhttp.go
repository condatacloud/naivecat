package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

type ForwardInfo struct {
	SrcIP   string
	SrcPort string
	DstIP   string
	DstPort string
}

type IForwardHttp interface {
	Start(info *ForwardInfo) error
	Close() error
	IsRunning() bool
}

type forwardHttpService struct {
	ForwardInfo *ForwardInfo
	server      *http.Server
	running     bool
}

var ProxyService IForwardHttp = &forwardHttpService{}

func (s *forwardHttpService) Start(info *ForwardInfo) error {
	s.running = true
	defer func() {
		s.running = false
	}()
	s.ForwardInfo = info
	s.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.ForwardInfo.DstIP, s.ForwardInfo.DstPort),
		Handler:      http.HandlerFunc(s.serveHTTP),
		ReadTimeout:  100 * time.Minute,
		WriteTimeout: 100 * time.Minute,
	}

	err := s.server.ListenAndServe()
	if err.Error() == "http: Server closed" {
		return nil
	}

	return err
}

func (s *forwardHttpService) Close() error {
	if s.running {
		s.running = false
	}

	// 使用context控制srv.Shutdown的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		return errors.New("关闭http代理错误")
	}
	return nil
}

func (s *forwardHttpService) IsRunning() bool {
	return s.running
}

func (s *forwardHttpService) handleHTTP(w http.ResponseWriter, req *http.Request, dialer proxy.Dialer) {
	tp := http.Transport{
		Dial: dialer.Dial,
	}
	resp, err := tp.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	s.copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (s *forwardHttpService) copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func (s *forwardHttpService) handleTunnel(w http.ResponseWriter, req *http.Request, dialer proxy.Dialer) {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	srcConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	dstConn, err := dialer.Dial("tcp", req.Host)
	if err != nil {
		srcConn.Close()
		return
	}

	srcConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))

	go s.transfer(dstConn, srcConn)
	go s.transfer(srcConn, dstConn)
}

func (s *forwardHttpService) transfer(dst io.WriteCloser, src io.ReadCloser) {
	defer dst.Close()
	defer src.Close()

	io.Copy(dst, src)
}

func (s *forwardHttpService) serveHTTP(w http.ResponseWriter, req *http.Request) {
	d := &net.Dialer{
		Timeout: 10 * time.Second,
	}
	dialer, _ := proxy.SOCKS5("tcp", fmt.Sprintf("%s:%s", s.ForwardInfo.SrcIP, s.ForwardInfo.SrcPort), nil, d)

	if req.Method == "CONNECT" {
		s.handleTunnel(w, req, dialer)
	} else {
		s.handleHTTP(w, req, dialer)
	}
}
