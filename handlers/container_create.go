package handlers

import (
	"encoding/json"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/container"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/utils"
)

// Create handler for creating a container
func Create(c *gin.Context) {
	header := "container.Create"

	ContainerName := c.PostForm("ContainerName")
	ImageName := c.PostForm("ImageName")
	HostPort := c.PostForm("HostPort")
	ContainerPort := c.PostForm("ContainerPort")
	PortBindings := c.PostForm("PortBindings")
	ExposedPorts := c.PostForm("ExposedPorts")
	CmdParam := c.PostForm("Cmd")

	// if HostPort & ContainerPort are set, then overwrite
	if HostPort != "" && ContainerPort != "" {
		openPort, _ := nat.NewPort("tcp", ContainerPort)
		exposedPorts := nat.PortSet{
			openPort: struct{}{}, //docker容器对外开放的端口
		}
		ExposedPortsSli, err := json.Marshal(exposedPorts)
		if err != nil {
			utils.ReportErr(c, err)
			logrus.Panic(err)
		}
		ExposedPorts = string(ExposedPortsSli)

		portBindings := nat.PortMap{
			openPort: []nat.PortBinding{{
				HostIP:   "0.0.0.0", //docker容器映射的宿主机的ip
				HostPort: HostPort,  //docker 容器映射到宿主机的端口
			}},
		}
		PortBindingsSli, err := json.Marshal(portBindings)
		if err != nil {
			utils.ReportErr(c, err)
			logrus.Panic(err)
		}
		PortBindings = string(PortBindingsSli)
	}

	createOpts := model.CreateOpts{
		ContainerName: ContainerName,
		ImageName:     ImageName,
		HostPort:      HostPort,
		ContainerPort: ContainerPort,
		ExposedPorts:  ExposedPorts,
		PortBindings:  PortBindings,
		Cmd:           CmdParam,
	}
	body, err := container.CreateContainer(createOpts)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Errorf("%s, CreateContainer err: %v", header, err)
		logrus.Panic(err)
	}
	c.JSON(200, gin.H{
		"result":      "success",
		"containerId": body.ID,
	})
}
