package cli

import (
	"github.com/spf13/cobra"
)

func SetupRootCommand(rootCmd *cobra.Command) {

	/*
		A flag can be 'persistent' meaning that this flag will be available
		to the command it's assigned to as well as every command under that command.
		For global flags, assign a flag as a persistent flag on the root.
	*/
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	rootCmd.PersistentFlags().MarkShorthandDeprecated("help", "please use --help")
}
