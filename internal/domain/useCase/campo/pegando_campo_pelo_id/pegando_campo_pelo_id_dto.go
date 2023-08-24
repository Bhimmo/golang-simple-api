package pegando_campo_pelo_id

type PegandoCampoPeloIdInput struct {
	Id uint `json:"id"`
}

type PegandoCampoPeloIdOutput struct {
	Id   uint   `json:"id"`
	Nome string `json:"nome"`
}
