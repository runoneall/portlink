package server

import (
	"io"
	"net"
	"strconv"
	"sync"
)

func (f *forward) Tcp() error {
	localAddr := net.JoinHostPort(f.LH, strconv.Itoa(f.LP))
	remoteAddr := net.JoinHostPort(f.RH, strconv.Itoa(f.RP))

	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	var connMap sync.Map

	go func() {
		defer listener.Close()

		go func() {
			<-f.stopChan
			listener.Close()

			connMap.Range(func(key, value interface{}) bool {
				conn := value.(net.Conn)
				conn.Close()
				connMap.Delete(key)
				return true
			})
		}()

		for {
			localConn, err := listener.Accept()
			if err != nil {
				break
			}

			wg.Add(1)

			go func() {
				defer wg.Done()
				defer localConn.Close()

				remoteConn, err := net.Dial("tcp", remoteAddr)
				if err != nil {
					return
				}
				defer remoteConn.Close()

				connMap.Store(localConn, localConn)
				defer connMap.Delete(localConn)

				connMap.Store(remoteConn, remoteConn)
				defer connMap.Delete(remoteConn)

				errChan := make(chan error, 2)

				go func() {
					_, err := io.Copy(remoteConn, localConn)
					errChan <- err
				}()

				go func() {
					_, err := io.Copy(localConn, remoteConn)
					errChan <- err
				}()

				<-errChan
			}()
		}
	}()

	wg.Wait()
	return nil
}
