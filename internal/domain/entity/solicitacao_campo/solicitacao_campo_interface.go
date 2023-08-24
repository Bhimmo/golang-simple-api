package solicitacao_campo

type SolicitacaoCampoInterface interface {
	BuscarCamposPelaSolicitacao(solicitacaoId uint) ([]SolicitacaoCampo, error)
}
