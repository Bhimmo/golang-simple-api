package salvar_solicitacao

import "github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"

type SalvarSolicitacaoInput struct {
	ServicoId     uint          `json:"servico_id"`
	Campos        []campo.Campo `json:"campos"`
	SolicitanteId uint          `json:"solicitante_id"`
}
