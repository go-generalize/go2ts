package main

import (
	"time"
)

type Status string

const (
	StatusOK      Status = "OK"
	StatusFailure Status = "Failure"
)

type Param struct {
	Status    Status
	Version   int
	Action    string
	CreatedAt time.Time
}
