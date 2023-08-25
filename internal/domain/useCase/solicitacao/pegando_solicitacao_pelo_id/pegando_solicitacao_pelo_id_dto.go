package pegando_solicitacao_pelo_id

import "time"

type PegandoSolicitacaoPeloIdOutput struct {
	Id            uint                                  `json:"id"`
	Servico       PegandoSolicitacaoPeloIdServicoOutput `json:"servico"`
	Status        string                                `json:"status"`
	Concluida     bool                                  `json:"concluida"`
	SolicitanteId uint                                  `json:"solicitante_id"`
	CreatedAt     time.Time                             `json:"createdAt"`
	Campos        []PegandoSolicitacaoPeloIdCampoOutput `json:"campos"`
}

type PegandoSolicitacaoPeloIdServicoOutput struct {
	Id   uint   `json:"id"`
	Nome string `json:"nome"`
}
type PegandoSolicitacaoPeloIdCampoOutput struct {
	Id    uint   `json:"id"`
	Nome  string `json:"nome"`
	Valor string `json:"valor"`
}
