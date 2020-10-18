package model

type MigrateReqOpts struct {
	Src           Address
	Dst           Address
	Container     string
	ServiceID     string // of worker
	CheckpointID  string
	CheckpointDir string
}

type MigrateOpts struct {
	Address
	Container     string
	ServiceID     string
	CheckpointID  string
	CheckpointDir string
}
