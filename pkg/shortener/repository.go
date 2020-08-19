package shortener

import (
	"errors"
	"fmt"
)

var (
	ErrorURLNotFound = errors.New("url not found")
)

type Repository map[string]string

func NewRepository() Repository {
	return make(map[string]string)
}

func (r Repository) Save(url *URL) {
	r[url.Short] = url.Full
}

func (r Repository) FindByShort(short string) (url *URL, err error) {
	full, exists := r[short]
	if !exists {
		return nil, fmt.Errorf("%w for %q", ErrorURLNotFound, short)
	}

	url = &URL{
		Full:  full,
		Short: short,
	}

	return url, nil
}
