package myDict

import "errors"

// alias
// type에 대한 alias를 만들 수 있다.

// type 이 method를 가질 수 있음

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound = errors.New("Not Found")
 errCantUpdate = errors.New("Cant update non-existing word")
 errWordExists = errors.New("That word already exist")

)

// 에러 처리를 먼저
// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	// value, ok(boolean)
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }

	switch err {
		case errNotFound:
			d[word] = def
		case nil:
			return errWordExists
	}

	return nil
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
		case nil:
			d[word] = definition
		case errNotFound:
			return errCantUpdate
	}
	return nil
}


// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
		case nil:
			delete(d, word)
		case errNotFound:
			return errNotFound
	}
	return nil
}