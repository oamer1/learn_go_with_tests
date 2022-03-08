package concurrency

// WebsiteChecker checks a url, returning a bool.
type WebsiteChecker func(string) bool

// CheckWebsites takes a WebsiteChecker and a slice of urls and returns  a map.
// of urls to the result of checking each url with the WebsiteChecker function
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// Anonymous functions have
	// 1-executed at the same time that they're declared
	// 2- lexical scope they are defined
	for _, url := range urls {

		go func() {
			results[url] = wc(url)
		}()

	}

	return results
}
