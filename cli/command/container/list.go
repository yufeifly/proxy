package container

//type listOptions struct {
//	quiet   bool
//	size    bool
//	all     bool
//	noTrunc bool
//	nLatest bool
//	last    int
//	format  string
//}
//
//func NewListCommand(proxyCli *command.ProxyCli) *cobra.Command {
//	opts := listOptions{}
//	cmd := &cobra.Command{
//		Use:   "list [OPTIONS]",
//		Short: "List containers",
//		Args:  cli.NoArgs,
//		RunE: func(cmd *cobra.Command, args []string) error {
//			return runList(proxyCli, &opts)
//		},
//	}
//	flags := cmd.Flags()
//	flags.BoolVarP(&opts.quiet, "quiet", "q", false, "Only display numeric IDs")
//	flags.BoolVarP(&opts.size, "size", "s", false, "Display total file sizes")
//	flags.BoolVarP(&opts.all, "all", "a", false, "Show all containers (default shows just running)")
//	//flags.VarP(&opts.filter, "filter", "f", "Filter output based on conditions provided")
//	return cmd
//}
//
//func runList(proxyCli *command.ProxyCli, opts *listOptions) error {
//	fmt.Println("i am list")
//	return nil
//}
