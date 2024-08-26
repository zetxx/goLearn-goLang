package test

type Ai interface {
	Aia() int
}

type As struct {
	asa int
}

func (a As) Aia() int {
	return 1
}

func main() {
	_a := As{asa: 33,}
	_a.Aia()
}
