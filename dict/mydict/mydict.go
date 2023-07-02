package mydict

import "errors"

var errNotFound = errors.New("not Found")

// Dictionary type
type Dictionary map[string]string

// Search for a world
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
