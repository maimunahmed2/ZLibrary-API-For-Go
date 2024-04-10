package zlibrary

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/maimunahmed2/ZLibrary-API-For-Go/types"
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

func (z *ZLibrary) Init(email string, password string)(*http.Response, error) {
	z.email = email
	z.password = password
	z.domain = "singlelogin.se"
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

	if email == "" && password == "" {
		return nil, errors.New("email and password cannot be empty")
	}
	
	return z.Login(email, password)

}

func (z ZLibrary) Login(email string, password string)(*http.Response, error) {
	credentials, err := json.Marshal(types.Login{Email: email, Password: password})
	if err != nil {
		return nil, err
	}
	return utils.MakePostRequest(z.domain+"/eapi/user/login", credentials)
}

func (z ZLibrary) GetProfile()(*http.Response, error) {
	return utils.MakeGetRequest(z.domain+"/eapi/user/profile")
}