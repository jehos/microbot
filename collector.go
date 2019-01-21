package microbot

import "runtime/pprof"

func GetGoroutines() {
	pprof.Lookup("goroutine")
}
