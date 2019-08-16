package couples

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/doskhanl/concurrent/app/model"
)

// ExecETL extract, transform and load
func ExecETL(wd *string) {
	fmt.Println("Executing extract, transform and load")
	start := time.Now()

	orders := extractOrders(wd)
	orders = transform(wd, orders)
	load(wd, orders)
	fmt.Println(time.Since(start))
	//fmt.Scanln()
}

// extractOrders data from file
func extractOrders(wd *string) []*model.Order {

	f, err := os.Open(*wd + "/data/orders.txt")
	if err != nil {
		log.Printf("Error while opening file: %s\n", err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)

	result := []*model.Order{}
	// read contents of a file
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(model.Order)
		if (len(record)) > 2 {
			customerNumber, err := strconv.Atoi(record[0])
			if err != nil {
				log.Printf("Error while converting customer number: %s\n", err.Error())
			}
			order.CustomerNumber = customerNumber

			order.PartNumber = record[1]

			quantity, err := strconv.Atoi(record[2])
			if err != nil {
				log.Printf("Error while converting quantity: %s\n", err.Error())
			}
			order.Quantity = quantity

			result = append(result, order)
		}
	}

	return result
}

//transform transforming order with product file
func transform(wd *string, orders []*model.Order) []*model.Order {
	f, err := os.Open(*wd + "/data/productList.txt")
	if err != nil {
		log.Printf("Error while opening product list %s\n", err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	productList := make(map[string]*model.Product)
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		if len(record) > 2 {
			product := new(model.Product)
			product.PartNumber = record[0]

			unitCost, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				log.Printf("Error while converting unit cost %s\n", err.Error())
			}
			product.UnitCost = unitCost

			unitPrice, err := strconv.ParseFloat(record[2], 64)
			if err != nil {
				log.Printf("Error while converting unit price %s\n", err.Error())
			}
			product.UnitPrice = unitPrice

			productList[product.PartNumber] = product
		}
	}

	for _, order := range orders {
		// Simulating web service or data base retrieving
		time.Sleep(3 * time.Millisecond)
		product := productList[order.PartNumber]
		if product != nil {
			order.UnitCost = product.UnitCost
			order.UnitPrice = product.UnitPrice
		}
	}

	return orders
}

// Load function
func load(wd *string, orders []*model.Order) {
	// creating destination file for writing data to file
	// simulation of data base
	f, err := os.Create(*wd + "/data/dest.txt")
	defer f.Close()
	if err != nil {
		log.Printf("Error while creating destination file %s\n", err.Error())
	}
	fmt.Fprintf(f, "%20s %15s %12s %12s %15s %15s\n",
		"Part Number", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")
	for _, o := range orders {
		time.Sleep(1 * time.Millisecond)
		fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n",
			o.PartNumber,
			o.Quantity,
			o.UnitCost,
			o.UnitPrice,
			o.UnitCost*float64(o.Quantity),
			o.UnitPrice*float64(o.Quantity),
		)
	}
}
