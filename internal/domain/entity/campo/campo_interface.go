package campo

type InterfaceCampoRepository interface {
	Salvar(campo Campo) error
	BuscarCampoPeloSolicitacaoId(solicitacaoId uint) ([]Campo, error)
}
