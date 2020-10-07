package flags

type CommonOptions struct {
	Debug    bool
	Hosts    []string
	LogLevel string
}

// NewCommonOptions returns a new CommonOptions
func NewCommonOptions() *CommonOptions {
	return &CommonOptions{}
}
