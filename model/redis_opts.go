package model

type RedisGetOpts struct {
	Key       string
	ServiceID string
	Node      Address
}

type RedisSetOpts struct {
	Key       string
	Value     string
	ServiceID string
	Node      Address
}
