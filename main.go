package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	// "strings"
)

var apiKey = os.Getenv("ALPHA_VANTAGE_API_KEY")

func getGlobalMarketInfo() {
	gloablSymbolList := []string {"BSE", "SHZ"}

	for _, symbol := range gloablSymbolList {
		fmt.Println(symbol)
		url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", symbol, apiKey)
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println("Error: ", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		fmt.Println(string(body))
	}
}

func main() {
	fmt.Println("api repo init")
	getGlobalMarketInfo()
}