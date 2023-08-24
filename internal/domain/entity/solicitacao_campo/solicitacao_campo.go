package solicitacao_campo

type SolicitacaoCampo struct {
	Id            uint
	CampoId       uint
	SolicitacaoId uint
	Valor         string
}

func NewSolicitacaoCampo(
	campoId uint,
	solicitacaoId uint,
	valor string,
) *SolicitacaoCampo {
	return &SolicitacaoCampo{
		CampoId:       campoId,
		SolicitacaoId: solicitacaoId,
		Valor:         valor,
	}
}
