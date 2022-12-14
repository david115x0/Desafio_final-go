package dto

// DentistaUpdateDTO is used by client when PUT update profile
type DentistaUpdateDTO struct {
	Id        uint64 `json:"id" form:"id"`
	Nome      string `json:"nome" form:"nome" binding:"required"`
	Sobrenome string `json:"sobrenome" form:"sobrenome" binding:"required"`
	Matricula string `json:"matricula" form:"matricula" binding:"required"`
}
