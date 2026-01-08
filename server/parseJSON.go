package server

import (
	"encoding/json"
)

type action struct {
	Name   string `json:"name"`
	RH     string `json:"rh"`
	RP     int    `json:"rp"`
	LH     string `json:"lh"`
	LP     int    `json:"lp"`
	StopId string `json:"stop"`
}

func parseJSON(raw []byte) (action, error) {
	var job action

	if err := json.Unmarshal(raw, &job); err != nil {
		return job, err
	}

	return job, nil
}
