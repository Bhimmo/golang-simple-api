package servico

type InterfaceServicoRepository interface {
	Inserir(nome string) error
	PegandoPeloId(id uint) (Servico, error)
}
