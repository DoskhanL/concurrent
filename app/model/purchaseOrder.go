package model

// PurchaseOrder model
type PurchaseOrder struct {
	Number int
	Value  float32
}

// SavePO function
func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number = 1234

	callback <- po
}
