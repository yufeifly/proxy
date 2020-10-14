package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/wal"
)

func LogConsumeAdder(c *gin.Context) {
	logrus.Info("consumed log add one")
	wal.ConsumedAddOne()
	c.JSON(200, gin.H{"result": "success"})
}
