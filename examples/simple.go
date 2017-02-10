package main

import (
	curl "github.com/wz2cool/go-curl"
)

func main() {
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		easy.Setopt(curl.OPT_URL, "http://www.google.com/")
		easy.Perform()
	}
}
