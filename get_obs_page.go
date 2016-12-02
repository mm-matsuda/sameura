package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func GetObsPage(domain string, body string) string {
	str := "IFRAME"
	iframe := strings.SplitAfter(body, str)
	deli := "\""
	s := strings.SplitAfter(iframe[1], deli)
	url := strings.Split(s[1], deli)

	obs, _ := http.Get(domain + string(url[0]))
	defer obs.Body.Close()
	b2, _ := ioutil.ReadAll(obs.Body)
	return string(b2)
}
