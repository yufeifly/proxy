package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/cuserr"
	"github.com/yufeifly/proxy/redis"
	"github.com/yufeifly/proxy/utils"
	"net/http"
)

// Get redis get
func Get(c *gin.Context) {
	// get params
	ProxyService := c.Query("service")
	key := c.Query("key")

	// verify params
	if ProxyService == "" || key == "" {
		logrus.Errorf("%s, err: %v", "handlers.Get", cuserr.ErrBadParams)
		utils.ReportErr(c, http.StatusBadRequest, cuserr.ErrBadParams)
		return
	}

	val, err := redis.Get(ProxyService, key)
	if err != nil {
		logrus.Errorf("%s, err: %v", "handlers.Get", err)
		utils.ReportErr(c, http.StatusInternalServerError, err)
		return
	}
	c.String(http.StatusOK, val)
}

// Set redis set
func Set(c *gin.Context) {
	// get params
	ProxyService := c.PostForm("service")
	key := c.PostForm("key")
	value := c.PostForm("value")

	// verify params
	if ProxyService == "" || key == "" || value == "" {
		logrus.Errorf("handlers.Set, err: %v", cuserr.ErrBadParams)
		c.JSON(http.StatusBadRequest, gin.H{"failed: ": cuserr.ErrBadParams.Error()})
		return
	}
	// do set
	err := redis.Set(ProxyService, key, value)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": value,
		}).Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"failed: ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"set": "success"})
}

// Delete
func Delete(c *gin.Context) {
	// get params
	ProxyService := c.PostForm("service")
	key := c.PostForm("key")
	// verify params
	if ProxyService == "" || key == "" {
		logrus.Errorf("handlers.Delete, err: %v", cuserr.ErrBadParams)
		c.JSON(http.StatusBadRequest, gin.H{"failed: ": cuserr.ErrBadParams.Error()})
		return
	}
	// do delete
	err := redis.Delete(ProxyService, key)
	if err != nil {
		logrus.Errorf("handlers.Delete, err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"failed: ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete": "success"})
}
