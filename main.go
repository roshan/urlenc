package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	decode = "urldec"
	encode = "urlenc"
)

func printDec(s string) string {
	return mustUnescape(url.QueryUnescape(s))
}

func printEnc(s string) string {
	return url.QueryEscape(s)
}

func mustUnescape(s string, err error) string {
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	name := os.Args[0]

	i := strings.Join(os.Args[1:], " ")

	var fn func(i string) string
	switch filepath.Base(name) {
	case encode:
		fn = printEnc
	case decode:
		fn = printDec
	default:
		panic("unknown program type")
	}

	fmt.Println(fn(i))
}
