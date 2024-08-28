package emailer

import (
    "fmt"
    "gopkg.in/gomail.v2"
    "os"
)

// SendEmail sends the specified file via email.
func SendEmail(filePath string) {
    if filePath == "" {
        fmt.Println("No file to send.")
        return
    }

    // Email configuration
    sender := "pranalidivekar@gmail.com"
    password := os.Getenv("EMAIL_PASSWORD") // Set this in your environment variables
    recipient := "pranalidivekar21@gmail.com"
    subject := "Daily Tech News Digest"
    body := "Please find attached the daily tech news digest."

    // Set up the email message
    m := gomail.NewMessage()
    m.SetHeader("From", sender)
    m.SetHeader("To", recipient)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)
    m.Attach(filePath)

    // Set up the email server
    d := gomail.NewDialer("smtp.gmail.com", 587, sender, password)

    // Send the email
    if err := d.DialAndSend(m); err != nil {
        fmt.Println("Error sending email:", err)
        return
    }

    fmt.Println("Email sent successfully.")
}
