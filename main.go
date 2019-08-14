package main

import (
	"fmt"
	"log"
	"os"

	"github.com/doskhanl/concurrent/app/channels"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Working directory is: %s\n", wd)
	// goroutines.Exec()
	//goroutines.ExecAsycWeb()
	//goroutines.ExecFileWatcher(&wd)
	//channels.SimpleChan()
	//channels.BufferedChan()
	channels.RangeOverChan()
}
