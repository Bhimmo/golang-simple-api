package cadastrar_campo

type CadastrarCampoInput struct {
	Nome string
}

type CadastrarCampoOutput struct {
	Id   uint   `json:"id"`
	Nome string `json:"nome"`
}
