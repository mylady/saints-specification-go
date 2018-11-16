package util

import (
	"time"
)

//TickerFunc :ticker function
type TickerFunc func()

//SimpleTicker :time.Ticker wrapper
type SimpleTicker struct {
	ticker *time.Ticker
	dur    time.Duration
	fun    TickerFunc
}

//NewSimpleTicker :create new timer
func NewSimpleTicker(d time.Duration, f TickerFunc) *SimpleTicker {
	return &SimpleTicker{
		dur: d,
		fun: f,
	}
}

//Start :start timer
func (st *SimpleTicker) Start() {
	st.ticker = time.NewTicker(st.dur)
	go func() {
		for range st.ticker.C {
			st.fun()
		}
	}()
}

//Stop :stop timer
func (st *SimpleTicker) Stop() {
	st.ticker.Stop()
}
