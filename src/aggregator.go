package aggregator

import (
    "fmt"
    "strings"
    "src/parser"
)

// Aggregate takes a slice of Article structs and formats them into a single string.
func Aggregate(articles []parser.Article) string {
    var sb strings.Builder

    // Write the header
    sb.WriteString("Daily Tech News Digest\n")
    sb.WriteString("======================\n\n")

    // Loop through the articles and format them
    for i, article := range articles {
        sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, article.Title))
        sb.WriteString(fmt.Sprintf("*Summary:* %s\n", article.Summary))
        sb.WriteString(fmt.Sprintf("*Link:* %s\n\n", article.Link))
    }

    return sb.String() // Return the aggregated content as a single string
}

