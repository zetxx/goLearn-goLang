package main

import (
	"fmt"
)
const (
	french = "French"
	spanish = "Spanish"

	prefEn = "hello, "
	prefEs = "hola, "
	prefFr = "ahoi, "
)

func main() {
	fmt.Println(hello("elin", "Spanish"))
}
func greetingsPrefix(lang string) (prefix string) {
	prefix = prefEn
	if lang == spanish {
		prefix = prefEs
	} else if lang == french {
		prefix = prefFr
	}
	return
}
func hello(name, lang string) string {
	if name == "" {
		name = "world"
	}
	
	return greetingsPrefix(lang) + name
}

