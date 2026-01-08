package server

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/runoneall/pgoipc/ipcserver"
)

func Start() {
	fmt.Println("正在启动服务端...")
	ipcserver.Serv("portlink-server", func(conn net.Conn) {

		// 解析请求
		r := bufio.NewReader(conn)
		raw, err := r.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				fmt.Println("不能读取客户端请求", err)
			}
			return
		}

		job, err := parseJSON(raw)
		if err != nil {
			fmt.Println("不能解析客户端请求", err)
			return
		}

		// 处理任务
		switch job.Name {

		case "list":
			actionList(conn)

		case "forward":
			actionForward(conn, job)

		case "stop":
			actionStop(conn, job)

		default:
			fmt.Fprint(conn, "未知命令")

		}
	})
}
