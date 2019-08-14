package goroutines

import (
	"fmt"
	"runtime"
	"time"
)

// Exec function
func Exec() {

	godur, _ := time.ParseDuration("10ms")
	runtime.GOMAXPROCS(2)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(fmt.Sprintf("%d - %s", i, "Hello"))
			time.Sleep(godur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(fmt.Sprintf("%d - %s", i, "Go"))
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("3s")
	time.Sleep(dur)
}
