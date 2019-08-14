package channels

import "fmt"

// SimpleChan func
func SimpleChan() {
	ch := make(chan string, 1)
	ch <- "Hello"
	fmt.Println(<-ch)
	close(ch)
}
