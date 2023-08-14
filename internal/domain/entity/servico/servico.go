package servico

type Servico struct {
	Id   uint
	Nome string
}

func NovoServico() *Servico {
	return &Servico{}
}
