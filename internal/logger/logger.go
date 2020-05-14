package logger

import (
	"SerialTest/configs"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

type Log struct {
	logObj   *os.File
	logWri   *bufio.Writer
	mutexWri sync.Mutex
	opened   bool
}

type Logger interface {
	Open() (err error)
	Write(data []byte) (len int, err error)
	Close()
}

var AppLogger Logger = &Log{}

func (l *Log) Open() (err error) {
	if _, err := os.Stat(configs.LogPathConfig); err == nil || os.IsExist(err) {
		if err := os.Remove(configs.LogPathConfig); err != nil {
			fmt.Println("已存在日志文件删除失败")
		}
	}
	if l.logObj, err = os.OpenFile(configs.LogPathConfig, os.O_RDWR|os.O_CREATE, 0644); err == nil {
		l.logWri = bufio.NewWriterSize(l.logObj, 32)
		l.opened = true
	}
	return
}

func (l *Log) Write(data []byte) (len int, err error) {
	if l.opened {
		l.mutexWri.Lock()
		bytes.ReplaceAll(data, []byte(`\n`), []byte(time.Now().Format("2006-01-02 15:04:05")+"\r\n"))
		//data = time.Now().Format("2006-01-02 15:04:05") + " : " + data
		len, err = l.logWri.Write(data)
		l.logWri.Flush()
		l.mutexWri.Unlock()
	}
	return
}

func (l *Log) Close() {
	if l.opened {
		defer l.logObj.Close()
		l.opened = false
	}
}
