package main

import (
	"portlink/client"
	"portlink/cmd"
	"portlink/server"
)

func main() {
	cmd.InitCli()
	cmd.Parse()

	switch *cmd.IsStartServer {

	case false:
		client.DoJob()

	case true:
		server.Start()

	}
}
