package store

import "desafioFinal/internal/domain"

type StoreInterfacePaciente interface {
	GetAll() []domain.Paciente
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}
