package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/server/router"
	"github.com/yufeifly/proxy/cluster"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/utils"
)

func init() {
	if utils.IsDebugEnabled() {
		utils.EnableDebug()
	}
}

func main() {
	// loading cluster
	err := cluster.LoadClusterConfig()
	if err != nil {
		logrus.Panicf("LoadClusterConfig failed, err: %v", err)
	}
	// init default scheduler
	scheduler.InitScheduler()
	// register services
	scheduler.PseudoRegister()
	// get gin engine and init API routes
	//if !utils.IsDebugEnabled(){
	//	gin.SetMode(gin.ReleaseMode)
	//}
	r := gin.Default()
	router.InitRoutes(r)
	if err := r.Run(":6788"); err != nil {
		logrus.Errorf("gin.run err: %v", err)
	}
}
