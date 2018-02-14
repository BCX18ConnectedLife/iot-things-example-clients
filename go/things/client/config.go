package client

import "time"

type Configuration struct {
	Proxy         string
	SkipSslVerify bool
	WebSocketPingInterval time.Time
}
