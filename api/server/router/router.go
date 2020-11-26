package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/api/server/router/container"
	"github.com/yufeifly/proxy/api/server/router/logger"
	"github.com/yufeifly/proxy/api/server/router/redis"
	"github.com/yufeifly/proxy/api/server/router/service"
)

// InitRoutes init all the routers
func InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	redis.InitRoutes(r)
	container.InitRoutes(r)
	logger.InitRoutes(r)
	service.InitRoutes(r)
	return r
}
