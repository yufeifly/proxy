package model

type MigrateReqOpts struct {
	Src Address // migration src
	Dst Address // migration destination
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
