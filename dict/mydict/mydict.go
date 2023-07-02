package mydict

import "errors"

var (
	errNotFound   = errors.New("not Found")
	errWordExists = errors.New("that word already exists")
	errCantUpdate = errors.New("can`t update non-existing word")
)

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

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
