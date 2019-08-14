package model

import "time"

// Invoice model
type Invoice struct {
	Number              string
	Amount              float32
	PurchaseOrderNumber int
	InvoiceDate         time.Time
}
