package utils

import (
	"net/http"
	"net/url"
)

func HttpGet(params map[string]string, requrl string) (err error) {
	Url, err := url.Parse(requrl)
	urlparams := url.Values{}
	for k, v := range params {
		urlparams.Set(k, v)
	}
	Url.RawQuery = urlparams.Encode()
	_, err = http.Get(Url.String())
	return
}
