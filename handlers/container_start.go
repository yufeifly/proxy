package handlers

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/container"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/utils"
)

// Start ContainerStart handler
func Start(c *gin.Context) {
	containerID := c.PostForm("ContainerID")
	checkpointID := c.PostForm("CheckpointID")
	checkpointDir := c.PostForm("CheckpointDir")
	address := c.PostForm("Address")

	startOpts := model.StartOpts{
		CStartOpts: types.ContainerStartOptions{
			CheckpointID:  checkpointID,
			CheckpointDir: checkpointDir,
		},
		ContainerID: containerID,
	}

	targetAddress, err := utils.ParseAddress(address)

	startReqOpts := model.StartReqOpts{
		Address:   targetAddress,
		StartOpts: startOpts,
	}

	err = container.StartContainer(startReqOpts)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	c.JSON(200, gin.H{"result": "success"})
}
