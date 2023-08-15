package solicitacao

type InterfaceSolicitacaoRepository interface {
	Salvar(servicoId uint, statusId uint) error
}
