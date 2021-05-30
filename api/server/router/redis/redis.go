package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/handlers"
)

// InitRoutes init routes for redis operation
func InitRoutes(r *gin.Engine) {
	r.GET("/redis/get", handlers.Get)
	r.POST("/redis/set", handlers.Set)
	r.POST("/redis/delete", handlers.Delete)
}
