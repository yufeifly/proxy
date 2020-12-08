package utils

import (
	"github.com/gin-gonic/gin"
)

// ReportErr report err to request sender side
func ReportErr(c *gin.Context, httpCode int, err error) {
	c.JSON(httpCode, gin.H{"result": err})
}
