package main

//func newProxyCommand(proxyCli *command.ProxyCli) *cobra.Command {
//	// debug
//	var author string
//	opts := cliflags.NewClientOptions()
//	var flags *pflag.FlagSet
//
//	cmd := &cobra.Command{
//		Use:   "proxy [OPTIONS] COMMAND [ARG...]",
//		Short: "A self-sufficient runtime for proxy",
//		Args:  noArgs,
//		RunE: func(cmd *cobra.Command, args []string) error {
//			logrus.Infof("author: %v", author)
//			return proxyCli.ShowHelp(cmd, args)
//		},
//		SilenceErrors:    true,
//		SilenceUsage:     true,
//		TraverseChildren: true,
//	}
//	cli.SetupRootCommand(cmd)
//
//	//cmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
//
//	flags = cmd.Flags()
//	flags.BoolVarP(&opts.Version, "version", "v", false, "Print version information and quit")
//
//	commands.AddCommands(cmd, proxyCli)
//
//	return cmd
//}
//
//func main() {
//	stderr := os.Stderr
//	proxyCli := command.NewProxyCli(stderr)
//	cmd := newProxyCommand(proxyCli)
//	if err := cmd.Execute(); err != nil {
//		logrus.Errorf("newProxyCommand err: %v", err)
//		os.Exit(1)
//	}
//}
//
//func noArgs(cmd *cobra.Command, args []string) error {
//	if len(args) == 0 {
//		return nil
//	}
//	return fmt.Errorf(
//		"proxy: '%s' is not a proxy command.\nSee 'proxy --help'", args[0])
//}
