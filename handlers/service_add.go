package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/model"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/utils"
)

func ServiceAdd(c *gin.Context) {
	ProxyService := c.Query("Service3")
	RawAddress := c.Query("Address")
	address, err := utils.ParseAddress(RawAddress)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	service := scheduler.NewService(model.ServiceOpts{
		ID:             utils.NameServiceByProxyService(ProxyService),
		ProxyServiceID: ProxyService,
		NodeAddr:       address,
	})
	scheduler.Default().AddService(ProxyService, service)
	//
	cli := client.Client{}
	cli.AddService(service)
}
