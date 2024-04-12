package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func encodeForm(data map[string]string) string {
    values := url.Values{}
    for key, value := range data {
        values.Add(key, value)
    }
    return values.Encode()
}

func MakePostRequest(url string, data url.Values, cookies map[string]string)(map[string]interface{}, error) {
    // formData := url.Values{}
    // formData.Set("key", "value")
    // for key, value := range data {
	// }
    // requestBody := strings.NewReader(formData.Encode())
    // formData := make(map[string]string)
    // formData["email"] = "studywithmaimun@gmail.com"
    // formData["password"] = "maimun123"

    // requestBody := strings.NewReader(encodeForm(formData))
	// for key, value := range data {
	// 	formData.Set(key, value)
	// }
    body := strings.NewReader(data.Encode())

    req, err := http.NewRequest("POST", url, body)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return nil, err
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
    req.Header.Set("accept-language", "en-US,en;q=0.9")
    req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

    for name, value := range cookies {
        req.AddCookie(&http.Cookie{Name: name, Value: value})
    }

    client := &http.Client{}

    res, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return nil, err
    }
    defer res.Body.Close()

    var resBody map[string]interface{}
    err = json.NewDecoder(res.Body).Decode(&resBody)
    if err != nil {
        return nil, err
    }
    return resBody, nil
}

func MakeGetRequest(url string, cookies map[string]string)(map[string]interface{}, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return nil, err
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
    req.Header.Set("accept-language", "en-US,en;q=0.9")
    req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

    for name, value := range cookies {
        req.AddCookie(&http.Cookie{Name: name, Value: value})
    }

    client := &http.Client{}

    res, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return nil, err
    }
    defer res.Body.Close()

    var resBody map[string]interface{}
    err = json.NewDecoder(res.Body).Decode(&resBody)
    if err != nil {
        return nil, err
    }
    return resBody, nil
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