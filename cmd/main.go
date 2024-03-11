package main

import (
	"expvar"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.RFC1123))

	testExpvarPackage()

	// Halte das Programm laufend
	select {}
}

func testExpvarPackage() {
	tick := time.NewTicker(1 * time.Second)
	num_go := expvar.NewInt("runtime.goroutines")
	counters := expvar.NewMap("counters")
	counters.Set("cnt1", new(expvar.Int))
	counters.Set("cnt2", new(expvar.Float))

	go http.ListenAndServe(":7777", nil) // --> http://localhost:7777/debug/vars

	for {
		select {
		case <-tick.C:
			num_go.Set(int64(runtime.NumGoroutine()))
			counters.Add("cnt1", 1)
			counters.AddFloat("cnt2", 1.452)
		}
	}

}
