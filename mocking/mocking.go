package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}
type DefSleeper struct{}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (ds ConfigurableSleeper) Sleep() {
	ds.sleep(ds.duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
