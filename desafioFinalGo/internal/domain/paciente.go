package domain

type Paciente struct {
	Id           int    `json:"id,omitempty"`
	Nome         string `json:"nome,omitempty"`
	Sobrenome    string `json:"sobrenome,omitempty"`
	RG           int    `json:"rg,omitempty"`
	DataCadastro string `json:"data_cadastro,omitempty"`
}
