package consulta

import (
	"desafioFinal/internal/domain"
	"desafioFinal/pkg/store"
)

type Repository interface {
	GetByID(id int) (domain.Consulta, error)
	Create(c domain.Consulta) (domain.Consulta, error)
	Update(id int, c domain.Consulta) (domain.Consulta, error)
	Delete(id int) error
	GetByRg(rg int) (domain.ConsultaDTO, error)
}

type repository struct {
	store store.StoreInterfaceConsulta
}

func (r *repository) GetByRg(rg int) (domain.ConsultaDTO, error) {
	c, err := r.store.GetByRg(rg)
	if err != nil {
		return c, err
	}
	return c, nil
}

// NewRepository cria um novo repositório
func NewRepository(store store.StoreInterfaceConsulta) Repository {
	return &repository{store}
}

// GetByID busca uma consulta pelo seu ID
func (r *repository) GetByID(id int) (domain.Consulta, error) {
	pReturn, err := r.store.GetByID(id)
	if err != nil {
		return domain.Consulta{}, err
	}

	return pReturn, nil
}

// Create adiciona uma nova consulta
func (r *repository) Create(c domain.Consulta) (domain.Consulta, error) {
	//if !r.validateRG(c.RG) {
	//	return domain.Consulta{}, errors.New("RG already exists")
	//}
	return r.store.Create(c)
}

// Delete exclui uma consulta
func (r *repository) Delete(id int) error {
	err := r.store.Delete(id)
	if err == nil {
		return nil
	}

	return err
}

// Update atualiza uma consulta
func (r *repository) Update(pid int, c domain.Consulta) (domain.Consulta, error) {
	return r.store.Update(pid, c)
}
