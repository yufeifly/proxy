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
	header := "redis.Get"
	// get params
	key := c.Query("key")
	ProxyService := c.Query("service")
	// verify params
	if key == "" || ProxyService == "" {
		logrus.Errorf("%s, err: %v", header, cusErr.ErrBadParams)
		c.JSON(http.StatusBadRequest, gin.H{"failed: ": cusErr.ErrBadParams.Error()})
		return
	}

	val, err := redis.Get(ProxyService, key)
	if err != nil {
		logrus.Errorf("%s, err: %v", header, err)
		c.JSON(200, gin.H{"failed: ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"value: ": val})
}

// Set redis set
func Set(c *gin.Context) {
	header := "redis.Set"
	// get params
	key := c.Query("key")
	val := c.Query("value")
	ProxyService := c.Query("service")
	// verify params
	if key == "" || val == "" || ProxyService == "" {
		logrus.Errorf("%s, err: %v", header, cusErr.ErrBadParams)
		c.JSON(http.StatusBadRequest, gin.H{"failed: ": cusErr.ErrBadParams.Error()})
		return
	}
	//
	err := redis.Set(ProxyService, key, val)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": val,
		}).Error("set pair failed")
		c.JSON(200, gin.H{"failed: ": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": "success"})
}
