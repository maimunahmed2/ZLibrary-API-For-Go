package utils

import (
	"bytes"
	"fmt"
	"net/http"
)

func MakePostRequest(url string, body []byte) {
    // Create a new HTTP request with POST method and set the JSON payload
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // Set the request headers
    req.Header.Set("Content-Type", "application/json")

    // Create an HTTP client
    client := &http.Client{}

    // Send the request
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }

    defer resp.Body.Close()

    // Print the response status code
    fmt.Println("Response Status:", resp.Status)
}