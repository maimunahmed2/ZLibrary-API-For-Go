package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func MakePostRequest(url string, body []byte)(*http.Response, error) {
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return nil, err
    }

    defer resp.Body.Close()
    return resp, nil
}

func MakeGetRequest(url string)(*http.Response, error) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return nil, err
    }
    defer resp.Body.Close()
    return resp, nil
}


func MarshalJSONData() ([]byte, error) {
    data := struct {
        UserID int    `json:"userId"`
        Title  string `json:"title"`
        Body   string `json:"body"`
    }{
		UserID: 1,
        Title:  "Sample Title",
        Body:   "This is the body of the sample post.",
	}

    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

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

    err := json.Unmarshal(jsonData, &data)
    if err != nil {
        return struct {
			UserID int `json:"userId"`
			Title  string `json:"title"`
			Body   string `json:"body"`
		}{}, err
    }

    return data, nil
}