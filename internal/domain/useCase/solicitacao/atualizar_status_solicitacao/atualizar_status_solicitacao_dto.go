package atualizar_status_solicitacao

type AtualizarStatusSolicitacaoOutput struct {
	Id        uint   `json:"id"`
	Concluida bool   `json:"concluida"`
	Status    string `json:"status"`
}
