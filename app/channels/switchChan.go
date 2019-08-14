package channels

import (
	"fmt"

	"github.com/doskhanl/concurrent/app/model"
)

// SwitchChan function for switching between multiple channels
func SwitchChan() {
	fmt.Println("Switching between multiple channels")

	msgCh := make(chan model.Message, 1)
	errCh := make(chan model.FailedMessage, 1)

	msg := model.Message{
		To:      []string{"frodo@underhill.me"},
		From:    "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe.",
	}

	failedMessage := model.FailedMessage{
		ErrorMessage:    "Message intercepted by a black rider",
		OriginalMessage: model.Message{},
	}

	// Pushing message and failed message object to channels
	msgCh <- msg
	errCh <- failedMessage

	// fmt.Println(<-msgCh)
	// fmt.Println(<-errCh)

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("No message received")
	}

}
