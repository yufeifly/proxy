package handlers

import (
	dockertypes "github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/container"
	"github.com/yufeifly/proxy/utils"
	"net/http"
)

// Start ContainerStart handler
func Start(c *gin.Context) {
	containerID := c.PostForm("ContainerID")
	checkpointID := c.PostForm("CheckpointID")
	checkpointDir := c.PostForm("CheckpointDir")
	address := c.PostForm("Address")

	startOpts := types.StartOpts{
		CStartOpts: dockertypes.ContainerStartOptions{
			CheckpointID:  checkpointID,
			CheckpointDir: checkpointDir,
		},
		ContainerID: containerID,
	}

	targetAddress, err := utils.ParseAddress(address)

	startReqOpts := container.StartReqOpts{
		Address:   targetAddress,
		StartOpts: startOpts,
	}

	err = container.StartContainer(startReqOpts)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
