package commands

import (
	"github.com/spf13/cobra"
	"github.com/yufeifly/proxy/cli/command"
	"github.com/yufeifly/proxy/cli/command/container"
	"github.com/yufeifly/proxy/cli/command/migrate"
)

// AddCommands adds all the commands from cli/command to the root command
func AddCommands(cmd *cobra.Command, proxyCli *command.ProxyCli) {
	cmd.AddCommand(
		migrate.NewMigrateCommand(proxyCli),
		container.NewListCommand(proxyCli),
	)
}
