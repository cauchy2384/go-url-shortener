package shortener

import "github.com/cauchy2384/go-url-shortener/pkg/hash"

// url is a combination of Long (value) and Short (key) urls.
type url struct {
	Long  string
	Short string
}

// Service provides functionality for URL storing and retrieving.
type Service struct {
	repo repository
}

// NewService returns new service instance with in-memory storage.
func NewService() Service {
	return Service{
		repo: newRepository(),
	}
}

// ShortenURL accepts full URL and returns corresponding short key.
func (s Service) ShortenURL(long string) (short string, err error) {
	if short, err = hash.Hash(long); err != nil {
		return "", err
	}

	if _, err := s.repo.findByShort(short); err == nil {
		return short, nil
	}

	s.repo.Save(&url{
		Long:  long,
		Short: short,
	})

	return short, nil
}

// RetrieveURL returns full url by given short key.
// Returns ErrorURLNotFound if nothing exists.
func (s Service) RetrieveURL(short string) (long string, err error) {
	record, err := s.repo.findByShort(short)
	if err != nil {
		return "", err
	}

	return record.Long, nil
}
