package campo

type Campo struct {
	Id            uint
	Valor         string
	SolicitacaoId uint
}

func NovoCampo(id uint, valor string, solicitacaoId uint) *Campo {
	return &Campo{
		Id:            id,
		Valor:         valor,
		SolicitacaoId: solicitacaoId,
	}
}
