package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types/svc"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/utils"
	"net/http"
)

// ServiceAdd handler of adding a redis service
func ServiceAdd(c *gin.Context) {
	Service := c.PostForm("Service")
	containerName := c.PostForm("ContainerName")
	RawAddress := c.PostForm("Address") // address of worker node
	address, err := utils.ParseAddress(RawAddress)
	if err != nil {
		utils.ReportErr(c, http.StatusBadRequest, err)
		logrus.Panic(err)
	}
	opts := svc.ServiceOpts{
		CName:    containerName,
		SID:      Service,
		NodeAddr: address,
	}

	scheduler.DefaultRegister(Service, opts)
	//
	cli := client.NewClient(opts.NodeAddr)
	err = cli.AddService(opts)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
