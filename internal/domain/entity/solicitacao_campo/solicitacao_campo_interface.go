package solicitacao_campo

type SolicitacaoCampoInterface interface {
	SalvarCamposDaSolicitacao(campoId uint, solicitacaoId uint, valor string) error
	BuscarCamposPelaSolicitacao(solicitacaoId uint) ([]SolicitacaoCampo, error)
}
