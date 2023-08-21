package status

const (
	SolicitacaoIniciada   uint = 1
	EsperaDaAprovacao          = 2
	SolicitacaoFinalizada      = 3
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
	case EsperaDaAprovacao:
		s.Id = EsperaDaAprovacao
		s.Nome = "Espera da aprovacao do responsavel"
	case SolicitacaoFinalizada:
		s.Id = SolicitacaoFinalizada
		s.Nome = "Solicitacao finalizada"
	}
}

func (s *Status) ProximoStatus() {
	switch s.Id {
	case SolicitacaoIniciada:
		s.Id = EsperaDaAprovacao
		s.Nome = "Espera da aprovacao do responsavel"
	case EsperaDaAprovacao:
		s.Id = SolicitacaoFinalizada
		s.Nome = "Solicitacao finalizada"
	}
}

func (s *Status) VerificaUltimoStatus() bool {
	if s.Id == SolicitacaoFinalizada {
		return true
	}
	return false
}
