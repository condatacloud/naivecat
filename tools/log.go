package tools

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	log     *log.Logger
	logFile *os.File
	level   int
}

// https://blog.51cto.com/u_10125763/3697502

const (
	DEBUG = 0
	INFO  = 1
	ERROR = 2
)

func NewLog(folder string, level ...string) *Logger {
	lel := ERROR
	if len(level) != 0 {
		l := level[0]
		if l == "debug" {
			lel = DEBUG
		} else if l == "error" {
			lel = ERROR
		} else if l == "info" {
			lel = INFO
		} else {
			lel = ERROR
		}
	}
	logFile, err := os.OpenFile(folder+"/system.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faild to open error logger file:", err)
	}

	logger := &Logger{}
	// 日期，时间，文件名
	logger.log = log.New(logFile, "DEFAULT: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.logFile = logFile
	logger.level = lel
	return logger
}

func (l *Logger) Close() {
	l.logFile.Close()
}

func (l *Logger) Debug(v ...interface{}) {
	if DEBUG >= l.level {
		l.log.SetPrefix("DEBUG: ")
		l.log.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	if INFO >= l.level {
		l.log.SetPrefix("INFO: ")
		l.log.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logger) Error(v ...interface{}) {
	if ERROR >= l.level {
		l.log.SetPrefix("ERROR: ")
		l.log.Output(2, fmt.Sprintln(v...))
	}
}
