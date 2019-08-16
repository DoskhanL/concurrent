package couples

import (
	"fmt"

	"github.com/doskhanl/concurrent/app/model"
)

// ExecCallback function
func ExecCallback() {
	fmt.Println("Executing of callback functions")

	po := new(model.PurchaseOrder)
	po.Value = 42.27

	ch := make(chan *model.PurchaseOrder)

	go model.SavePO(po, ch)

	newPo := <-ch

	fmt.Printf("PO: %v\n", newPo)
}
