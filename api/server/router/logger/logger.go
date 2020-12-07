package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/handlers"
)

// InitRoutes init routes for logger
func InitRoutes(r *gin.Engine) {
	r.POST("/log/consume", handlers.LogConsumedAdder)
}
