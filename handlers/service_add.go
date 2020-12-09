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
	ProxyService := c.PostForm("Service")
	RawAddress := c.PostForm("Address") // address of worker node
	address, err := utils.ParseAddress(RawAddress)
	if err != nil {
		utils.ReportErr(c, http.StatusBadRequest, err)
		logrus.Panic(err)
	}
	opts := svc.ServiceOpts{
		ID:             utils.NameServiceByProxyService(ProxyService),
		ProxyServiceID: ProxyService,
		NodeAddr:       address,
	}

	scheduler.DefaultRegister(ProxyService, opts)
	//
	cli := client.NewClient(opts.NodeAddr)
	err = cli.AddService(opts)
	if err != nil {
		utils.ReportErr(c, http.StatusInternalServerError, err)
		logrus.Panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
