package model

type MigrateReqOpts struct {
	Src           Address // migration src
	Dst           Address // migration destination
	ServiceID     string  // of worker
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}

type MigrateOpts struct {
	Address
	ServiceID     string
	ProxyService  string
	CheckpointID  string
	CheckpointDir string
}
