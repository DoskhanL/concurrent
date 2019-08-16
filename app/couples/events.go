package couples

import (
	"fmt"

	"github.com/doskhanl/concurrent/app/model"
)

// SimulateEvents function
func SimulateEvents() {
	fmt.Println("Simulate events function")
	btn := model.MakeButton()

	handlerOne := make(chan string)
	handlerTwo := make(chan string)

	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)

	go func() {
		for {
			msg := <-handlerOne
			fmt.Println("Handler One: " + msg)
		}
	}()

	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println("Handler Two: " + msg)
		}
	}()

	btn.TriggerEvent("click", "Button clicked!")

	btn.RemoveEventListener("click", handlerTwo)

	btn.TriggerEvent("click", "Button clicked again!")

	fmt.Scanln()
}
