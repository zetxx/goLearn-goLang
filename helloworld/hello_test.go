package main
import "testing"
func assertHello (t testing.TB, got, wanted string, ) {
	t.Helper()
	if got != wanted {
		t.Errorf("got %q want %q", got, wanted)
	}
}
func TestHello(t *testing.T) {
	t.Run("Pass string for name and nothing for lang", func(t *testing.T){
		assertHello(t, hello("elin",""), "hello, elin")
	})
	t.Run("Pass string for name on spanish", func(t *testing.T){
		assertHello(t, hello("elin", "Spanish"), "hola, elin")
	})
	t.Run("Pass string for name on French", func(t *testing.T){
		assertHello(t, hello("elin", "French"), "ahoi, elin")
	})
	t.Run("Pass empty val for name", func(t *testing.T){
		assertHello(t, hello("", ""), "hello, world")
	})
	t.Run("Pass empty val for name and Spanish for language", func(t *testing.T){
		assertHello(t, hello("", "Spanish"), "hola, world")
	})
	t.Run("Pass empty val for name and French for language", func(t *testing.T){
		assertHello(t, hello("", "French"), "ahoi, world")
	})
	
}

