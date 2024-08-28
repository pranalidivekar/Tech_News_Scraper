package parser

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "strings"
)

// Article represents a news article with a title, summary, and link.
type Article struct {
    Title   string
    Summary string
    Link    string
}

// siteSelectors holds the CSS selectors for each site.
var siteSelectors = map[string]map[string]string{
    "techcrunch.com": {
        "article": "article",
        "title":   "h2 a",
        "link":    "h2 a",
        "summary": "p",
    },
    "theverge.com": {
        "article": "div.c-compact-river__entry",
        "title":   "h2.c-entry-box--compact__title a",
        "link":    "h2.c-entry-box--compact__title a",
        "summary": "div.c-entry-summary p",
    },
    "news.ycombinator.com": {
        "article": "tr.athing",
        "title":   "a.storylink",
        "link":    "a.storylink",
        "summary": "", // Hacker News typically doesn't have a summary
    },
    "arstechnica.com": {
        "article": "article",
        "title":   "h2 a",
        "link":    "h2 a",
        "summary": "p.excerpt",
    },
    "wired.com": {
        "article": "li.archive-item-component",
        "title":   "h2.archive-item-component__title a",
        "link":    "h2.archive-item-component__title a",
        "summary": "p.archive-item-component__desc",
    },
}

// ParseHTML parses the provided HTML content and returns a slice of Article structs.
func ParseHTML(html string, site string) []Article {
    var articles []Article

    selectors, ok := siteSelectors[site]
    if !ok {
        fmt.Println("No selectors found for site:", site)
        return articles
    }

    // Load the HTML content into goquery
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
    if err != nil {
        fmt.Println("Error loading HTML:", err)
        return articles
    }

    // Find article elements and extract data using site-specific selectors
    doc.Find(selectors["article"]).Each(func(i int, s *goquery.Selection) {
        title := s.Find(selectors["title"]).Text()
        link, _ := s.Find(selectors["link"]).Attr("href")
        summary := s.Find(selectors["summary"]).Text()

        // Create a new Article and add it to the slice
        articles = append(articles, Article{
            Title:   title,
            Summary: summary,
            Link:    link,
        })
    })

    return articles
}
