package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleep struct {
	Calls int
}

func (s *SpySleep) Sleep() {
	s.Calls++
}

type SpySleepOrder struct {
	Calls []string
}

func (s *SpySleepOrder) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpySleepOrder) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("count", func(t *testing.T) {
		buf := &bytes.Buffer{}
		sl := &SpySleep{}
		Countdown(buf, sl)

		got := buf.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %s != want %s", got, want)
		}
		if sl.Calls != 3 {
			t.Errorf("Sleep repeat not correct, got %d want %d", sl.Calls, 3)
		}
	})

	t.Run("order", func(t *testing.T) {
		sl := &SpySleepOrder{}
		Countdown(sl, sl)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}
		if !reflect.DeepEqual(want, sl.Calls) {
			t.Errorf("got %v != want %v", sl.Calls, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
