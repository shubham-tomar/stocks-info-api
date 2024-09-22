package main

import (
    "fmt"
    "log"
    "encoding/json"
)

func main() {

    // fetchIndexData("1Day")
    // fetchIndexData("7Days")

    newsItems, err := fetchMarketNewsFromMultipleRSS()
    if err != nil {
        log.Fatalf("Error fetching market news via RSS feeds: %v", err)
    }

    newsJSON, err := json.Marshal(newsItems)
    if err != nil {
        log.Fatalf("Error marshalling news items to JSON: %v", err)
    }
    fmt.Println(string(newsJSON))

    formattedNews := formatRSSNews(newsItems)
    fmt.Println(formattedNews)
}