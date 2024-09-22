package main

import (
    "os"
)

var GLOBAL_IDX_API_KEY = os.Getenv("IDX_API_KEY")
var NEWS_API_KEY       = os.Getenv("FINNHUB_API_KEY")

var GLOBAL_INDICES = []string{
    "SP500",   // S&P 500
    "DJI",     // Dow Jones Industrial
    // "FTSE100", // FTSE 100 Index
    // "CAC40",   // CAC 40 Index
    // "DAX",     // DAX
    // "N225",    // Nikkei 225
    // "HSI",     // Hang Seng Index
}

var FEEDS = []string{
    "https://www.moneycontrol.com/rss/MCtopnews.xml",
    "https://economictimes.indiatimes.com/markets/rssfeeds/1977021501.cms",
    "https://www.business-standard.com/rss/markets-106.rss",
    "https://www.livemint.com/rss/markets",
}