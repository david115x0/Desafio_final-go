package dentista

import (
	"desafioFinal/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Dentista, error)
	GetByID(id int) (domain.Dentista, error)
	Create(d domain.Dentista) (domain.Dentista, error)
	Update(id int, d domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

// NewService cria um novo serviço
func NewService(r Repository) Service {
	return &service{r}
}

// GetAll retorna todos os produtos
func (s *service) GetAll() ([]domain.Dentista, error) {
	l := s.r.GetAll()
	return l, nil
}

// GetByID busca um produto pelo seu id
func (s *service) GetByID(id int) (domain.Dentista, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	return d, nil
}

// Create adiciona um novo produto
func (s *service) Create(d domain.Dentista) (domain.Dentista, error) {
	dentist, err := s.r.Create(d)
	if err != nil {
		return domain.Dentista{}, err
	}
	return dentist, nil
}

// Delete apaga um produto
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update atualiza um produto
func (s *service) Update(id int, u domain.Dentista) (domain.Dentista, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	if u.Nome != "" {
		d.Nome = u.Nome
	}
	if u.Sobrenome != "" {
		d.Sobrenome = u.Sobrenome
	}
	if u.Matricula != 0 {
		d.Matricula = u.Matricula
	}
	d, err = s.r.Update(id, d)
	if err != nil {
		return domain.Dentista{}, err
	}
	return d, nil
}
