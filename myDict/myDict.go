package myDict

import "errors"

// alias
// type에 대한 alias를 만들 수 있다.

// type 이 method를 가질 수 있음

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

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

