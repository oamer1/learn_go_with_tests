package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// Always make channels For channels the zero value is nil and if you try and send to it with <- it will block forever because you cannot send to nil channels
// Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory
func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
