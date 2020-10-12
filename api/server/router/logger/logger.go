package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/handlers"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/log/consumed", handlers.LogConsumeAdder)
}
