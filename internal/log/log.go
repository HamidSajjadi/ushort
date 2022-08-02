package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
)

var logLevelMap = map[string]zapcore.Level{
	"trace": zapcore.DebugLevel,
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"err":   zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func New(level string) *zap.SugaredLogger {
	logCfg := zap.NewProductionConfig()
	var logLevel zapcore.Level
	if val, ok := logLevelMap[strings.ToLower(level)]; ok {
		logLevel = val
	} else {
		logLevel = zapcore.InfoLevel
	}
	log.Printf("loaded logger with level `%s` mapped to `%d`\n", level, logLevel)
	logCfg.EncoderConfig.StacktraceKey = zapcore.OmitKey
	logCfg.Level = zap.NewAtomicLevelAt(logLevel)
	logger, err := logCfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	return sugar
}
