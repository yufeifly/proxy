package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/handlers"
)

// InitRoutes init routes for services
func InitRoutes(r *gin.Engine) {
	r.GET("/service/list", handlers.ListService)
	r.POST("/service/add", handlers.ServiceAdd)
	r.POST("/service/migrate", handlers.MigrateService)
}
