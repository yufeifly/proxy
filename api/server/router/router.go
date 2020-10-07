package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/api/server/router/container"
	"github.com/yufeifly/proxy/api/server/router/redis"
)

func InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	redis.InitRoutes(r)
	container.InitRoutes(r)
	return r
}
