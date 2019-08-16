package model

import (
	"errors"
	"time"
)

// Promise model
type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

// Then function
func (promise *Promise) Then(success func(interface{}) error, failure func(error)) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	// Setting timeout for waiting response
	timeout := time.After(1 * time.Second)
	go func() {
		select {
		case obj := <-promise.successChannel:
			newErr := success(obj)
			if newErr == nil {
				result.successChannel <- obj
			} else {
				result.failureChannel <- newErr
			}
		case err := <-promise.failureChannel:
			failure(err)
			result.failureChannel <- err
		case <-timeout:
			err := errors.New("Promise timed out")
			failure(err)
		}
	}()

	return result
}
