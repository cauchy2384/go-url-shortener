package internal

import "fmt"

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
		return nil, fmt.Errorf("url not found for %q", short)
	}

	url = &URL{
		Full:  full,
		Short: short,
	}
	return url, nil
}
