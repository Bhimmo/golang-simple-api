package atualizar_status_solicitacao

type AtualizarStatusSolicitacaoOutput struct {
	Id        uint                                   `json:"id"`
	Concluida bool                                   `json:"concluida"`
	Status    AtualizarStatusSolicitacaoStatusOutput `json:"status"`
}

type AtualizarStatusSolicitacaoStatusOutput struct {
	Id   uint   `json:"id"`
	Nome string `json:"nome"`
}
