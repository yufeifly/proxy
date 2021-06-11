package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/scheduler"
	"net/http"
)

func ProxyService(c *gin.Context) {
	containerServ := c.Query("Service")
	serv, err := scheduler.Default().GetService(containerServ)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	node := serv.GetNode()

	c.JSON(http.StatusOK, node)
}
