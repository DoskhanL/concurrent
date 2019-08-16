package model

// Button model
type Button struct {
	eventListeners map[string][]chan string
}

// MakeButton constructor
func MakeButton() *Button {
	button := new(Button)
	button.eventListeners = make(map[string][]chan string)
	return button
}

// AddEventListener function
func (button *Button) AddEventListener(event string, responseChannel chan string) {
	if _, present := button.eventListeners[event]; present {
		button.eventListeners[event] =
			append(button.eventListeners[event], responseChannel)
	} else {
		button.eventListeners[event] = []chan string{responseChannel}
	}
}

//RemoveEventListener function
func (button *Button) RemoveEventListener(event string, listenerChannel chan string) {
	if _, present := button.eventListeners[event]; present {
		for idx := range button.eventListeners[event] {
			if button.eventListeners[event][idx] == listenerChannel {
				button.eventListeners[event] = append(button.eventListeners[event][:idx],
					button.eventListeners[event][idx+1:]...)
				break
			}
		}
	}
}

// TriggerEvent function
func (button *Button) TriggerEvent(event string, response string) {
	if _, present := button.eventListeners[event]; present {
		for _, handler := range button.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}
