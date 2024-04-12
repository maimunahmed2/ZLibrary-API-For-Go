package zlibrary

import (
	"errors"
	"net/url"
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

func (z *ZLibrary) Init()(map[string]interface{}, error) {
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
	
	return map[string]interface{}{"success": 1}, nil
	// return z.Login(email, password)
}

func (z ZLibrary) Login(email string, password string)(map[string]interface{}, error) {
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

	res, err := utils.MakePostRequest(z.domain+"/eapi/user/login", formData, z.cookies)
	if err != nil {
		return nil, err
	}

	if errVal, ok := res["error"]; ok {
		errMsg := strings.ToLower(errVal.(string))
		return nil, errors.New(errMsg)
	}

	z.cookies["remix_userid"] = strconv.FormatFloat(res["user"].(map[string]interface{})["id"].(float64), 'f', -1, 64)
	z.cookies["remix_userkey"] = res["user"].(map[string]interface{})["remix_userkey"].(string)

	return res, nil
}

func (z ZLibrary) GetProfile()(map[string]interface{}, error) {
	return utils.MakeGetRequest(z.domain+"/eapi/user/profile", z.cookies)
}

func (z ZLibrary) GetSimilar(id string, hash string)(map[string]interface{}, error) {
	return utils.MakeGetRequest(z.domain+"/eapi/book/"+id+"/"+hash+"/similar", z.cookies)
}