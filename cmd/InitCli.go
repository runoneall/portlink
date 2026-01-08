package cmd

import "github.com/spf13/cobra"

var Cli *cobra.Command

func InitCli() {
	Cli = &cobra.Command{
		Use:   "portlink",
		Short: "================| Port Link |================",
	}

	initCmdArg()
}
