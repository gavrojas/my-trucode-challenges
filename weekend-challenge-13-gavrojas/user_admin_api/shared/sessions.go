package shared

import "time"

type Session struct {
	Uid        uint
	ExpiryTime time.Time
}

var Sessions = make(map[string]Session)
