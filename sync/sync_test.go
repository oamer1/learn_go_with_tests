package sync

import (
	"sync"
	"testing"
)

// create a constructor which shows readers of your API that it would be better to not initialise the type yourself.
func NewCounter() *Counter {
	return &Counter{}
}
func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)

	})
	t.Run("Runs safely concurrently", func(t *testing.T) {

		wantedCount := 1000
		counter := NewCounter()

		// A WaitGroup waits for a collection of goroutines to finish.
		// The main goroutine calls Add to set the number of goroutines to wait for.
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()

			}()
		}

		// block until all routines finish
		wg.Wait()
		assertCounter(t, counter, wantedCount)

	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

// When we pass our Counter (by value) to assertCounter it will try and create a copy of the mutex.
