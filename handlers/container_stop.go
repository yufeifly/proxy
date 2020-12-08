package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/container"
	"github.com/yufeifly/proxy/utils"
	"net/http"
)

// Stop stop a container
func Stop(c *gin.Context) {
	ContainerID := c.PostForm("ContainerID")
	Timeout := c.PostForm("Timeout")
	address := c.PostForm("Address")

	targetAddress, err := utils.ParseAddress(address)

	stopReqOpts := container.StopReqOpts{
		Address: targetAddress,
		StopOpts: types.StopOpts{
			ContainerID: ContainerID,
			Timeout:     Timeout,
		},
	}

	err = container.StopContainer(stopReqOpts)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		panic(err)
	}

	logrus.WithFields(logrus.Fields{
		"ContainerID": ContainerID,
	}).Info("the container has been stopped")

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
