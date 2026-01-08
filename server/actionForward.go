package server

import (
	"fmt"
	"net"

	"github.com/google/uuid"
)

func actionForward(conn net.Conn, job action) {
	cfg := &forward{
		ID: uuid.New().String(),
		RH: job.RH,
		RP: job.RP,
		LH: job.LH,
		LP: job.LP,
	}

	fmanager().New(cfg)
	fmt.Fprint(conn, "添加成功")
}
