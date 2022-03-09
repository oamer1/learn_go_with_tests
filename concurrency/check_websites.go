package concurrency

// WebsiteChecker checks a url, returning a bool.
type WebsiteChecker func(string) bool

// As we don't need either value to be named,
type result struct {
	string
	bool
}

// CheckWebsites takes a WebsiteChecker and a slice of urls and returns  a map.
// of urls to the result of checking each url with the WebsiteChecker function
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// Parralized part
	// Anonymous functions have
	// 1-executed at the same time that they're declared
	// 2- lexical scope they are defined
	for _, url := range urls {
		// Use url parameter in anyon functions or re-init
		// Generate WARNING: DATA RACE if run "go test -race"
		url := url
		go func() {
			resultChannel <- result{url, wc(url)}
		}()

	}

	// Linear part to avoid race condition
	// ensuring that it happens one at a time
	// go test -race now passes
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
