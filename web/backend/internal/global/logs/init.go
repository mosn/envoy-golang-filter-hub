package logs

import (
	"envoy-go-fliter-hub/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
)

const (
	defaultLogFilePath = "./logs"
	defaultLogFileName = "log"
)

var logger *zap.Logger

func Init() {
	switch config.Config.Mode {
	case config.ReleaseMode:
		initReleaseLogger()
	default:
		initDebugLogger()
	}
}

func initDebugLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	logger, err = cfg.Build(zap.AddCallerSkip(1), zap.AddCaller())
	if err != nil {
		panic(err)
	}
}

func initReleaseLogger() {
	var err error
	logger, err = zap.NewProduction(
		zap.AddCallerSkip(1),
		zap.AddCaller(),
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return newLogWriteCore()
		}),
	)
	if err != nil {
		panic(err)
	}
}

func newLogWriteCore() zapcore.Core {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(cfg)
	ws := getLogWriter(config.Config.Log.FilePath)
	return zapcore.NewCore(encoder, ws, zap.InfoLevel)
}

func getLogWriter(path string) zapcore.WriteSyncer {
	if path == "" {
		path = defaultLogFilePath
	}

	path = filepath.Clean(path)
	logFileName := filepath.Join(path, defaultLogFileName) // path + filename
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFileName, // 日志文件路径
		MaxSize:    128,         // 设置日志文件最大尺寸 MB
		MaxBackups: 8,           // 设置日志文件最多保存多少个备份
		MaxAge:     30,          // 设置日志文件最多保存多少天
		Compress:   true,        // 是否压缩 disabled by default
	}

	return zapcore.AddSync(lumberJackLogger)
}
