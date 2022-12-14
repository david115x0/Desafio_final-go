package paciente

import (
	"desafioFinal/internal/domain"
	"desafioFinal/pkg/store"
	"errors"
)

type Repository interface {
	GetAll() []domain.Paciente
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

type repository struct {
	store store.StoreInterfacePaciente
}

// NewRepository cria um novo repositório
func NewRepository(store store.StoreInterfacePaciente) Repository {
	return &repository{store}
}

// GetAll devolve todos os pacientes
func (r *repository) GetAll() []domain.Paciente {
	return r.store.GetAll()
}

// GetByID busca um paciente pelo seu ID
func (r *repository) GetByID(id int) (domain.Paciente, error) {
	pReturn, err := r.store.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}

	return pReturn, nil
}

// Create adiciona um novo paciente
func (r *repository) Create(c domain.Paciente) (domain.Paciente, error) {
	if !r.validateRG(c.RG) {
		return domain.Paciente{}, errors.New("RG already exists")
	}
	return r.store.Create(c)
}

// validateRG valida que o RG não existe na lista de pacientes
func (r *repository) validateRG(rg int) bool {
	for _, registro := range r.store.GetAll() {
		if registro.RG == rg {
			return false
		}
	}
	return true
}

// Delete exclui um paciente
func (r *repository) Delete(id int) error {
	err := r.store.Delete(id)
	if err == nil {
		return nil
	}

	return err
}

// Update atualiza um paciente
func (r *repository) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	for _, paciente := range r.store.GetAll() {
		if paciente.Id == id {
			return r.store.Update(id, p)
		}
	}
	return domain.Paciente{}, errors.New("consulta not found")
}
