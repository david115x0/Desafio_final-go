package consulta

import (
	"desafioFinal/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Consulta, error)
	Create(c domain.Consulta) (domain.Consulta, error)
	Update(id int, c domain.Consulta) (domain.Consulta, error)
	Delete(id int) error
	GetByRg(rg int) (domain.ConsultaDTO, error)
}

type service struct {
	r Repository
}

// NewService cria um novo serviço
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByRg(rg int) (domain.ConsultaDTO, error) {
	c, err := s.r.GetByRg(rg)
	if err != nil {
		return c, err
	}
	return c, nil
}

// GetAll retorna todos as consultas
//func (s *service) GetAll() ([]domain.Consulta, error) {
//	l := s.r.GetAll()
//	return l, nil
//}

// GetByID busca uma consulta pelo seu id
func (s *service) GetByID(id int) (domain.Consulta, error) {
	c, err := s.r.GetByID(id)
	if err != nil {
		return domain.Consulta{}, err
	}
	return c, nil
}

// Create adiciona uma nova consulta
func (s *service) Create(c domain.Consulta) (domain.Consulta, error) {
	consult, err := s.r.Create(c)
	if err != nil {
		return domain.Consulta{}, err
	}
	return consult, nil
}

// Delete apaga uma consulta
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update atualiza uma consulta
func (s *service) Update(id int, u domain.Consulta) (domain.Consulta, error) {
	c, err := s.r.GetByID(id)
	if err != nil {
		return domain.Consulta{}, err
	}
	if u.Data != "" {
		c.Data = u.Data
	}
	if u.Descricao != "" {
		c.Descricao = u.Descricao
	}
	c, err = s.r.Update(id, c)
	if err != nil {
		return domain.Consulta{}, err
	}
	return c, nil
}
