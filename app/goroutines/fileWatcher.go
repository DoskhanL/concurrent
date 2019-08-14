package goroutines

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/doskhanl/concurrent/app/model"
)

// ExecFileWatcher function
func ExecFileWatcher(wd *string) {
	fmt.Println("Execution file watcher. Watch for file")
	/*
		files, err := ioutil.ReadDir(*wd + "/invoices")

		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			fmt.Println(f.Name())
		}
	*/

	sourceFolder := *wd + "/source"

	for {
		d, err := os.Open(sourceFolder)
		if err != nil {
			log.Print(fmt.Sprintf("Error while opening source folder: %s", err.Error()))
		}
		fileInfos, err := d.Readdir(-1)
		if err != nil {
			log.Print(fmt.Sprintf("Error while reading directory file infos: %s", err.Error()))
		}
		for _, fi := range fileInfos {
			filePath := sourceFolder + "/" + fi.Name()
			f, err := os.Open(filePath)
			if err != nil {
				log.Print(fmt.Sprintf("Error while opening %s: %s", fi.Name(), err.Error()))
			}

			data, err := ioutil.ReadAll(f)
			f.Close()
			err = os.Remove(filePath)
			if err != nil {
				log.Print(fmt.Sprintf("Error while removing %s from %s: %s", fi.Name(), filePath, err.Error()))
			}
			// Calling anonymous function for reading data from file asynchronously
			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, err := reader.ReadAll()
				if err != nil {
					log.Print(fmt.Sprintf("Error while reading csv file: %s", err.Error()))
				}

				// Operating with csv file data
				for _, r := range records {
					invoice := new(model.Invoice)
					if len(r) > 3 {
						invoice.Number = r[0]

						amount, err := strconv.ParseFloat(r[1], 64)
						if err != nil {
							log.Print(fmt.Sprintf("Error while parsing amount: %s", err.Error()))
						}
						invoice.Amount = float32(amount)

						purchaseOrderNumber, err := strconv.Atoi(r[2])
						if err != nil {
							log.Print(fmt.Sprintf("Error while parsing purchase order number order: %s", err.Error()))
						}
						invoice.PurchaseOrderNumber = purchaseOrderNumber

						unixTime, err := strconv.ParseInt(r[3], 10, 64)
						if err != nil {
							log.Print(fmt.Sprintf("Error while parsing time order: %s", err.Error()))
						}
						invoice.InvoiceDate = time.Unix(unixTime, 0)

						fmt.Println(fmt.Sprintf("Received invoice '%v' for $%.2f and submitted",
							invoice.Number, invoice.Amount))
					}

				}
			}(string(data))
		}
	}
}
