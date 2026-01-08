package server

import (
	"fmt"
	"net"
)

func actionStop(conn net.Conn, job action) {
	fmanager().Stop(job.StopId)
	fmt.Fprint(conn, "请求成功")
}
