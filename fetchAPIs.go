package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
    "github.com/mmcdole/gofeed"
    "sort"
)

type FinnhubNewsItem struct {
    Category string `json:"category"`
    DateTime int64  `json:"datetime"`
    Headline string `json:"headline"`
    Image    string `json:"image"`
    Related  string `json:"related"`
    Source   string `json:"source"`
    Summary  string `json:"summary"`
    URL      string `json:"url"`
}

type NewsItem struct {
    Item        *gofeed.Item
    SourceTitle string
}

func fetchIndexData(time_range string) {

    for _, index := range GLOBAL_INDICES {

        url := fmt.Sprintf("https://global-market-indices-data.p.rapidapi.com/v1/index_price_change?index=%s&period=%s", index, time_range)

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            log.Printf("Error creating request for index %s: %v", index, err)
            return
        }

        req.Header.Add("x-rapidapi-key", GLOBAL_IDX_API_KEY)
        req.Header.Add("x-rapidapi-host", "global-market-indices-data.p.rapidapi.com")

        res, err := http.DefaultClient.Do(req)
        if err != nil {
            log.Printf("Error fetching data for index %s: %v", index, err)
            return
        }
        defer res.Body.Close()

        body, err := io.ReadAll(res.Body)
        if err != nil {
            log.Printf("Error reading response for index %s: %v", index, err)
            return
        }

        fmt.Printf("Response for %s:\n", index)
        fmt.Println(string(body))
        fmt.Println("----------------------------------------------------")
    }
}

func fetchMarketNews() ([]FinnhubNewsItem, error) {
    url := fmt.Sprintf("https://finnhub.io/api/v1/news?category=general&token=%s", NEWS_API_KEY)

    res, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making HTTP request: %v", err)
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        bodyBytes, _ := io.ReadAll(res.Body)
        return nil, fmt.Errorf("non-200 response: %d - %s", res.StatusCode, string(bodyBytes))
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %v", err)
    }

    var newsItems []FinnhubNewsItem
    err = json.Unmarshal(body, &newsItems)
    if err != nil {
        return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
    }

    return newsItems, nil
}

func formatNews(newsItems []FinnhubNewsItem) string {
    formattedNews := "Market Bulletin\n"
    specialChars := []string{"•", "➤", "★", "✔", "▶", "♦", "→", "✶", "✓"}

    // Limit the number of news items
    maxItems := 10
    if len(newsItems) < maxItems {
        maxItems = len(newsItems)
    }

    for i := 0; i < maxItems; i++ {
        item := newsItems[i]
        // Choose a random special character
        prefix := specialChars[i%len(specialChars)]
        // Format the date
        date := time.Unix(item.DateTime, 0).Format("02 Jan 2006")
        // Append to the formatted news
        formattedNews += fmt.Sprintf("%s %s (Source: %s, Date: %s)\n", prefix, item.Headline, item.Source, date)
    }
    return formattedNews
}

func fetchMarketNewsFromMultipleRSS() ([]NewsItem, error) {
    fp := gofeed.NewParser()
    var allItems []NewsItem

    for _, feedURL := range FEEDS {
        feed, err := fp.ParseURL(feedURL)
        if err != nil {
            log.Printf("Error parsing RSS feed %s: %v", feedURL, err)
            continue
        }
        for _, item := range feed.Items {
            allItems = append(allItems, NewsItem{
                Item:        item,
                SourceTitle: feed.Title,
            })
        }
    }

    return allItems, nil
}

func formatRSSNews(items []NewsItem) string {
    formattedNews := "Market Bulletin\n"
    specialChars := []string{"•", "➤", "★", "✔", "▶", "♦", "→", "✶", "✓"}

    // Sort items by published date (optional)
    sort.SliceStable(items, func(i, j int) bool {
        return items[i].Item.PublishedParsed.After(*items[j].Item.PublishedParsed)
    })

    // Limit the number of news items
    maxItems := 10
    if len(items) < maxItems {
        maxItems = len(items)
    }

    for i := 0; i < maxItems; i++ {
        item := items[i]
        prefix := specialChars[i%len(specialChars)]
        date := item.Item.PublishedParsed.Format("02 Jan 2006")
        newsText := fmt.Sprintf("%s %s : %s (Source: %s, Date: %s)", prefix, item.Item.Title, item.Item.Description, item.SourceTitle, date)
        formattedNews += newsText + "\n"
    }
    return formattedNews
}

