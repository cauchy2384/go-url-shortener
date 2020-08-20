package hash

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"io"
	"strings"
)

const shortenLength = 8

// Hash takes an url and returns its hash trimmed to shortenLength.
func Hash(url string) (hash string, err error) {
	hashed := sha512.New()
	_, err = io.Copy(
		hashed,
		strings.NewReader(url),
	)
	if err != nil {
		return "", err
	}

	encoded := bytes.NewBuffer([]byte{})
	_, err = io.Copy(
		base64.NewEncoder(base64.URLEncoding, encoded),
		bytes.NewReader(hashed.Sum(nil)),
	)
	if err != nil {
		return "", err
	}

	encoded.Truncate(shortenLength)

	return encoded.String(), nil
}
