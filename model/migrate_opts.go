package model

type MigrateOpts struct {
	Src           Address
	Dst           Address
	Container     string
	CheckpointID  string
	CheckpointDir string
	//DestIP        string
	//DestPort      string
}
