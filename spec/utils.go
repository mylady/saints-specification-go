package spec

import (
	"time"
)

//TickerFunc :ticker function
type TickerFunc func(ticker *SimpleTicker)

//SimpleTicker :time.Ticker wrapper
type SimpleTicker struct {
	ticker *time.Ticker
	dur    time.Duration
	fun    TickerFunc
}

//NewSimpleTicker :create new ticker
func NewSimpleTicker(d time.Duration, f TickerFunc) *SimpleTicker {
	return &SimpleTicker{
		dur: d,
		fun: f,
	}
}

//Start :start ticker
func (st *SimpleTicker) Start() {
	st.ticker = time.NewTicker(st.dur)
	go func() {
		for range st.ticker.C {
			st.fun(st)
		}
	}()
}

//Stop :stop ticker
func (st *SimpleTicker) Stop() {
	st.ticker.Stop()
}

//TimerFunc :timer function
type TimerFunc func(timer *SimpleTimer)

//SimpleTimer :time.Timer wrapper
type SimpleTimer struct {
	timer *time.Timer
	dur   time.Duration
	fun   TimerFunc
}

//NewSimpleTimer :create new timer
func NewSimpleTimer(d time.Duration, f TimerFunc) *SimpleTimer {
	return &SimpleTimer{
		dur: d,
		fun: f,
	}
}

//Start :start timer
func (st *SimpleTimer) Start() {
	st.timer = time.NewTimer(st.dur)
	go func() {
		<-st.timer.C
		st.fun(st)
	}()
}

//Stop :stop timer
func (st *SimpleTimer) Stop() {
	st.timer.Stop()
}
