package errortypes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// DumbGetter will get the string body of url if it gets a 200
func DumbGetter(url string) (string, error) {
	res, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("problem fetching from %s, %v", url, err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("did not get 200 from %s, got %d", url, res.StatusCode)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body) // ignoring err for brevity

	return string(body), nil
}
func TestDumbGetter(t *testing.T) {
	t.Run("when you don't get a 200 you get a status error", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusTeapot)
		}))
		defer svr.Close()

		_, err := DumbGetter(svr.URL)

		if err == nil {
			t.Fatal("expected an error")
		}

		want := fmt.Sprintf("did not get 200 from %s, got %d", svr.URL, http.StatusTeapot)
		got := err.Error()

		if got != want {
			t.Errorf(`got "%v", want "%v"`, got, want)
		}
	})
}
