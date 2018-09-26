package model

import (
	"time"
)

type WorkRequest struct {
	Name  string
	Delay time.Duration
	File  string
}

//COLLECTOR
// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 100)
