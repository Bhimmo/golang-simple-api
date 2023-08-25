package pegando_campo_servico

type PegandoCampoServicoInput struct {
	ServicoId uint `json:"servico_id"`
}

type PegandoCampoServicoOutput struct {
	ServicoId uint                             `json:"servico_id"`
	Campos    []PegandoCampoServicoCampoOutput `json:"campos"`
}

type PegandoCampoServicoCampoOutput struct {
	Id   uint   `json:"id"`
	Nome string `json:"nome"`
}
