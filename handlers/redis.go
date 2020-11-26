package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/cusErr"
	"github.com/yufeifly/proxy/redis"
	"net/http"
)

// Get redis get
func Get(c *gin.Context) {
	header := "handlers.Get"
	// get params
	ProxyService := c.Query("service")
	key := c.Query("key")

	// verify params
	if ProxyService == "" || key == "" {
		logrus.Errorf("%s, err: %v", header, cusErr.ErrBadParams)
		c.JSON(http.StatusBadRequest, gin.H{"failed: ": cusErr.ErrBadParams.Error()})
		return
	}

	val, err := redis.Get(ProxyService, key)
	if err != nil {
		logrus.Errorf("%s, err: %v", header, err)
		c.JSON(http.StatusOK, gin.H{"failed: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value: ": val})
}

// Set redis set
func Set(c *gin.Context) {
	// get params
	ProxyService := c.PostForm("service")
	key := c.PostForm("key")
	value := c.PostForm("value")

	// verify params
	if ProxyService == "" || key == "" || value == "" {
		logrus.Errorf("handlers.Set, err: %v", cusErr.ErrBadParams)
		c.JSON(http.StatusBadRequest, gin.H{"failed: ": cusErr.ErrBadParams.Error()})
		return
	}
	// do set
	err := redis.Set(ProxyService, key, value)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": value,
		}).Error("set pair failed")
		c.JSON(http.StatusInternalServerError, gin.H{"failed: ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
