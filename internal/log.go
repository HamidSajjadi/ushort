package internal

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.SugaredLogger {
	logCfg := zap.NewProductionConfig()
	logCfg.EncoderConfig.StacktraceKey = zapcore.OmitKey
	logger, err := logCfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	return sugar
}
