package migrate

//type migrateOpts struct {
//	container     string
//	checkpoint    string
//	checkpointDir string
//	leaveRunning  bool
//	destination   string
//}
//
//// proxy migrate container checkpoint --checkpoint-dir=string --leave-running=bool --dest-url=ip:port
//// NewMigrateCommand
//func NewMigrateCommand(proxyCli *command.ProxyCli) *cobra.Command {
//	var opts migrateOpts
//
//	cmd := &cobra.Command{
//		Use:   "migrate",
//		Short: "Manage a running container to target node",
//		Args:  cli.ExactArgs(2),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			opts.container = args[0]
//			opts.checkpoint = args[1]
//
//			return runMigrate(proxyCli, opts)
//		},
//	}
//	flags := cmd.Flags()
//	flags.Bool("help", false, "Print usage")
//	return cmd
//}
//
//func runMigrate(proxyCli *command.ProxyCli, opts migrateOpts) error {
//	client := proxyCli.Client()
//	migOpts := model.MigrateReqOpts{
//		Container:     opts.container,
//		CheckpointID:  opts.checkpoint,
//		CheckpointDir: opts.checkpointDir,
//	}
//	//err := client.Migrate(context.Background(), migOpts)
//	//if err != nil {
//	//	return err
//	//}
//	return nil
//}
