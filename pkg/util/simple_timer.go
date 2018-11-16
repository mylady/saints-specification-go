package util

import (
	"time"
)

//TimerFunc :timer function
type TimerFunc func()

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
		st.fun()
	}()
}

//Stop :stop timer
func (st *SimpleTimer) Stop() {
	st.timer.Stop()
}
