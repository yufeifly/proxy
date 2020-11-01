package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/scheduler"
	"net/http"
)

// ListService list all services
func ListService(c *gin.Context) {
	services := scheduler.Default().ListService()
	for _, ser := range services {
		logrus.Infof("services: %v", ser)
	}
	c.JSON(http.StatusOK, services)
}
