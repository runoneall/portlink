package server

import (
	"fmt"
	"net"
)

func actionList(conn net.Conn) {
	list, err := fmanager().GetAll()

	if err != nil {
		fmt.Fprintln(conn, "加载列表失败", err)
		return
	}

	conn.Write(list)
}
