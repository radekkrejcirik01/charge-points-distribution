package controller

import "github.com/radekkrejcirik01/charge-points-distribution/pkg/model/statuses"

type Response struct {
	Status  string            `json:""`
	Message string            `json:",omitempty"`
	Data    statuses.Statuses `json:",omitempty"`
}
