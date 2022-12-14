package domain

type Consulta struct {
	Id         int      `json:"id,omitempty"`
	PacienteID int      `json:"paciente_id,omitempty"`
	DentistaID int      `json:"dentista_id,omitempty"`
	Data       string   `json:"data,omitempty"`
	Hora       string   `json:"hora,omitempty"`
	Descricao  string   `json:"descricao,omitempty"`
	Paciente   Paciente `json:"paciente,omitempty"`
	Dentista   Dentista `json:"dentista,omitempty"`
}

type ConsultaDTO struct {
	Consulta
	Paciente Paciente `json:"paciente,omitempty"`
	Dentista Dentista `json:"dentista,omitempty"`
}
type ConsultaDTORgMatricula struct {
	Consulta
	PacienteID int    `json:"paciente_id,omitempty"`
	DentistaID int    `json:"dentista_id,omitempty"`
	Data       string `json:"data,omitempty"`
	Hora       string `json:"hora,omitempty"`
	Descricao  string `json:"descricao,omitempty"`
}
