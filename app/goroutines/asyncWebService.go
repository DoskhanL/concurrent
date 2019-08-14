package goroutines

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"

	"github.com/doskhanl/concurrent/app/model"
)

// ExecAsycWeb method
func ExecAsycWeb() {
	fmt.Println("Executing asynchronous web service")
	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}
	runtime.GOMAXPROCS(4)
	start := time.Now()
	numComplete := 0
	for _, symbol := range stockSymbols {
		go func(symbol string) {
			resp, err := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/quote?symbol=" + symbol)
			if err == nil {
				defer resp.Body.Close()
			}

			body, _ := ioutil.ReadAll(resp.Body)
			quote := new(model.QuoteResponse)
			xml.Unmarshal(body, &quote)
			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}

	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}

	elapsedTime := time.Since(start)

	fmt.Println(fmt.Sprintf("Execution time: %s", elapsedTime))
}
