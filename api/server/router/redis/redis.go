package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/handlers"
)

// InitRoutes init routes for redis operation
func InitRoutes(r *gin.Engine) {
	r.GET("/redis/get", handlers.Get)
	// redis set func
	r.POST("/redis/set", handlers.Set)
	// @deprecated redis migrate
	//r.POST("/redis/migration", r.MigrateRedis)
}
