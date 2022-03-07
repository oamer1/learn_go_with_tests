package main

// Dictionary type which acts as a thin wrapper around map. With the custom type defined
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	return d[word], nil
}
