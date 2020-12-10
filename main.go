package main

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/server/router"
	"github.com/yufeifly/proxy/cluster"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/ticket"
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
	// init ticket
	ticket.InitTicket()
	//
	r := router.InitRoutes()
	if err := r.Run(":6788"); err != nil {
		logrus.Errorf("gin.run err: %v", err)
	}
}
