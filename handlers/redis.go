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
	serviceID := c.Query("service")

	val, err := redis.Get(serviceID, key)
	if err != nil {
		logrus.Errorf("%s, err: %v", header, err)
		c.JSON(200, gin.H{"failed: ": err})
	} else {
		c.JSON(200, gin.H{"value: ": val})
	}

}

// Set redis set
func Set(c *gin.Context) {

	key := c.Query("key")
	val := c.Query("value")
	serviceID := c.Query("service")

	// todo validate kv pair

	err := redis.Set(serviceID, key, val)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": val,
		}).Error("set pair failed")
		c.JSON(200, gin.H{"failed: ": err})
	} else {
		c.JSON(200, gin.H{"result": "success"})
	}
}
