package cmd

import "github.com/spf13/cobra"

var Cli *cobra.Command

func InitCli() {
	Cli = &cobra.Command{
		Use:   "portlink",
		Short: "纯 Golang 编写端口转发工具，用于将远程端口转发到本地",
	}

	initCmdArg()
}
