package zlibrary

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/maimunahmed2/ZLibrary-API-For-Go/utils"
)

type ZLibrary struct {
	email        string
	password     string
	domain       string
	isLoggedIn   bool
	headers      map[string]string
	cookies      map[string]string
}

func (z *ZLibrary) Init()(map[string] interface{}, error) {
	z.domain = "https://singlelogin.se"
	z.isLoggedIn = false
	z.headers = map[string]string{
		"Content-Type":    "application/x-www-form-urlencoded",
		"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"accept-language": "en-US,en;q=0.9",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
	}
	z.cookies = map[string]string{
		"siteLanguageV2": "en",
	}
	
	return map[string] interface{}{"success": 1}, nil
	// return z.Login(email, password)
}

func (z ZLibrary) LoginWithCredentials(email string, password string)(map[string]interface{}, error) {
	formData := url.Values{"email": {email}, "password": {password}}
	// Alternatively
	// formData := url.Values{}
    // formData.Set("key", "value")
	// OR
	// formData := map[string]string{"email": email, "password": password}
	// OR
	// formData := make(map[string]string)
    // formData["email"] = email
    // formData["password"] = password

	res, err := utils.MakePostRequest(z.domain+"/eapi/user/login", formData, z.headers, z.cookies)
	if err != nil {
		return nil, err
	}

	if errVal, ok := res["error"]; ok {
		errMsg := strings.ToLower(errVal.(string))
		return nil, errors.New(errMsg)
	}

	z.cookies["remix_userid"] = strconv.FormatFloat(res["user"].(map[string]interface{})["id"].(float64), 'f', -1, 64)
	z.cookies["remix_userkey"] = res["user"].(map[string] interface{})["remix_userkey"].(string)

	return res, nil
}

func (z ZLibrary) LoginWithToken(remix_userid string, remix_userkey string)(map[string]interface{}, error) {
	z.cookies["remix_userid"] = remix_userid
	z.cookies["remix_userkey"] = remix_userkey

	return utils.MakeGetRequest(z.domain+"/eapi/user/profile", z.headers, z.cookies)
}

func (z ZLibrary) GetProfile()(map[string] interface{}, error) {
	return utils.MakeGetRequest(z.domain+"/eapi/user/profile", z.headers, z.cookies)
}

func (z ZLibrary) GetSimilar(id string, hash string)(map[string] interface{}, error) {
	return utils.MakeGetRequest(z.domain+"/eapi/book/"+id+"/"+hash+"/similar", z.headers, z.cookies)
}

func (z ZLibrary) GetRecommended()(map[string] interface{}, error) {
	return utils.MakeGetRequest(z.domain+"/eapi/user/book/recommended", z.headers, z.cookies)
}

func (z ZLibrary) GetMostPopular()(map[string] interface{}, error) {
	return utils.MakeGetRequest(z.domain+"/eapi/book/most-popular", z.headers, z.cookies)
}

func (z ZLibrary) GetRecently()(map[string] interface{}, error) {
	return utils.MakeGetRequest(z.domain+"/eapi/book/recently", z.headers, z.cookies)
}

func (z ZLibrary) GetDownloads()(map[string] interface{}, error) {
	return utils.MakeGetRequest(z.domain+"/eapi/user/book/downloaded", z.headers, z.cookies)
}

//TODO Make it a map and use for loop to simplify this func.
// TODO Languages not working
func (z ZLibrary) Search(searchTerm string, yearFrom *int, yearTo *int, languages *[]string, extensions *[]string, order *string, page *int, limit *int)(map[string]interface{}, error) {
	data := url.Values{
		"message": {searchTerm},
	}
	if yearFrom != nil {
		data.Add("yearFrom", strconv.Itoa(*yearFrom))
	}
	if yearTo != nil {
		data.Add("yearTo", strconv.Itoa(*yearTo))
	}
	if order != nil {
		data.Add("order", *order)
	}
	if page != nil {
		data.Add("page", strconv.Itoa(*page))
	}
	if limit != nil {
		data.Add("limit", strconv.Itoa(*limit))
	}
	if languages != nil {
		for _, language := range *languages {
			data.Add("languages[]", language)
		}
	}
	if extensions != nil {
		for _, extension := range *extensions {
			data.Add("extensions[]", extension)
		}
	}

	return utils.MakePostRequest(z.domain+"/eapi/book/search", data, z.headers, z.cookies)
}

func (z ZLibrary) GetExpirableDownloadLink(bookId string, bookHash string, filename string)(*string, error) {
	res, err := utils.MakeGetRequest(z.domain+"/eapi/book/"+bookId+"/"+bookHash+"/file", z.headers, z.cookies)
	if err != nil {
		return nil, err
	}

	if res["file"].(map[string] interface{})["allowDownload"] != true {
		return nil, errors.New("this file doesn't allow downloads")
	}
	downloadLink := res["file"].(map[string] interface{})["downloadLink"].(string)

	parsedURL, err := url.Parse(downloadLink)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	originalFilename := parsedURL.Query().Get("filename")
	extension := filepath.Ext(originalFilename)

	newFilename := filename + extension
	query := parsedURL.Query()
	query.Set("filename", newFilename)
	parsedURL.RawQuery = query.Encode()

	downloadLink = parsedURL.String()
	return &downloadLink, nil
}

func (z ZLibrary) GetBookData(bookId string, bookHash string)(map[string]interface{}, error) {
	return utils.MakeGetRequest(z.domain+"/eapi/book/"+bookId+"/"+bookHash, z.headers, z.cookies)
}