package main

import (
	http "github.com/bogdanfinn/fhttp"
	tlsClient "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func createClient(proxyChannel chan string) tlsClient.HttpClient {
	client, _ := tlsClient.NewHttpClient(tlsClient.NewNoopLogger(), []tlsClient.HttpClientOption{
		tlsClient.WithClientProfile(profiles.Safari_IOS_18_0),
		tlsClient.WithTimeoutSeconds(60),
		tlsClient.WithRandomTLSExtensionOrder(),
		tlsClient.WithTransportOptions(&tlsClient.TransportOptions{
			MaxIdleConns:        1,
			MaxConnsPerHost:     1,
			MaxIdleConnsPerHost: 1,
			ReadBufferSize:      8096,
			WriteBufferSize:     8096,
		})}...)
	if proxyChannel != nil {
		client.SetProxy(<-proxyChannel)
	}
	return client
}

func getSafariHeaders() map[string][]string {
	return http.Header {
		"sec-fetch-dest": {"document"},
		"user-agent": {"Mozilla/5.0 (iPhone; CPU iPhone OS 18_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.5 Mobile/15E148 Safari/604.1"},
		"accept": {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
		"referer": {"https://www.google.com/"},
		"sec-fetch-site": {"same-origin"},
		"sec-fetch-mode": {"navigate"},
		"sec-fetch-user": {"?1"},
		"accept-language": {"en-US,en;q=0.9"},
		"priority": {"u=0, i"},
	}
}

func getInstagramHeaders() map[string][]string {
    return http.Header{
        "accept-language":                       {"en-US;q=1.0"},
        "content-type":                          {"application/x-www-form-urlencoded; charset=UTF-8"},
        "ig-intended-user-id":                   {"0"},
        "sec-fetch-site":                        {"same-origin"},
        "priority":                              {"u=2, i"},
        "user-agent":                            {"Instagram 381.1.2.26.83 (iPhone14,5; iOS 16_3_1; en_US; en; scale=3.00; 1170x2532; 737297623) AppleWebKit/420+"},
        "x-bloks-is-prism-enabled":              {"true"},
        "x-bloks-prism-ax-base-colors-enabled":  {"true"},
        "x-bloks-prism-button-version":          {"INDIGO_PRIMARY_BORDERED_SECONDARY"},
        "x-bloks-prism-colors-enabled":          {"true"},
        "x-bloks-prism-font-enabled":            {"false"},
        "x-bloks-prism-link-colors-enabled":     {"1"},
        "x-bloks-version-id":                    {"89260ab7c284bc53283ddb1870bf272c0c189a1a497762c002b28865952b5415"},
        "x-fb-client-ip":                        {"True"},
        "x-fb-connection-type":                  {"wifi"},
        "x-fb-friendly-name":                    {"api"},
        "x-fb-privacy-context":                  {"4760009080727693"},
        "x-fb-request-analytics-tags":           {"{\"network_tags\":{\"product\":\"124024574287414\",\"purpose\":\"fetch\",\"surface\":\"settings\",\"is_ad\":\"0\",\"request_category\":\"api\",\"retry_attempt\":\"0\"},\"application_tags\":{\"is_app_start_media\":\"0\",\"is_nav_critical\":\"0\"}}"},
        "x-fb-server-cluster":                   {"True"},
        "x-ig-abr-connection-speed-kbps":        {"289"},
        "x-ig-app-id":                           {"124024574287414"},
        "x-ig-app-locale":                       {"en"},
        "x-ig-bandwidth-speed-kbps":             {"4146.000"},
        "x-ig-bloks-serialize-payload":          {"true"},
        "x-ig-capabilities":                     {"36r/F/8="},
        "x-ig-connection-speed":                 {"25kbps"},
        "x-ig-connection-type":                  {"WiFi"},
        "x-ig-device-id":                        {"EB1507B6-8959-46EA-958D-ABFEE4C9AA6D"},
        "x-ig-device-locale":                    {"en-US"},
        "x-ig-family-device-id":                 {"18146D2B-7691-4B12-A5C6-2932CBA2EAC5"},
        "x-ig-mapped-locale":                    {"en_US"},
        "x-ig-timezone-offset":                  {"-18000"},
        "x-mid":                                 {"ZDHqpQAAAAEP2s27NnTYI9cmO1UM"},
        "x-pigeon-rawclienttime":                {"1771558782.004573"},
        "x-pigeon-session-id":                   {"UFS-EDFC2394-F9C7-4434-B4C3-1D82C624FF9C-1"},
        "x-tigon-is-retry":                      {"False"},
        "x-fb-conn-uuid-client":                 {"5b69872b8000e89d8770019e377fa39b"},
        "x-fb-http-engine":                      {"MNS/TCP"},
        "x-fb-rmd":                              {"state=URL_ELIGIBLE"},
        "x-fb-tasos-td-config":                  {"prod_signal:1"},
        http.HeaderOrderKey: {
            "accept-language",
            "content-type",
            "ig-intended-user-id",
            "priority",
            "user-agent",
            "x-bloks-is-prism-enabled",
            "x-bloks-prism-ax-base-colors-enabled",
            "x-bloks-prism-button-version",
            "x-bloks-prism-colors-enabled",
            "x-bloks-prism-font-enabled",
            "x-bloks-prism-link-colors-enabled",
            "x-bloks-version-id",
            "x-fb-client-ip",
            "x-fb-connection-type",
            "x-fb-friendly-name",
            "x-fb-privacy-context",
            "x-fb-request-analytics-tags",
            "x-fb-server-cluster",
            "x-ig-abr-connection-speed-kbps",
            "x-ig-app-id",
            "x-ig-app-locale",
            "x-ig-bandwidth-speed-kbps",
            "x-ig-bloks-serialize-payload",
            "x-ig-capabilities",
            "x-ig-connection-speed",
            "x-ig-connection-type",
            "x-ig-device-id",
            "x-ig-device-locale",
            "x-ig-family-device-id",
            "x-ig-mapped-locale",
            "x-ig-timezone-offset",
            "x-mid",
            "x-pigeon-rawclienttime",
            "x-pigeon-session-id",
            "x-tigon-is-retry",
            "x-fb-conn-uuid-client",
            "x-fb-http-engine",
            "x-fb-rmd",
            "x-fb-tasos-td-config",
        },
    }
}

func getChromeHeaders() map[string][]string {
	return http.Header {
		"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"accept-language":           {"en-US,en;q=0.9"},
		"cache-control":             {"max-age=0"},
		"priority":                  {"u=0, i"},
		"sec-ch-ua":                 {`"Google Chrome";v="149", "Chromium";v="149", "Not)A;Brand";v="24"`},
		"sec-ch-ua-mobile":          {"?0"},
		"sec-ch-ua-platform":        {`"macOS"`},
		"sec-fetch-dest":            {"document"},
		"sec-fetch-mode":            {"navigate"},
		"sec-fetch-site":            {"same-origin"},
		"sec-fetch-user":            {"?1"},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/149.0.0.0 Safari/537.36"},
	}
}

func parseCookies(response *http.Response) map[string]string {
	cookies := make(map[string]string)
	for _, cookie := range response.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}
	return cookies
}