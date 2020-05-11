package serialcom

import (
	"SerialTest/configs"
	"SerialTest/internal/logger"
	"SerialTest/internal/utils"
	"bufio"
	"fmt"
	"github.com/tarm/goserial"
	"go.uber.org/zap"
	"io"
	"os"
	"strconv"
	"time"
)

type OpenComType struct {
	sercom io.ReadWriteCloser
	opened bool
}

var OpenCom OpenComType = OpenComType{opened: false}

func ComOpen(com string, baud int) (err error) {
	if OpenCom.opened {
		return
	}
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

func RuntimeSerialcom() {
	for {
		if !OpenCom.opened {
			continue
		}
		readbuf := make([]byte, 1024)
		rlen, err := ComRead(readbuf)
		if rlen == 0 || err != nil {
			continue
		}
		readbuf = readbuf[:rlen]
		fmt.Print(string(readbuf))
		logger.AppLogger.Info("read", zap.String("content", string(readbuf)))
		//将读取到的串口数据推送到注册地址
		for _, v := range configs.ClinetListened {
			params := map[string]string{"content": string(readbuf), "timestamp": strconv.FormatInt(time.Now().UnixNano(), 10)}
			utils.HttpGet(params, v)
		}
	}
}

func RuntimeScreen() {
	for {
		if OpenCom.opened {
			break
		}
	}
	cmdio := bufio.NewReader(os.Stdin)
	for {
		if cmd, err := cmdio.ReadString('\n'); err == nil {
			ComWrite([]byte(cmd))
		}
	}
}
