package status

const (
	SolicitacaoIniciada    uint = 1
	ESPERA_DA_APROVACAO         = 2
	SOLICITACAO_FINALIZADA      = 3
)

type Status struct {
	Id   uint
	Nome string
}

func NovoStatus() Status {
	return Status{}
}

func (s *Status) TendoStatusInicial() {
	s.Id = SolicitacaoIniciada
	s.Nome = "Solicitacao iniciada"
}
func (s *Status) TendoStatusDesejado(id uint) {
	switch id {
	case SolicitacaoIniciada:
		s.Id = SolicitacaoIniciada
		s.Nome = "Solicitacao iniciada"
	case ESPERA_DA_APROVACAO:
		s.Id = ESPERA_DA_APROVACAO
		s.Nome = "Espera da aprovacao do responsavel"
	case SOLICITACAO_FINALIZADA:
		s.Id = SOLICITACAO_FINALIZADA
		s.Nome = "Solicitacao finalizada"
	}
}
