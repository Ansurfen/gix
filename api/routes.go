package api

import (
	. "gix/config"

	"github.com/gin-gonic/gin"
)

func (s *SEngine) UseRouter() *SEngine {
	s.POST("/", Push)
	return s
}

func Push(ctx *gin.Context) {
	topic, msg := GetLvm().PreHandler(ctx.PostForm("msg"))
	GetMQ().Pub(topic, msg)
}
