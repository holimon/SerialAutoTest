package logger

import (
	"SerialTest/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var AppLogger, _ = zap.NewProduction()

func LoggerInit() {
	hook := lumberjack.Logger{
		Filename:   configs.LogPathConfig, // 日志文件路径
		MaxSize:    1,                     // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 1,                     // 日志文件最多保存多少个备份
		MaxAge:     7,                     // 文件最多保存多少天
		Compress:   false,                 // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "AppLogger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到文件
		atomicLevel, // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	//开启文件及行号
	development := zap.Development()
	AppLogger = zap.New(core, caller, development)
	//AppLogger = zap.New(core)
	return
}
