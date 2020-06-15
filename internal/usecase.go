package internal

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"strings"
)

const shortenLength = 8

type URLShortener struct {
	repository Repository
}

func NewURLShortener() URLShortener {
	return URLShortener{
		repository: NewRepository(),
	}
}

func (s URLShortener) ShortenURL(full string) (short string, err error) {
	if short, err = shorten(full); err != nil {
		return "", err
	}

	if _, err := s.repository.FindByShort(short); err == nil {
		return short, nil
	}

	s.repository.Save(&URL{
		Full:  full,
		Short: short,
	})

	return short, nil
}

func (s URLShortener) RetrieveURL(short string) (url string, err error) {
	record, err := s.repository.FindByShort(short)
	if err != nil {
		return "", err
	}

	return record.Full, nil
}

func shorten(url string) (short string, err error) {
	hashed := sha1.New()
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
