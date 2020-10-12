package model

type MigrateReqOpts struct {
	Src           Address
	Dst           Address
	Container     string
	CheckpointID  string
	CheckpointDir string
}

type MigrateOpts struct {
	Address
	Container     string
	CheckpointID  string
	CheckpointDir string
}
