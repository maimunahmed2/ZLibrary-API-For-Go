package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func MakePostRequest(url string, body []byte)(*http.Response, error) {
    // Create a new HTTP request with POST method and set the JSON payload
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return nil, err
    }

    // Set the request headers
    req.Header.Set("Content-Type", "application/json")

    // Create an HTTP client
    client := &http.Client{}

    // Send the request
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return nil, err
    }

    defer resp.Body.Close()

    // responseBody, err := ioutil.ReadAll(resp.Body)
    // if err != nil {
    //     fmt.Println("Error reading response body:", err)
    //     return nil, err
    // }
    // // Print the response status code
    // fmt.Println("Response Status:", responseBody)
    return resp, nil
}

func MakeGetRequest(url string)(*http.Response, error) {
    // Send the GET request
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return nil, err
    }
    defer resp.Body.Close()

    // Read the response body
    // body, err := io.ReadAll(resp.Body)
    // if err != nil {
    //     fmt.Println("Error reading response:", err)
    //     return nil
    // }

    // // Print the response body
    // fmt.Println("Response Body:", string(body))
    return resp, nil
}


func MarshalJSONData() ([]byte, error) {
    // Data to be sent
    data := struct {
        UserID int    `json:"userId"`
        Title  string `json:"title"`
        Body   string `json:"body"`
    }{
		UserID: 1,
        Title:  "Sample Title",
        Body:   "This is the body of the sample post.",
	}

    // Marshal data into JSON string
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    // Return JSON string
    return jsonData, nil
}

func UnmarshalJSONData(jsonData []byte) (struct {
	UserID int `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}, error) {
    var data struct {
        UserID int `json:"userId"`
        Title  string `json:"title"`
        Body   string `json:"body"`
    }

    // Unmarshal JSON string into struct
    err := json.Unmarshal(jsonData, &data)
    if err != nil {
        return struct {
			UserID int `json:"userId"`
			Title  string `json:"title"`
			Body   string `json:"body"`
		}{}, err
    }

    // Return unmarshaled data
    return data, nil
}