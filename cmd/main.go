package main

import (
	"SerialTest/configs"
	"SerialTest/internal/logger"
	"SerialTest/internal/serialcom"
	"SerialTest/internal/server"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func usage() {
	fmt.Fprintf(os.Stderr, `SerialTest version: serialtest/0.1.0
Usage: server [-s serial device] [-b baud rate] [-l server addr] [-t token]

Options:
`)
	flag.PrintDefaults()
}

func init() {
	h := false
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&(configs.ServerConfig.SerialCom), "s", "COM13", "Serial device name")
	flag.IntVar(&(configs.ServerConfig.BaudRate), "b", 38400, "Serial device baud rate")
	flag.StringVar(&(configs.LogPathConfig), "l", "server.log", "SerialTest server log")
	flag.StringVar(&(configs.ServerConfig.ServerAddr), "p", ":8820", "IP and port monitored by the server")
	flag.StringVar(&(configs.ServerConfig.Token), "t", "cwioehfewowwoefowfh", "Authentication token")
	flag.Usage = usage
	flag.Parse()
	if h {
		flag.Usage()
	}

	sha1 := sha1.New()
	sha1.Write([]byte(configs.ServerConfig.Token))
	configs.ServerConfig.AccessToken = hex.EncodeToString(sha1.Sum([]byte(nil)))

	fmt.Println("**************************************************************************************************************")
	fmt.Println("serial device name is ", configs.ServerConfig.SerialCom)
	fmt.Println("Serial device baud rate is ", configs.ServerConfig.BaudRate)
	fmt.Println("Serial server log is ", configs.LogPathConfig)
	fmt.Println("server addr is ", configs.ServerConfig.ServerAddr)
	fmt.Println("token is ", configs.ServerConfig.Token)
	fmt.Println("access token is ", configs.ServerConfig.AccessToken)
	fmt.Println("Please carry access_token for HTTP request header.Like:access_token:\"4e80c5fc061e67daefea8ab92d9060e5776fe6c4\"")
	fmt.Println("**************************************************************************************************************")
}

func main() {
	if err := logger.AppLogger.Open(); err == nil {
		defer logger.AppLogger.Close()
	}
	var sig chan os.Signal = make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	if err := serialcom.ComOpen(configs.ServerConfig.SerialCom, configs.ServerConfig.BaudRate); err == nil {
		defer serialcom.ComClose()
	}
	go serialcom.RuntimeSerialcom(sig)
	go serialcom.RuntimeScreen()
	go server.RuntimeServer()
	<-sig
	fmt.Println("收到终止信号")
}
