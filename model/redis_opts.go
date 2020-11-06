package model

type RedisGetOpts struct {
	Key       string
	ServiceID string
	Node      Address
}

type RedisSetOpts struct {
	Key, Value string
	ServiceID  string
	Node       Address
}
