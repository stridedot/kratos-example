package log

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/natefinch/lumberjack"
)

var _ log.Logger = (*ZapLogger)(nil)

// ZapLogger is a logger impl.
type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

// NewZapLogger return a zap logger.
func NewZapLogger(encoder zapcore.Encoder, level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {
	// 日志保存方式
	writeSyncer := getWriterSyncer()

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(writeSyncer),
		), level)
	zapLogger := zap.New(core, opts...)
	return &ZapLogger{log: zapLogger, Sync: zapLogger.Sync}
}

func Logger(mode string) log.Logger {
	encoder := getEncoder(mode)
	
	return NewZapLogger(
	   encoder,
	   zap.NewAtomicLevelAt(zapcore.DebugLevel),
	   zap.AddStacktrace(
		  zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
	   zap.AddCaller(),
	   zap.AddCallerSkip(2),
	   zap.Development(),
	)
 }

 // getLogWriter 日志自动切割，采用 lumberjack 实现
func getWriterSyncer() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
	   Filename:   "../../logs/zap.log", // 指定日志存储位置
	   MaxSize:    10,                   // 日志的最大大小（M）
	   MaxBackups: 5,                    // 日志的最大保存数量
	   MaxAge:     30,                   // 日志文件存储最大天数
	   Compress:   false,                // 是否执行压缩
	}
	return zapcore.AddSync(lumberJackLogger)
 }

 // getEncoder 以不同格式写入文件
func getEncoder(mode string) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	if mode == "dev" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

// Log Implementation of logger interface.
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	case log.LevelFatal:
		l.log.Fatal("", data...)
	}
	return nil
}
