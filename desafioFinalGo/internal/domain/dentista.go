package domain

type Dentista struct {
	Id        int    `json:"id,omitempty"`
	Nome      string `json:"nome,omitempty"`
	Sobrenome string `json:"sobrenome,omitempty"`
	Matricula int    `json:"matricula,omitempty"`
}
