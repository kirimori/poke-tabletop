package pokeapi

import (
	"errors"
)

// Link is a link for further information in the api
type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (l Link) Get(target interface{}) error {
	if len(l.URL) == 0 {
		return errors.New("URL was not initialized for retrieval")
	}
	return getFromURL(l.URL, target)
}
