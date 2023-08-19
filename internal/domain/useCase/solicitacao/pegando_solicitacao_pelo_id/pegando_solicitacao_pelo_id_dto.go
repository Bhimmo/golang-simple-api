package pegando_solicitacao_pelo_id

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
)

type PegandoSolicitacaoPeloIdOutput struct {
	Solicitacao solicitacao.Solicitacao
	Campos      []campo.Campo
}
