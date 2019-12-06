package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

func printdec(encodedValue string) {
	res, _ := url.ParseQuery(encodedValue)
	for k := range res {
		fmt.Println(k)
	}
}

func printenc(plainValue string) {
	t := &url.URL{Path: plainValue}
	fmt.Println(t.String())
}

func main() {
	programName := os.Args[0]

	for _, value := range os.Args[1:] {
		if filepath.Base(programName) == "urldec" {
			printdec(value)
		} else {
			printenc(value)
		}
	}
}
