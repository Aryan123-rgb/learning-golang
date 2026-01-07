package websiteracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("compares the speed of servers, returning the url of the fastest one", func(t *testing.T) {
		serverA := createDelayedServer(20 * time.Millisecond)
		serverB := createDelayedServer(0)

		defer serverA.Close()
		defer serverB.Close()

		want := serverB.URL
		got, err := Racer(serverA.URL, serverB.URL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})

	t.Run("returns an error if the server does not respond within 10 second", func(t *testing.T) {
		serverA := createDelayedServer(12 * time.Second)
		serverB := createDelayedServer(11 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		urlA := serverA.URL
		urlB := serverB.URL

		_, err := Racer(urlA, urlB)

		if err == nil {
			t.Errorf("timed out waiting for %s and %s", urlA, urlB)
		}
	})
}
