package dto

// PacienteUpdateDTO is used by client when PUT update profile
type PacienteUpdateDTO struct {
	Id           uint64 `json:"id" form:"id"`
	Nome         string `json:"nome" form:"nome" binding:"required"`
	Sobrenome    string `json:"sobrenome" form:"sobrenome" binding:"required"`
	RG           string `json:"rg" form:"rg" binding:"required"`
	DataCadastro string `json:"data_cadastro" form:"data_cadastro" binding:"required"`
}
