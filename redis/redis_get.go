package redis

import "github.com/yufeifly/proxy/dal"

// Get redis get value by key
func Get(key string) (string, error) {
	val, err := dal.GetKV(key)
	if err != nil {
		return "", err
	}
	return val, err
}
