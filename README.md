# URL shortener in GO  

[![PkgGoDev](https://pkg.go.dev/badge/github.com/cauchy2384/go-url-shortener)](https://pkg.go.dev/github.com/cauchy2384/go-url-shortener)
[![Go Report Card](https://goreportcard.com/badge/github.com/cauchy2384/go-url-shortener)](https://goreportcard.com/report/github.com/cauchy2384/go-url-shortener)
 [![codecov](https://codecov.io/gh/cauchy2384/go-url-shortener/branch/master/graph/badge.svg)](https://codecov.io/gh/cauchy2384/go-url-shortener)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Test

`go test ./...`

## Build & Run

Command-line utility

`go build ./cmd/go-shortener-cli/`

```
user$ ./go-shortener-cli 
>> http://google.com
<< I0mIVmya
>> I0mIVmya
<< http://google.com
>> asdf
<< url not found for "asdf"
```
