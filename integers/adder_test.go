package integers

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	got := Add(1, 2)
	want := 3

	if got != want {
		t.Errorf("wrong adition. want: %d got: %d", want, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 6)
	if sum != 7 {
		fmt.Println(sum)
	}
}
