package shortener

import (
	"fmt"
)

// Repository defines in-memory storage for urls.
type repository map[string]string

// NewRepository returns new repository instance.
func newRepository() repository {
	return make(map[string]string)
}

// Save URL to storage.
func (r repository) Save(record *url) {
	r[record.Short] = record.Long
}

// FindByShort returns URL that corresponds to given short key.
// Returns ErrorURLNotFound if nothing exists.
func (r repository) findByShort(short string) (record *url, err error) {
	full, exists := r[short]
	if !exists {
		return nil, fmt.Errorf("%w for %q", ErrorURLNotFound, short)
	}

	record = &url{
		Long:  full,
		Short: short,
	}

	return record, nil
}
