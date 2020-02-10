package nlogs

import (
	"fmt"
	"his6/base/message"
	"time"

	"his6/base/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.Logger
)

func init() {
	hook := lumberjack.Logger{
		Filename:   "./log/his6.log", // 日志文件路径
		MaxSize:    128,                // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                 // 日志文件最多保存多少个备份
		MaxAge:     7,                  // 文件最多保存多少天
		Compress:   true,               // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	// 配置文件中获取LEVEL值
	loglevel := config.GetConfigString("logs", "level", "debug")
	switch loglevel {
	case "info":
		atomicLevel.SetLevel(zap.InfoLevel)
		break
	case "notice":
		atomicLevel.SetLevel(zap.WarnLevel)
		break
	case "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
		break
	default:
		atomicLevel.SetLevel(zap.DebugLevel)
		break
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),               // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到控制台和文件
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 由于nlog包装层，调用堆栈增加1层
	callerskip := zap.AddCallerSkip(1)
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	app := config.GetConfigString("app", "name", "")
	filed := zap.Fields(zap.String("serviceName", app))
	// 构造日志
	logger = zap.New(core, caller, callerskip, development, filed)
}

// Debug d
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Info d
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Notice d
func Notice(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
	// 消息发送
	message.Send("notice", msg)
}

// Error d
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
	res := fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05.000"))
	fmt.Println(res + " [Error] " + msg)
	// 消息发送
	message.Send("error", msg)
}
