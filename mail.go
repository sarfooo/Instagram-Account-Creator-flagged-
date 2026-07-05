package main

import (
	"fmt"
	"io"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	tlsClient "github.com/bogdanfinn/tls-client"
)

type Email struct {
	Email string
	Username string
	Domain string
	Cookie map[string]string
}

func newEmail(client tlsClient.HttpClient) *Email {
	request, _ := http.NewRequest("GET", "https://10minutemail.net/", nil)
	request.Header = getChromeHeaders()

	response, err := client.Do(request)
	if err != nil {
		return nil
	}

	bytes, _ := io.ReadAll(response.Body)
	body := string(bytes)
	if !strings.Contains(body, `data-clipboard-text="`) {
		return nil
	}
	
	fmt.Println("HERE?")
	email := strings.Split(strings.Split(body, `data-clipboard-text="`)[1], `"`)[0]
	emailParts := strings.Split(email, "@")
	return &Email {
		Email: email,
		Username: emailParts[0],
		Domain: emailParts[1],
		Cookie: parseCookies(response),
	}
}

func getCode(client tlsClient.HttpClient) string {
	request, _ := http.NewRequest("GET", "https://10minutemail.net/", nil)
	request.Header = getChromeHeaders()

	response, err := client.Do(request)
	if err != nil {
		return ""
	}

	io.ReadAll(response.Body)
	return ""
}