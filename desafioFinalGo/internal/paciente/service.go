package paciente

import "desafioFinal/internal/domain"

type Service interface {
	GetAll() ([]domain.Paciente, error)
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

// NewService cria um novo serviço
func NewService(r Repository) Service {
	return &service{r}
}

// GetAll retorna todos os pacientes
func (s *service) GetAll() ([]domain.Paciente, error) {
	l := s.r.GetAll()
	return l, nil
}

// GetByID busca um paciente pelo seu id
func (s *service) GetByID(id int) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

// Create adiciona um novo paciente
func (s *service) Create(p domain.Paciente) (domain.Paciente, error) {
	pacient, err := s.r.Create(p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return pacient, nil
}

// Delete apaga um paciente
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update atualiza um paciente
func (s *service) Update(id int, u domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	if u.Nome != "" {
		p.Nome = u.Nome
	}
	if u.Sobrenome != "" {
		p.Sobrenome = u.Sobrenome
	}
	if u.RG != 0 {
		p.RG = u.RG
	}
	if u.DataCadastro != "" {
		p.DataCadastro = u.DataCadastro
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}
