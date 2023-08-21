package solicitacao

type InterfaceSolicitacaoRepository interface {
	Salvar(
		servicoId uint,
		statusId uint,
		concluida bool,
		solicitanteId uint,
	) (uint, error)

	BuscarPeloId(id uint) (Solicitacao, error)
	AtualizarSolicitacao(solicitacao Solicitacao) error
}
