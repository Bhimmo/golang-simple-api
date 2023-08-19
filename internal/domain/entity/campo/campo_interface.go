package campo

type InterfaceCampoRepository interface {
	Salvar(campo Campo) error
	BuscarCampoPeloSolicitanteId(solicitanteId uint) ([]Campo, error)
}
