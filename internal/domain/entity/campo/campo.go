package campo

type Campo struct {
	Id   uint
	Nome string
}

func NovoCampo() *Campo {
	return &Campo{}
}
