package model

type MigrateReqOpts struct {
	Src Address
	Dst Address
	//Container     string
	ServiceID     string // of worker
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}

type MigrateOpts struct {
	Address
	//Container     string
	ServiceID     string
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}
