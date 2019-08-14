package channels

import (
	"fmt"
	"strings"
)

// BufferedChan func
func BufferedChan() {
	phrase := "There are times that try men's souls.\n"

	words := strings.Split(phrase, " ")

	// Providing buffer for channels
	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}
	close(ch)
	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
	// It will throw an error, because ch is closed
	ch <- "Hello my friend"
}
