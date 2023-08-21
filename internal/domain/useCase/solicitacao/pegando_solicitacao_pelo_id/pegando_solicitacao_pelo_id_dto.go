package pegando_solicitacao_pelo_id

type PegandoSolicitacaoPeloIdOutput struct {
	Id            uint                                  `json:"id"`
	Servico       PegandoSolicitacaoPeloIdServicoOutput `json:"servico"`
	Status        PegandoSolicitacaoPeloIdStatusOutput  `json:"status"`
	Concluida     bool                                  `json:"concluida"`
	SolicitanteId uint                                  `json:"solicitante_id"`
	Campos        []PegandoSolicitacaoPeloIdCampoOutput `json:"campos"`
}

type PegandoSolicitacaoPeloIdServicoOutput struct {
	Id   uint   `json:"id"`
	Nome string `json:"nome"`
}
type PegandoSolicitacaoPeloIdStatusOutput struct {
	Id   uint   `json:"id"`
	Nome string `json:"nome"`
}
type PegandoSolicitacaoPeloIdCampoOutput struct {
	Id    uint   `json:"id"`
	Valor string `json:"valor"`
}
