package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cauchy2384/go-url-shortener/internal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	shortener := internal.NewURLShortener()

	for scanner.Scan() {
		line := scanner.Text()

		var (
			result string
			err error
		)
		if strings.HasPrefix(line, "http") {
			result, err = shortener.ShortenURL(line)
		} else {
			result, err = shortener.RetrieveURL(line)
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
