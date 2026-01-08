package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func jobList(conn net.Conn) {
	job := action{
		Name: "list",
	}

	raw, err := json.Marshal(&job)
	if err != nil {
		fmt.Println("不能创建请求", err)
		return
	}

	conn.Write(raw)
	conn.Write([]byte("\n"))

	resp, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println("不能读取响应", err)
		return
	}

	fmt.Println(string(resp))
}
