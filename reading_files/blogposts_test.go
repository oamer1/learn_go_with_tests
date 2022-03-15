package blogposts_test

// By appending _test to our intended package name, we only access exported members from our package - just like a real user of our package.

import (
	blogposts "learn_go_with_tests/reading_files"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	// in-memory file system for testing
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
