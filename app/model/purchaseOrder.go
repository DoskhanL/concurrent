package model

import (
	"errors"
	"time"
)

// PurchaseOrder model
type PurchaseOrder struct {
	Number int
	Value  float64
}

// SavePO function
func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number = 1234

	callback <- po
}

// SavePromisePO with promise function
func SavePromisePO(po *PurchaseOrder, shouldFail bool) *Promise {
	result := new(Promise)
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		// checking timeout for promise
		time.Sleep(2 * time.Second)
		if shouldFail {
			result.failureChannel <- errors.New("Failed to save purchase order")
		} else {
			po.Number = 1234
			result.successChannel <- po
		}
	}()

	return result
}
