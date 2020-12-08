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
	for _, service := range services {
		logrus.Debugf("services: %v", service)
	}
	c.JSON(http.StatusOK, services)
}
