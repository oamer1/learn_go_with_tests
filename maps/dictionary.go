package main

import "errors"

// So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.
// Dictionary type which acts as a thin wrapper around map. With the custom type defined
type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	defintion, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defintion, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err

	}
	return nil
}
