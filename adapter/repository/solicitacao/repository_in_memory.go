package solicitacao

type InMemorySolicitacaoRepository struct {
	Solicitacao [][]uint
}

func (r *InMemorySolicitacaoRepository) Salvar(servicoId uint, statusId uint) error {
	maps := []uint{servicoId, statusId}
	_ = append(r.Solicitacao, maps)

	return nil
}
