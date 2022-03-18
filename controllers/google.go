package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sumitroajiprabowo/go-oauth2/config"
	"golang.org/x/oauth2"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request) {

	// Set the config
	config := config.SetupConfig()

	// Get the url
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)

	// Redirect the user
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)

	// Print the url
	fmt.Println(url)

}

func GoogleCallback(res http.ResponseWriter, req *http.Request) {

	// Set the config
	config := config.SetupConfig()

	// Get the code
	code := req.URL.Query().Get("code")

	// Get the token
	token, err := config.Exchange(oauth2.NoContext, code)

	// Check the error
	if err != nil {
		fmt.Println(err)
	}

	// Print the token
	fmt.Println(token)

	// Get the user info
	client := config.Client(oauth2.NoContext, token)

	// Get the user info
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	// Check the error
	if err != nil {
		fmt.Println(err)
	}

	// Read the body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// Check the error
	if err != nil {
		fmt.Println(err)
	}

	// Print the body to console
	fmt.Println(string(body))

	// Print body to browser
	fmt.Fprintf(res, string(body))

}
