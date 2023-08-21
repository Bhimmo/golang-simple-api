package salvar_solicitacao

type SalvarSolicitacaoInput struct {
	ServicoId     uint                           `json:"servico_id"`
	SolicitanteId uint                           `json:"solicitante_id"`
	Campos        []SalvarSolicitacaoCampoOutput `json:"campos"`
}

type SalvarSolicitacaoOutput struct {
	Id            uint                           `json:"id"`
	Concluida     bool                           `json:"concluida"`
	SolicitanteId uint                           `json:"solicitante_id"`
	ServicoId     uint                           `json:"servico_id"`
	Campos        []SalvarSolicitacaoCampoOutput `json:"campos"`
}

type SalvarSolicitacaoCampoOutput struct {
	Id    uint   `json:"id"`
	Valor string `json:"valor"`
}
