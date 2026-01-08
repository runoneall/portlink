package client

import (
	"fmt"
	"net"
	"os"
	"portlink/cmd"

	"github.com/runoneall/pgoipc/ipcclient"
	"github.com/runoneall/pgoipc/ipcdial"
)

type action struct {
	Name   string `json:"name"`
	RH     string `json:"rh"`
	RP     int    `json:"rp"`
	LH     string `json:"lh"`
	LP     int    `json:"lp"`
	StopId string `json:"stop"`
}

func DoJob() {
	if _, err := ipcdial.Dial("portlink-server"); err != nil {
		fmt.Println("无法连接到服务端")
		os.Exit(1)
	}

	ipcclient.Connect("portlink-server", func(conn net.Conn) {
		if *cmd.IsList {
			jobList(conn)
			return
		}

		if *cmd.DoStop != "" {
			jobStop(conn, *cmd.DoStop)
			return
		}

		jobForward(conn)
	})
}
