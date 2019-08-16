package couples

import (
	"fmt"

	"github.com/doskhanl/concurrent/app/model"
)

// ExecPromises func
func ExecPromises() {
	fmt.Println("Executing promises block")
	po := new(model.PurchaseOrder)

	po.Value = 42.27
	model.SavePromisePO(po, false).Then(func(obj interface{}) error {
		po := obj.(*model.PurchaseOrder)
		fmt.Printf("Purchase Order saved with ID: %d\n", po.Number)
		return nil
		// checking of error handling second promse block
		//return errors.New("First promise failed")
	}, func(err error) {
		fmt.Printf("Failed to save Purchase Order: %s\n", err.Error())
	}).Then(func(obj interface{}) error {
		fmt.Println("Second promise success")
		return nil
	}, func(err error) {
		fmt.Println("Second promise failed: " + err.Error())
	})
	fmt.Scanln()
}
