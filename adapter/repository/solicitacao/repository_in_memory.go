package solicitacao

type InMemorySolicitacaoRepository struct {
	Solicitacao []any
}

func (r *InMemorySolicitacaoRepository) Salvar(
	servicoId uint,
	statusId uint,
	concluida bool,
	solicitanteId uint,
) (uint, error) {
	itemSalvar := map[uint][]any{
		1: {servicoId, statusId, concluida, solicitanteId},
	}
	r.Solicitacao = append(r.Solicitacao, itemSalvar)
	return 1, nil
}
