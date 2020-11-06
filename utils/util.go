package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// report err to request sender side
func ReportErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"result": err})
}
