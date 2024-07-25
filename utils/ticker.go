package utils

import "time"

type Ticker interface {
	C() <-chan time.Time
	Stop()
}

type realTicker struct {
	*time.Ticker
}

func NewRealTicker(d time.Duration) Ticker {
	return &realTicker{
		Ticker: time.NewTicker(d),
	}
}

func (rt realTicker) C() <-chan time.Time {
	return rt.Ticker.C
}

//Implement mock ticket, could do in a separable file
