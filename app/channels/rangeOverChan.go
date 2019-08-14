package channels

import (
	"fmt"
	"strings"
)

// RangeOverChan func
func RangeOverChan() {
	fmt.Println("Rangin over channels")

	phrase := "These are times that try men's souls.\n"

	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	close(ch)
	// infinite loop for watching the channel
	for {
		if msg, ok := <-ch; ok {
			fmt.Print(msg + " ")
		} else {
			break
		}
	}
}
