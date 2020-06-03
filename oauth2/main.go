package main

import (
	"crypto/rand"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

const authorizeURL = "https://www.dropbox.com/oauth2/authorize"
const tokenURL = "https://www.dropbox.com/oauth2/token"
const redirectURL = "http://localhost:8080/callback"
const appID = "kx1y9ttwvd57l69"
const secret = "t6ix1vd1lm26z75"

func main() {

	conf := &oauth2.Config{
		ClientID:     appID,
		ClientSecret: secret,
		Scopes:       []string{},
		RedirectURL:  redirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeURL,
			TokenURL: tokenURL,
		},
	}

	r := gin.Default()

	r.GET("/oauth/dropbox/connect", func(c *gin.Context) {
		state := randomString()
		url := conf.AuthCodeURL(state)
		c.SetCookie("state", state, 3600, "/", "localhost", false, true)
		c.Redirect(302, url)
	})

	r.GET("/callback", func(c *gin.Context) {
		state := c.Query("state")
		setState, err := c.Cookie("state")
		if err != nil {
			c.Status(400)
			return
		}
		if state != setState {
			c.Status(400)
			return
		}
		code := c.Query("code")
		token, err := conf.Exchange(c.Request.Context(), code)
		if err != nil {
			c.Status(400)
			return
		}
		c.String(200, listFolders(token.AccessToken))
	})

	r.Run()
}

func randomString() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func listFiles(token string) string {
	url := "https://api.dropboxapi.com/2/file_requests/list_v2"
	body := `{
		"limit": 1000
	}`
	return makeRequest(token, url, body)
}

func listFolders(token string) string {
	url := "https://api.dropboxapi.com/2/files/list_folder"
	body := `{
		"path": ""
	}`
	return makeRequest(token, url, body)
}

func makeRequest(token, url, body string) string {
	payload := strings.NewReader(body)
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(result)
}
