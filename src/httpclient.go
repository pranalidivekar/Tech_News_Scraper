package httpclient

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

// FetchHTML sends an HTTP GET request to the specified URL and returns the HTML content as a string.
func FetchHTML(url string) string {
    client := &http.Client{
        Timeout: 30 * time.Second, // Set a timeout for the HTTP request
    }

    // Send the GET request
    resp, err := client.Get(url)
    if err != nil {
        fmt.Println("Error fetching URL:", url, err)
        return ""
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return ""
    }

    return string(body) // Return the HTML content as a string
}
