package blogposts_test

// By appending _test to our intended package name, we only access exported members from our package - just like a real user of our package.

import (
	"errors"
	"io/fs"
	blogposts "learn_go_with_tests/reading_files"
	"testing"
	"testing/fstest"
)

// inject a failing fs.FS test-double to make fs.ReadDir return an error.
type StubFailingFS struct {
}

func (s StubFailingFS) open(name string) (fs.File, error) {
	return nil, errors.New("Always failing file")
}

func TestNewBlogPosts(t *testing.T) {
	// in-memory file system for testing
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

}
