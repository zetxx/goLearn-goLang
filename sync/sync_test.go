package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 tames and live it at 3", func(t *testing.T) {
		counter := Counter{}

		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got: %d != want %d", counter.Value(), 3)
		}
	})

	t.Run("runs safe in concurent env", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < 1000; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		if counter.Value() != 1000 {
			t.Errorf("got %d != want %d", counter.Value(), wantedCount)
		}
	})
}
