package cmd

var (
	IsStartServer *bool
	RemoteHost    *string
	RemotePort    *int
	LocalHost     *string
	LocalPort     *int
	DoStop        *string
	IsList        *bool
)

func initCmdArg() {
	IsStartServer = Cli.Flags().Bool("s", false, "启动后端服务")
	RemoteHost = Cli.Flags().String("rh", "0.0.0.0", "远程主机")
	RemotePort = Cli.Flags().Int("rp", 0, "远程端口")
	LocalHost = Cli.Flags().String("lh", "0.0.0.0", "本地主机")
	LocalPort = Cli.Flags().Int("lp", 0, "本地端口")
	DoStop = Cli.Flags().String("stop", "", "停止转发")
	IsList = Cli.Flags().Bool("l", false, "列出转发的端口")
}
