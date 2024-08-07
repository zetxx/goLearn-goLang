package iteration

func Repeat(s string, times int) (ss string) {
	for i := 0; i < times; i++ {
		ss = ss + s
	}
	return
}
