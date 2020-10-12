package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/redis"
)

// Get redis get
func Get(c *gin.Context) {
	header := "redis.Get"
	key := c.Query("key")

	val, err := redis.Get(key)
	if err != nil {
		logrus.Errorf("%s, err: %v", header, err)
		logrus.Panic(err)
	}
	c.JSON(200, gin.H{"value: ": val})
}

// Set redis set
func Set(c *gin.Context) {

	key := c.Query("key")
	val := c.Query("value")

	// todo validate kv pair

	err := redis.Set(key, val)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": val,
		}).Error("set pair failed")
		logrus.Panic(err)
	}

	c.JSON(200, gin.H{"result": "success"})
}
