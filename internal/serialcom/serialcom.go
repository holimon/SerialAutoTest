package serialcom

import (
	"SerialTest/configs"
	"SerialTest/internal/logger"
	"SerialTest/internal/utils"
	"github.com/tarm/goserial"
	"go.uber.org/zap"
	"io"
	"os"
	"strconv"
	"syscall"
	"time"
)

type OpenComType struct {
	sercom io.ReadWriteCloser
	opened bool
}

var OpenCom OpenComType = OpenComType{opened: false}

func ComOpen(com string, baud int) (err error) {
	sercfg := &serial.Config{Name: com, Baud: baud, ReadTimeout: 1000 * time.Millisecond}
	OpenCom.sercom, err = serial.OpenPort(sercfg)
	if err == nil {
		OpenCom.opened = true
	}
	return
}

func ComClose() (err error) {
	err = OpenCom.sercom.Close()
	if err == nil {
		OpenCom.opened = false
	}
	return
}

func ComRead(p []byte) (n int, err error) {
	n, err = OpenCom.sercom.Read(p)
	return
}

func ComWrite(p []byte) (n int, err error) {
	if OpenCom.opened {
		n, err = OpenCom.sercom.Write(p)
	}
	return
}

func RuntimeSerialcom(sig chan os.Signal) {
	err := ComOpen(configs.ServerConfig.SerialCom, configs.ServerConfig.BaudRate)
	if err != nil {
		logger.AppLogger.Error("打开串口设备失败")
		sig <- syscall.SIGTERM
		return
	}
	defer ComClose()
	for {
		readbuf := make([]byte, 1024)
		rlen, err := ComRead(readbuf)
		if rlen == 0 || err != nil {
			continue
		}
		readbuf = readbuf[:rlen]
		logger.AppLogger.Info("read", zap.String("content", string(readbuf)))
		//将读取到的串口数据推送到注册地址
		for _, v := range configs.ClinetListened {
			params := map[string]string{"content": string(readbuf), "timestamp": strconv.FormatInt(time.Now().UnixNano(), 10)}
			utils.HttpGet(params, v)
		}
	}
}
