package handlers

import (
//"github.com/gin-gonic/gin"
//"github.com/sirupsen/logrus"
//"github.com/yufeifly/proxy/container"
//"github.com/yufeifly/proxy/utils"
//"time"
)

// stop a container
//func Stop(c *gin.Context) {
//	ContainerID := c.Query("ContainerID")
//	timeout := time.Second * 10
//
//	//err := cli.ContainerStop(ctx, ContainerID, &timeout)
//	err := container.StopContainer(ContainerID, timeout)
//	if err != nil {
//		utils.ReportErr(c, err)
//		panic(err)
//	}
//
//	logrus.WithFields(logrus.Fields{
//		"ContainerID": ContainerID,
//	}).Info("the container has been stopped")
//
//	c.JSON(200, gin.H{
//		"result": "success",
//	})
//}
