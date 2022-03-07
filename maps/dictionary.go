package main

import "errors"

// Dictionary type which acts as a thin wrapper around map. With the custom type defined
type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	defintion, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defintion, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
