package cadastrando_campo_servico

type CadastrandoCampoServicoInput struct {
	ServicoId uint   `json:"servico_id"`
	Campos    []uint `json:"campos"`
}
