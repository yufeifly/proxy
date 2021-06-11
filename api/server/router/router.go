package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/handlers"
)

// InitRoutes init all the routers
func InitRoutes(r *gin.Engine) {

	r.POST("/container/start", handlers.Start)
	r.GET("/container/list", handlers.ListContainer)
	r.POST("/container/stop", handlers.Stop)

	r.POST("/migrate", handlers.Migrate)

	r.GET("/service/list", handlers.ListService)
	r.POST("/service/add", handlers.ServiceAdd)

	r.GET("/proxy/service/get", handlers.ProxyService)
}
