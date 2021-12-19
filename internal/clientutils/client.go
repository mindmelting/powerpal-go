package clientutils

import "net/http"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HttpClient

func init() {
	Client = &http.Client{}
}
