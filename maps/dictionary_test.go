package main

import "testing"

// The key type is special. It can only be a comparable type because without the ability to
// tell if 2 keys are equal, we have no way to ensure that we are getting the correct valu
func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
