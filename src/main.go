package main

import (
    "fmt"
    "tech-news-scraper/aggregator.go"
    "tech-news-scraper/email.go"
    "tech-news-scraper/filewriter.go"
    "tech-news-scraper/httpclient.go"
    "tech-news-scraper/parser.go"
)

func main() {
    // Define the URLs to scrape
    urls := []string{
        "https://techcrunch.com",
        "https://www.theverge.com",
        "https://news.ycombinator.com",
        "https://arstechnica.com",
        "https://www.wired.com",
    }

    var allArticles []Article

    // Loop through each URL to fetch and parse the HTML
    for _, url := range urls {
        html := httpclient.FetchHTML(url)         // Fetch HTML content
        articles := parser.ParseHTML(html)        // Parse HTML to extract articles
        allArticles = append(allArticles, articles...) // Aggregate all articles
    }

    // Format the aggregated content
    content := aggregator.Aggregate(allArticles)
    
    // Save the content to a file
    filePath := filewriter.WriteToFile(content)
    
    // Send the file via email
    emailer.SendEmail(filePath)

    fmt.Println("Daily Tech News Scraper completed successfully.")
}
