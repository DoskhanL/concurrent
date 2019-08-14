package couples

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

//MutexLock function mutual exclusion lock
func MutexLock() {
	mutex := new(sync.Mutex)
	fmt.Println("Mutual exclusion lock")

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex.Lock()
			go func(a, b int) {
				fmt.Printf("%d + %d = %d\n", a, b, a+b)
				mutex.Unlock()
			}(i, j)
		}
	}
	fmt.Scanln()
}

// MutexLockChan function making like mutex.Lock function
func MutexLockChan() {
	fmt.Println("Mutual exclusion lock by channels")

	logFileName := "./log.txt"
	f, err := os.Create(logFileName)
	if err != nil {
		log.Println(fmt.Sprintf("Error while creating a log file: %s", err.Error()))
		return
	}
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				if err != nil {
					log.Println(fmt.Sprintf("Error while opening a log file: %s", err.Error()))
					break
				}
				logTime := time.Now().Format(time.RFC3339)
				_, err = f.WriteString(logTime + " - " + msg)
				if err != nil {
					log.Println(fmt.Sprintf("Error while writing a string to file: %s", err.Error()))
					break
				}
				err = f.Sync()
				if err != nil {
					log.Println(fmt.Sprintf("Error while sync to file: %s", err.Error()))
					break
				}
				err = f.Close()
				if err != nil {
					log.Println(fmt.Sprintf("Error while closing file: %s", err.Error()))
					break
				}
			} else {
				fmt.Print("Channel closed! \n")
				break
			}
		}
	}()

	// Iteration faster than writing a file
	mutex := make(chan bool, 1)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func(a, b int) {
				msg := fmt.Sprintf("%d + %d = %d\n", a, b, a+b)
				logCh <- msg
				fmt.Print(msg)
				<-mutex
			}(i, j)
		}
	}
	fmt.Scanln()
}
