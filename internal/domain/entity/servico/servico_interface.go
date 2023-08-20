package servico

type InterfaceServicoRepository interface {
	Inserir(nome string) (uint, error)
	PegandoPeloId(id uint) (Servico, error)
}
