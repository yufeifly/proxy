package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/container"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/utils"
)

// Stop stop a container
func Stop(c *gin.Context) {
	ContainerID := c.PostForm("ContainerID")
	Timeout := c.PostForm("Timeout")
	address := c.PostForm("Address")

	targetAddress, err := utils.ParseAddress(address)

	stopReqOpts := model.StopReqOpts{
		Address: targetAddress,
		StopOpts: model.StopOpts{
			ContainerID: ContainerID,
			Timeout:     Timeout,
		},
	}

	err = container.StopContainer(stopReqOpts)
	if err != nil {
		utils.ReportErr(c, err)
		panic(err)
	}

	logrus.WithFields(logrus.Fields{
		"ContainerID": ContainerID,
	}).Info("the container has been stopped")

	c.JSON(200, gin.H{
		"result": "success",
	})
}
