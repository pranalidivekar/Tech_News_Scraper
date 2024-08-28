package filewriter

import (
    "fmt"
    "os"
    "time"
)

// WriteToFile writes the aggregated content to a file and returns the file path.
func WriteToFile(content string) string {
    // Generate the file name based on the current date
    fileName := fmt.Sprintf("tech_news_%s.txt", time.Now().Format("2006-01-02"))
    
    // Create the file
    file, err := os.Create(fileName)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return ""
    }
    defer file.Close()

    // Write the content to the file
    _, err = file.WriteString(content)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return ""
    }

    fmt.Println("File written successfully:", fileName)
    return fileName // Return the file path
}
