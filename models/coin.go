package models

import (
	"io"
	"log"
	"net/http"
)

/*
Coin represent a response of
the marketcap request
*/
type Coin struct {
	Name     string
	Symbol   string
	Rank     string
	PriceUSD string `json:"price_usd"`
	PriceEUR string `json:"price_eur"`
}

func (r *Coin) Do(method, uri string, headers map[string]string, body io.Reader) (*http.Response, *error) {
	req, err :=http.NewRequest(metgit inithod, uri, body)
	if err != nil {
		return nil, &err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req.Response, &err
}

func (r *Coin) Get(uri string, params ...map[string]string) (*http.Response, error) {
	var headers map[string]string
	switch len(params) {
	case 0:
	case 1:
		headers = params[0]
	default:
		log.Panic("too much parameters")
	}
	resp := &http.Response{}
	err := new(error)
	resp, err = r.Do("GET", uri, headers, nil)
	return resp, *err
}