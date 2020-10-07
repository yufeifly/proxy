package main

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/server/router"
)

func main() {
	r := router.InitRoutes()
	if err := r.Run(":6788"); err != nil {
		logrus.Errorf("gin.run err: %v", err)
	}
}
