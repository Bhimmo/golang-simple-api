package campo

type InterfaceCampoRepository interface {
	Salvar(campo Campo) (uint, error)

	BuscarPeloId(id uint) (Campo, error)
}
