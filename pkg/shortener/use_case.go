package shortener

import "github.com/cauchy2384/go-url-shortener/pkg/hash"

type URL struct {
	Full  string
	Short string
}

type Service struct {
	repository Repository
}

func NewService() Service {
	return Service{
		repository: NewRepository(),
	}
}

func (s Service) ShortenURL(full string) (short string, err error) {
	if short, err = hash.Hash(full); err != nil {
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

func (s Service) RetrieveURL(short string) (url string, err error) {
	record, err := s.repository.FindByShort(short)
	if err != nil {
		return "", err
	}

	return record.Full, nil
}
