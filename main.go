package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/doskhanl/concurrent/app/couples"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Working directory is: %s\n", wd)
	// Setting max processors size
	runtime.GOMAXPROCS(4)
	// goroutines.Exec()
	//goroutines.ExecAsycWeb()
	//goroutines.ExecFileWatcher(&wd)
	//channels.SimpleChan()
	//channels.BufferedChan()
	//channels.RangeOverChan()
	//channels.SwitchChan()
	//couples.MutexLock()
	//couples.MutexLockChan()
	//couples.SimulateEvents()
	//couples.ExecCallback()
	//couples.ExecPromises()
	couples.ExecPipeFilter()
}
