package api

import (
	"gix/config"

	"github.com/gin-gonic/gin"
)

type SEngine struct {
	*gin.Engine
}

func NewApp() *SEngine {
	gin.SetMode(gin.ReleaseMode)
	return &SEngine{
		Engine: gin.New(),
	}
}

func (s *SEngine) Defalut() *SEngine {
	s.Engine = gin.Default()
	return s
}

func (s *SEngine) Run() error {
	return s.Engine.Run(":" + config.GetConf().GetString("serve.port"))
}
