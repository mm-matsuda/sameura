package main

import (
	"io/ioutil"
	"bytes"
	"net/http"
)

func GetObsPage(domain string, body []byte) []byte {
	str := []byte("IFRAME")
	iframe := bytes.SplitAfter(body, str)
	deli := []byte("\"")
	s := bytes.SplitAfter(iframe[1], deli)
	url := bytes.Split(s[1], deli)

	obs, _ := http.Get(domain + string(url[0]))
	defer obs.Body.Close()
	b2, _ := ioutil.ReadAll(obs.Body)
	return b2
}
