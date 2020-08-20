package shortener

import "errors"

var (
	// ErrorURLNotFound is returned if given URL is not found in storage.
	ErrorURLNotFound = errors.New("url not found")
)
