package http_engine

import (
	"github.com/HamidSajjadi/ushort/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	ProductionMode  = "production"
	DevelopmentMode = "development"
)

func New(deploymentMode string, logger *zap.SugaredLogger) *gin.Engine {
	var httpStub *gin.Engine
	if deploymentMode == ProductionMode {
		gin.SetMode(gin.ReleaseMode)
		httpStub = gin.New()
		httpStub.Use(gin.Recovery())
	} else {
		httpStub = gin.Default()
	}
	httpStub.Use(api.ErrorHandler(logger))
	return httpStub
}
