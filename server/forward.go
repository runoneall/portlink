package server

import (
	"encoding/json"
	"fmt"
	"sync"
)

type forward struct {
	stopChan chan bool

	ID string `json:"id"`
	RH string `json:"rh"`
	RP int    `json:"rp"`
	LH string `json:"lh"`
	LP int    `json:"lp"`
}

type forwards struct {
	mu         sync.RWMutex
	forward    []*forward
	forward_id []string
}

var f = &forwards{
	forward: make([]*forward, 0),
}

func fmanager() *forwards {
	return f
}

func (f *forwards) GetAll() ([]byte, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	return json.MarshalIndent(f.forward, "", "  ")
}

func (f *forwards) New(cfg *forward) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.forward = append(f.forward, cfg)
	f.forward_id = append(f.forward_id, cfg.ID)
	cfg.Start()
}

func (f *forwards) Stop(id string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	newForward := make([]*forward, 0, len(f.forward))
	for _, fw := range f.forward {
		if fw.ID == id {
			select {
			case fw.stopChan <- true:
			default:
			}

			close(fw.stopChan)
			continue
		}
		newForward = append(newForward, fw)
	}
	f.forward = newForward

	newForwardID := make([]string, 0, len(f.forward_id))
	for _, fid := range f.forward_id {
		if fid != id {
			newForwardID = append(newForwardID, fid)
		}
	}
	f.forward_id = newForwardID
}

func (f *forward) Start() {
	f.stopChan = make(chan bool)

	go func() {
		err := f.Tcp()
		if err != nil {
			fmt.Println("TCP", f.ID, "错误", err)
		}
		fmanager().Stop(f.ID)
	}()

	go func() {
		err := f.Udp()
		if err != nil {
			fmt.Println("UDP", f.ID, "错误", err)
		}
		fmanager().Stop(f.ID)
	}()
}
