package iteration

import "testing"

func IterationTest(t *testing.T) {
	got := Repeat("a", 3)
	want := "aaa"
	if got != want {
		t.Errorf("Repetition fail. got: %s want: %s", got, want)
	}
}
