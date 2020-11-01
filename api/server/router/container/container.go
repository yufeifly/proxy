package container

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/handlers"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/container/run", handlers.Run)
	// start a container
	r.POST("/container/start", handlers.Start)
	//  list containers
	r.GET("/container/list", handlers.ListContainer)
	//  stop a container
	r.POST("/container/stop", handlers.Stop)
	//  create a container
	//r.POST("/container/create", handlers.Create)
	//  create a container checkpoint
	//r.POST("/container/checkpoint/create", handlers.CheckpointCreate)
	// receive checkpoint and restore from it
	//r.POST("/container/checkpoint/restore", handlers.FetchCheckpointAndRestore)
	// push checkpoint to destination
	//r.POST("/container/checkpoint/push", handlers.CheckpointPush)
	// migrate a container
	r.POST("/container/migrate", handlers.MigrateContainer)
}
