package cmd

import "os"

func Parse() {
	err := Cli.Execute()

	if len(os.Args) == 1 {
		Cli.Usage()
		os.Exit(0)
	}

	if err != nil {
		Cli.Help()
		os.Exit(1)
	}
}
