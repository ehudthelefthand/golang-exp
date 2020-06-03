package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

const authURL = "https://accounts.google.com/o/oauth2/v2/auth"
const tokenURL = "https://oauth2.googleapis.com/token"
const redirectURL = "http://localhost:8080/callback"
const appID = "[GOOGLE_APP_ID]"
const secret = "[GOOGLE_APP_SECRET]"

// http://localhost:8080/oauth/google/connect

func main() {
	// ctx := context.Background()
	// provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	// if err != nil {
	// 	panic(err)
	// }
	// oidcConfig := &oidc.Config{
	// 	ClientID: appID,
	// }
	// verifier := provider.Verifier(oidcConfig)

	conf := &oauth2.Config{
		ClientID:     appID,
		ClientSecret: secret,
		Scopes:       []string{"openid", "profile", "email"},
		RedirectURL:  redirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
	}

	r := gin.Default()

	r.GET("/oauth/google/connect", func(c *gin.Context) {
		state := randString()
		url := conf.AuthCodeURL(state)
		c.SetCookie("state", state, 3600, "/", "localhost", false, true)
		c.Redirect(302, url)
	})

	r.GET("/callback", func(c *gin.Context) {
		state := c.Query("state")
		code := c.Query("code")
		setState, err := c.Cookie("state")
		if err != nil {
			panic(err)
		}
		if state != setState {
			panic("state not match")
		}
		token, err := conf.Exchange(c.Request.Context(), code)
		if err != nil {
			panic(err)
		}
		rawIDToken, ok := token.Extra("id_token").(string)
		if !ok {
			panic("can not get id token")
		}
		fmt.Println("rawIDToken ===> ", rawIDToken)

		encodedProfile := strings.Split(rawIDToken, ".")[1]
		fmt.Println("encoded profile ===> ", encodedProfile)

		b, err := base64.StdEncoding.DecodeString(encodedProfile)
		if err != nil {
			fmt.Println("base64 decode error: ", err)
		}
		fmt.Println("profile ===> ", string(b))

		c.String(200, string(b))
	})

	r.Run()
}

func randString() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func verify(idToken, appID string) {

}
