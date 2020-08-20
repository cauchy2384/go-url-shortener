// Package describes url shortener application with CLI interface
// and in-memory storage.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cauchy2384/go-url-shortener/pkg/shortener"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	urlShortener := shortener.NewService()

	for scanner.Scan() {
		line := scanner.Text()

		var (
			result string
			err    error
		)
		if strings.HasPrefix(line, "http") {
			result, err = urlShortener.ShortenURL(line)
		} else {
			result, err = urlShortener.RetrieveURL(line)
		}
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
