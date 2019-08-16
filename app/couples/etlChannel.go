package couples

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/doskhanl/concurrent/app/model"
)

// ExecETLChannel extract, transform and load
func ExecETLChannel(wd *string) {
	fmt.Println("Executing extract, transform and load")
	start := time.Now()

	// Creating channels
	extractChannel := make(chan *model.Order)
	transformChannel := make(chan *model.Order)
	doneChannel := make(chan bool)

	// calling asynchronously and using pipes
	go extractOrderChan(wd, extractChannel)
	go transformOrderChan(wd, extractChannel, transformChannel)
	go loadChan(wd, transformChannel, doneChannel)

	// waiting for completing load process
	<-doneChannel

	fmt.Println(time.Since(start))
}

// extractOrderChan data from file
func extractOrderChan(wd *string, ch chan *model.Order) {

	f, err := os.Open(*wd + "/data/orders.txt")
	if err != nil {
		log.Printf("Error while opening file: %s\n", err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)

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

			ch <- order
		}
	}
	fmt.Println("Extracting channel closed")
	close(ch)
}

//transformOrderChan transforming order with product file
func transformOrderChan(wd *string, extractChannel, transformChannel chan *model.Order) {
	f, err := os.Open(*wd + "/data/productList.txt")
	if err != nil {
		log.Printf("Error while opening product list %s\n", err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Printf("Error while reading all line of file %s\n", err.Error())
	}
	productList := make(map[string]*model.Product)
	for _, record := range records {
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
	messagesCount := 0
	var waitGroup sync.WaitGroup

	for o := range extractChannel {
		messagesCount++
		waitGroup.Add(messagesCount)
		go func(o *model.Order) {
			// Simulating web service or data base retrieving
			time.Sleep(3 * time.Millisecond)
			o.UnitCost = productList[o.PartNumber].UnitCost
			o.UnitPrice = productList[o.PartNumber].UnitPrice
			transformChannel <- o
			waitGroup.Done()
		}(o)
	}
	fmt.Println("Transfer channel closed")
	close(transformChannel)
	waitGroup.Wait()
}

// LoadChan function
func loadChan(wd *string, transformChannel chan *model.Order, doneChannel chan bool) {
	// creating destination file for writing data to file
	// simulation of data base
	f, err := os.Create(*wd + "/data/dest.txt")
	defer f.Close()
	if err != nil {
		log.Printf("Error while creating destination file %s\n", err.Error())
	}
	messagesCount := 0
	fmt.Fprintf(f, "%20s %15s %12s %12s %15s %15s\n",
		"Part Number", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")
	var waitGroup sync.WaitGroup
	for o := range transformChannel {
		messagesCount++
		waitGroup.Add(messagesCount)
		go func(o *model.Order) {
			time.Sleep(1 * time.Millisecond)
			fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n",
				o.PartNumber,
				o.Quantity,
				o.UnitCost,
				o.UnitPrice,
				o.UnitCost*float64(o.Quantity),
				o.UnitPrice*float64(o.Quantity))
			waitGroup.Done()
		}(o)
	}
	doneChannel <- true
	waitGroup.Wait()
}
