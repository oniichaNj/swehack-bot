package main

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)
import "errors"

type GiphyResponse struct {
	Data GiphyObj    `json:"data"`
	Meta GiphyStatus `json:"meta"`
}

type GiphyStatus struct {
	Msg    string `json:"msg"`
	Status uint   `json:"status"`
}

type GiphyObj struct {
	Url string `json:"image_url"`
}

func GetGif(search string) (string, error) {
	var giphy GiphyResponse

	clientPtr := prepareProxyClient()
	xx := url.QueryEscape(search)
	// Giphy's public beta api key..
	resp, err := httpGet(clientPtr, "http://api.giphy.com/v1/gifs/random?api_key=dc6zaTOxFJmzC&tag="+xx)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(data, &giphy); err != nil {
		return "", err
	}

	if giphy.Meta.Status != 200 {
		return "", errors.New("Status code != 200")
	}
	return giphy.Data.Url, nil
}
