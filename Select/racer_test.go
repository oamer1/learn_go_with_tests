package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// In the mocking and dependency injection chapters, we covered how ideally we don't want to be relying on external services to test our code because they can be
// Slow
// Flaky
// Can't test edge cases

func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	// execute at end but keep near
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
