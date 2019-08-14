package main

import (
	"log"
	"os"

	"github.com/doskhanl/concurrent/app/goroutines"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// goroutines.Exec()
	//goroutines.ExecAsycWeb()
	goroutines.ExecFileWatcher(&wd)
}
