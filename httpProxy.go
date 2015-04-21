package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"html"

	"h12.me/socks"
)

func prepareProxyClient() *http.Client {
	dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, "127.0.0.1:9050")
	transport := &http.Transport{Dial: dialSocksProxy}
	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(5 * time.Second)}
}

func httpGet(httpClient *http.Client, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.8 [en] (FreeBSD; U)")
	resp, err = httpClient.Do(req)
	return
}

func httpGetBody(httpClient *http.Client, url string) (body string, err error) {
	resp, err := httpGet(httpClient, url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyb, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<18))
	if err != nil {
		return "", err
	}
	body = string(bodyb)
	return
}

func httpGetTitle(httpClient *http.Client, url string) (title string, err error) {
	body, err := httpGetBody(httpClient, url)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile("<title>(.{1,128})<\\/title>")
	titleMatch := re.FindStringSubmatch(body)
	if len(titleMatch) > 0 {
		return html.UnescapeString(titleMatch[1]), nil
	}
	return "", nil
}
