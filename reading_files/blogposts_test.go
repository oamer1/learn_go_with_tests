package blogposts_test

// By appending _test to our intended package name, we only access exported members from our package - just like a real user of our package.

import (
	"errors"
	"io/fs"
	blogposts "learn_go_with_tests/reading_files"
	"reflect"
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
	const (
		firstBody = `Title: Post 1
Description: Description 1`
		secondBody = `Title: Post 2
Description: Description 2`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}
	posts, _ := blogposts.NewPostsFromFS(fs)

	// rest of test code cut for brevity
	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
	})

}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
