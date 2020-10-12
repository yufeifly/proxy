package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yufeifly/proxy/wal"
)

func LogConsumeAdder(c *gin.Context) {
	wal.ConsumedAddOne()
	c.JSON(200, gin.H{"result": "success"})
}
