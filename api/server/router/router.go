package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/api/server/router/container"
	"github.com/yufeifly/proxy/api/server/router/logger"
	"github.com/yufeifly/proxy/api/server/router/redis"
	"github.com/yufeifly/proxy/api/server/router/service"
)

// InitRoutes init all the routers
func InitRoutes(r *gin.Engine) {
	redis.InitRoutes(r)
	container.InitRoutes(r)
	logger.InitRoutes(r)
	service.InitRoutes(r)
}
