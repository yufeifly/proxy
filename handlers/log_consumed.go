package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/scheduler"
	"github.com/yufeifly/proxy/utils"
	"net/http"
)

// LogConsumeAdder
// todo the logic should be reconsidered
func LogConsumedAdder(c *gin.Context) {
	logrus.Info("destination node has consumed a log")
	proxyServiceID := c.PostForm("ProxyServiceID")
	proxyService, err := scheduler.Default().GetService(proxyServiceID)
	if err != nil {
		utils.ReportErr(c, err)
		logrus.Panic(err)
	}
	proxyService.ConsumedAdder()
	logrus.Info("LogConsumedAdder finish")
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
