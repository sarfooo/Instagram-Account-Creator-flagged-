package main

import (
	http "github.com/bogdanfinn/fhttp"
	"strings"
	"io"
)

func parseInstagramCookies(response *http.Response) map[string]string {
	cookies := parseCookies(response)
	for _, key := range []string{"datr", "csrftoken", "ig_did"} {
		if _, exist := cookies[key]; !exist {
			return nil
		}
	}
	return cookies
}

func parseHashes(response *http.Response) map[string]string {
	cookies := parseInstagramCookies(response)
	if cookies == nil {
		return nil
	}
	bytes, _ := io.ReadAll(response.Body)
	body := string(bytes)
	if !strings.Contains(body, "&jazoest=") || !strings.Contains(body, `["LSD",[],{"token":"`) || !strings.Contains(body, `mid":{"value":"`) {
		return nil
	}
	return map[string]string {
		"datr": cookies["datr"],
		"csrftoken": cookies["csrftoken"],
		"deviceId": cookies["ig_did"],
		"jazoest": strings.Split(strings.Split(body, `&jazoest=`)[1], `"`)[0],
		"lsd": strings.Split(strings.Split(body, `["LSD",[],{"token":"`)[1], `"`)[0],
		"machineId": strings.Split(strings.Split(body, `mid":{"value":"`)[1], `"`)[0],
	}
}