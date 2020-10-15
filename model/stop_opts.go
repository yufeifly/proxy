package model

type StopOpts struct {
	ContainerID string
	Timeout     string
}

type StopReqOpts struct {
	Address
	StopOpts
}
