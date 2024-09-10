package httptest

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func spawnServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRace(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		slowServer := spawnServer(2 * time.Second)
		fastServer := spawnServer(0 * time.Second)
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL
		want := fastUrl

		got, _ := Race(slowUrl, fastUrl, 5*time.Second)

		if got != want {
			t.Errorf("got %s != want %s", got, want)
		}
	})
	t.Run("2", func(t *testing.T) {
		slowServer := spawnServer(4 * time.Second)
		fastServer := spawnServer(0 * time.Second)
		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Race(slowServer.URL, fastServer.URL, 3*time.Second)
		if err != nil {
			t.Errorf("want error, didnt get error")
		}
	})
}
