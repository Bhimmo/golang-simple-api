package servico_campo

type ServicoCampo struct {
	Id        uint
	ServicoId uint
	CampoId   uint
}

func NewServicoCampo() *ServicoCampo {
	return &ServicoCampo{}
}
