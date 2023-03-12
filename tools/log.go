package tools

import (
	"log"
	"os"
)

type Logger struct {
	debugLog  *log.Logger
	errorLog  *log.Logger
	debugFile *os.File
	errorFile *os.File
}

func NewLog(folder string) *Logger {
	debugFile, err := os.OpenFile(folder+"/debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faild to open error logger file:", err)
	}

	errorFile, err := os.OpenFile(folder+"/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faild to open error logger file:", err)
	}

	logger := &Logger{}
	// 日期，时间，文件名
	logger.debugLog = log.New(debugFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.errorLog = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.debugFile = debugFile
	logger.errorFile = errorFile
	return logger
}

func (l *Logger) Close() {
	l.debugFile.Close()
	l.errorFile.Close()
}

func (l *Logger) Debug(txt string) {
	l.debugLog.Println(txt)
}

func (l *Logger) Error(txt string) {
	l.errorLog.Println(txt)
}
