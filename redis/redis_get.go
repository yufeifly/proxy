package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/cusErr"
	"github.com/yufeifly/proxy/scheduler"
)

// Get redis get value by key
func Get(service string, key string) (string, error) {
	// get service
	ser, err := scheduler.DefaultScheduler.GetService(service)
	if err != nil {
		return "", err
	}
	// send get request
	cli := client.Client{}
	val, err := cli.RedisGet(ser, key)
	if err != nil {
		return "", err
	}
	return val, err
}

func doGetKV(cli *redis.Client, key string) (string, error) {
	header := "redis.doGetKV"
	val, err := cli.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", cusErr.ErrNotFound
	} else if err != nil {
		logrus.Errorf("%s, err: %v", header, err)
		return "", err
	} else {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": val,
		}).Info("the (key, value) pair")
	}
	return val, nil
}
