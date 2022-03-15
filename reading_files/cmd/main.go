package main

// As maintainers, we can be confident our tests are useful because they're from a consumer's point of view.
// We're not testing implementation details or other incidental details, so we can be reasonably confident
// that our tests will help us, rather than hinder us when refactoring.

import (
	blogposts "learn_go_with_tests/reading_files"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
