package command

import (
	"github.com/spf13/cobra"
	"github.com/yufeifly/proxy/client"
	"io"
)

type ProxyCli struct {
	err    io.Writer
	client client.APIClient
}

func NewProxyCli(err io.Writer) *ProxyCli {
	return &ProxyCli{err: err}
}

// Client return the api Client
func (cli *ProxyCli) Client() client.APIClient {
	return cli.client
}

// ShowHelp shows the command help.
func (cli *ProxyCli) ShowHelp(cmd *cobra.Command, args []string) error {
	cmd.SetOut(cli.err)
	cmd.HelpFunc()(cmd, args)
	return nil
}
