package mydict

import "errors"

var errNotFound = errors.New("not Found")
var errWordExists = errors.New("that word already exists")

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

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	//if err == errNotFound {
	//	d[word] = def
	//} else if err == nil {
	//	return errWordExists
	//}
	//return nil

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
