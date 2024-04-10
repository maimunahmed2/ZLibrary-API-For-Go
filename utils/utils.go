package utils

import (
	"bytes"
	"fmt"
	"net/http"
)

func MakePostRequest(url string, body []byte) {
    // URL endpoint to send the POST request
    // url := "https://jsonplaceholder.typicode.com/posts"

    // Create a Post struct instance with sample data
    // post := Post{
    //     UserID: 1,
    //     Title:  "Sample Title",
    //     Body:   "This is the body of the sample post.",
    // }
    //Alternatively
    // post := map[string]interface{}{
	// 	"UserID": 1,
	// 	"Title":  "Sample Title",
	// 	"Body":   "This is the body of the sample post.",
	// }
    //Or
    // post := struct {
    //     UserID int
    //     Title  string
    //     Body   string
    // }{
    //     UserID: 1,
    //     Title:  "Sample Title",
    //     Body:   "This is the body of the sample post.",
    // }
    // Convert the Post struct to JSON byte slice
    // jsonStr, err := json.Marshal(data)
    // if err != nil {
    //     fmt.Println("Error marshalling JSON:", err)
    //     return
    // }

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

// func MakePostRequest(url string, data string)() {
// 	fmt.Println(data)
// }