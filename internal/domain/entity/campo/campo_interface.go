package campo

type InterfaceCampoRepository interface {
	Salvar(campo Campo) error
}
