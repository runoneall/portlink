package server

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func (f *forward) Udp() error {
	remoteAddrStr := net.JoinHostPort(f.RH, strconv.Itoa(f.RP))

	remoteAddr, err := net.ResolveUDPAddr("udp", remoteAddrStr)
	if err != nil {
		return fmt.Errorf("解析远程UDP地址失败: %w", err)
	}

	localConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(f.LH),
		Port: f.LP,
	})

	if err != nil {
		return fmt.Errorf("监听UDP端口失败: %w", err)
	}
	defer localConn.Close()

	var wg sync.WaitGroup
	var connMap sync.Map

	go func() {
		<-f.stopChan
		localConn.Close()

		connMap.Range(func(key, value interface{}) bool {
			connMap.Delete(key)
			return true
		})
	}()

	buf := make([]byte, 65535)
	for {
		n, clientAddr, err := localConn.ReadFromUDP(buf)
		if err != nil {
			if _, ok := err.(net.Error); ok {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			break
		}

		wg.Add(1)
		go func(clientAddr *net.UDPAddr, data []byte) {
			defer wg.Done()

			clientKey := clientAddr.String()

			remoteConnIface, ok := connMap.Load(clientKey)
			if !ok {
				remoteConn, err := net.DialUDP("udp", nil, remoteAddr)
				if err != nil {
					return
				}

				connMap.Store(clientKey, remoteConn)

				go func() {
					remoteBuf := make([]byte, 65535)
					for {
						n, err := remoteConn.Read(remoteBuf)
						if err != nil {
							connMap.Delete(clientKey)
							remoteConn.Close()
							return
						}
						_, err = localConn.WriteToUDP(remoteBuf[:n], clientAddr)
						if err != nil {
							connMap.Delete(clientKey)
							remoteConn.Close()
							return
						}
					}
				}()

				remoteConnIface = remoteConn
			}

			remoteConn := remoteConnIface.(*net.UDPConn)
			_, err := remoteConn.Write(data[:n])
			if err != nil {
				connMap.Delete(clientKey)
				remoteConn.Close()
			}

		}(clientAddr, buf[:n])
	}

	wg.Wait()
	return nil
}
