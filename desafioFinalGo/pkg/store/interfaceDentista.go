package store

import "desafioFinal/internal/domain"

type StoreInterfaceDentista interface {
	GetAll() []domain.Dentista
	GetByID(id int) (domain.Dentista, error)
	Create(d domain.Dentista) (domain.Dentista, error)
	Update(id int, d domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
}
