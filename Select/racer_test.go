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

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
