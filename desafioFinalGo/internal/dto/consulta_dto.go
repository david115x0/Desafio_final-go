package dto

// ConsultaUpadateDTO  is a model that client use when updating a consulta
type ConsultaDTO struct {
	Id        int    `json:"id" form:"id" binding:"required"`
	Data      string `mysql:"type:date" json:"data"`
	Hora      string `mysql:"type:time" json:"hora"`
	Descricao string `mysql:"type:varchar(255)" json:"descricao"`
}

// ConsultaCreateDTO is a model that clinet use when create a new book
//type ConsultaCreateDTO struct {
//	PacienteID uint64 `json:"paciente_id,omitempty"  form:"paciente_id,omitempty"`
//	DentistaID uint64 `json:"dentista_id,omitempty"  form:"dentista_id,omitempty"`
//	Data       string `mysql:"type:date" json:"data"`
//	Hora       string `mysql:"type:time" json:"hora"`
//	Descricao  string `mysql:"type:varchar(255)" json:"descricao"`
//}
