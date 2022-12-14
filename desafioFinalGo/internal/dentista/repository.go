package dentista

import (
	"desafioFinal/internal/domain"
	"desafioFinal/pkg/store"
	"errors"
)

type Repository interface {
	GetAll() []domain.Dentista
	GetByID(id int) (domain.Dentista, error)
	Create(d domain.Dentista) (domain.Dentista, error)
	Update(id int, d domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
}

type repository struct {
	store store.StoreInterfaceDentista
}

// NewRepository cria um novo repositório
func NewRepository(store store.StoreInterfaceDentista) Repository {
	return &repository{store}
}

// GetAll devolve todos os dentistas
func (r *repository) GetAll() []domain.Dentista {
	return r.store.GetAll()
}

// GetByID busca um dentistas pelo seu ID
func (r *repository) GetByID(id int) (domain.Dentista, error) {
	pReturn, err := r.store.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}

	return pReturn, nil
}

// Create adiciona um novo dentistas
func (r *repository) Create(d domain.Dentista) (domain.Dentista, error) {
	if !r.validateCodeValue(d.Matricula) {
		return domain.Dentista{}, errors.New("matricula already exists")
	}
	return r.store.Create(d)
}

// validateCodeValue valida que a matricula não existe na lista de dentistas
func (r *repository) validateCodeValue(matricula int) bool {
	for _, dentist := range r.store.GetAll() {
		if dentist.Matricula == matricula {
			return false
		}
	}
	return true
}

// Delete exclui um dentista
func (r *repository) Delete(id int) error {
	err := r.store.Delete(id)
	if err == nil {
		return nil
	}

	return err
}

// Update atualiza um dentista
func (r *repository) Update(id int, d domain.Dentista) (domain.Dentista, error) {
	for _, dentist := range r.store.GetAll() {
		if dentist.Id == id {
			if !r.validateCodeValue(d.Matricula) && dentist.Matricula != d.Matricula {
				return domain.Dentista{}, errors.New("matricula already exists")
			}
			return r.store.Update(id, d)
		}
	}
	return domain.Dentista{}, errors.New("dentist not found")
}
