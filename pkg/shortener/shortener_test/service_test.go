package shortener_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cauchy2384/go-url-shortener/pkg/shortener"
)

const (
	mediumURL = "https://medium.com/"
	golangURL = "https://golang.org/"
)

func TestShortenURL(t *testing.T) {
	s := shortener.NewService()

	short, err := s.ShortenURL(mediumURL)
	assert.NoError(t, err)
	assert.NotZero(t, short)
}

func TestShortenURL_SameURLs(t *testing.T) {
	s := shortener.NewService()

	short1, err := s.ShortenURL(mediumURL)
	assert.NoError(t, err)

	short2, err := s.ShortenURL(mediumURL)
	assert.NoError(t, err)

	assert.EqualValues(t, short1, short2)
}

func TestShortenURL_DiffURLs(t *testing.T) {
	s := shortener.NewService()

	short1, err := s.ShortenURL(mediumURL)
	assert.NoError(t, err)

	short2, err := s.ShortenURL(golangURL)
	assert.NoError(t, err)

	assert.NotEqualValues(t, short1, short2)
}

func TestRetrieveURL_Exists(t *testing.T) {
	s := shortener.NewService()

	short, err := s.ShortenURL(mediumURL)
	assert.NoError(t, err)

	url, err := s.RetrieveURL(short)
	assert.NoError(t, err)

	assert.EqualValues(t, url, mediumURL)
}

func TestRetrieveURL_NotExists(t *testing.T) {
	s := shortener.NewService()

	short := "SMTH"

	url, err := s.RetrieveURL(short)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, shortener.ErrorURLNotFound))
	assert.Zero(t, url)
}
