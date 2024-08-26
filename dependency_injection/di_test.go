package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "you")
	got := buffer.String()
	want := "Haloz, you"
	fmt.Println(got)
	if got != want {
		t.Errorf("Got %s != Want %s", got, want)
	}
}
