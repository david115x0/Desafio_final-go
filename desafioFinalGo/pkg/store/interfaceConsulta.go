package store

import "desafioFinal/internal/domain"

type StoreInterfaceConsulta interface {
	GetByID(id int) (domain.Consulta, error)
	Create(c domain.Consulta) (domain.Consulta, error)
	Update(pid int, c domain.Consulta) (domain.Consulta, error)
	Delete(id int) error
	GetByRg(rg int) (domain.ConsultaDTO, error)
}
