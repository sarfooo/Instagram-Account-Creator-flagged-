package main

import (
	"github.com/bogdanfinn/tls-client/profiles"
	tlsClient "github.com/bogdanfinn/tls-client"
	http "github.com/bogdanfinn/fhttp"
	"strings"
)

func createClient(proxyChannel chan string) tlsClient.HttpClient {
	client, _ := tlsClient.NewHttpClient(tlsClient.NewNoopLogger(), []tlsClient.HttpClientOption {
		tlsClient.WithClientProfile(profiles.Safari_IOS_18_0),
		tlsClient.WithTimeoutSeconds(60),
		tlsClient.WithRandomTLSExtensionOrder(),
		tlsClient.WithNotFollowRedirects(),
		tlsClient.WithTransportOptions(&tlsClient.TransportOptions {
            MaxIdleConns:           1,
            MaxConnsPerHost:        1,
            MaxIdleConnsPerHost:    1,
            ReadBufferSize:         8096,
            WriteBufferSize:        8096,
	})}...)
	if proxyChannel != nil {
		client.SetProxy(<- proxyChannel)
	}
	return client
}

func createRequest(method, url, contentType string, payload *strings.Reader) *http.Request {
	request, _ := http.NewRequest(method, url, payload)
	headers := getHeaders()
    if contentType != "" {
        headers["content-type"] = []string{contentType}
    }
	request.Header = getHeaders()
    return request
}

func getHeaders() map[string][]string {
	return http.Header {
        "sec-fetch-dest": {"document"},
		"user-agent": {"Mozilla/5.0 (iPhone; CPU iPhone OS 18_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.5 Mobile/15E148 Safari/604.1"},
		"accept": {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
		"referer": {"https://www.google.com/"},
		"sec-fetch-site": {"cross-site"},
		"sec-fetch-mode": {"navigate"},
		"sec-fetch-user": {"?1"},
		"accept-language": {"en-US,en;q=0.9"},
		"priority": {"u=0, i"},
	}
}