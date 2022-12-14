package handler

import (
	"desafioFinal/internal/domain"
	"desafioFinal/internal/paciente"
	"desafioFinal/pkg/web"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type pacienteHandler struct {
	s paciente.Service
}

// NewProductHandler crea un nuevo controller de pacientes
func NewProductHandlerPaciente(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

// validateExpiration valida que la fecha de expiracion sea valida
func validateDataCadastro(exp string) (bool, error) {
	dates := strings.Split(exp, "-")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid date, must be in format: yyyy-mm-dd")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}

// GetAll busca todos os pacientes
// @Summary List pacientes
// @Tags CRUD PACIENTES
// @Description get pacientes
// @Accept json
// @Produce json
// @Param SECRET_TOKEN header string true "SECRET_TOKEN"
// @Success 200 {object} domain.Paciente
// @Router /pacientes [get]
func (h *pacienteHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, _ := h.s.GetAll()
		web.Success(c, http.StatusOK, p)
	}
}

// GetByID gosdoc
// @Summary Get a paciente by ID
// @Tags CRUD PACIENTES
// @Description Get a paciente by ID
// @Accept json
// @Produce json
// @Param SECRET_TOKEN header string true "SECRET_TOKEN"
// @Param id path int true "Paciente ID"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes/{id} [get]
func (h *pacienteHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		paciente, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("pacient not found"))
			return
		}
		web.Success(c, 200, paciente)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptysPacient(p *domain.Paciente) (bool, error) {
	switch {
	case p.Nome == "" || p.Sobrenome == "" || p.RG == 0 || p.DataCadastro == "":
		return false, errors.New("fields can't be empty, our zero")
	}
	return true, nil
}

// Post cria um paciente
// @Summary Post a paciente
// @Tags CRUD PACIENTES
// @Description Post a paciente
// @Accept json
// @Produce json
// @Param SECRET_TOKEN header string true "SECRET_TOKEN"
// @Param json body domain.Paciente true "Paciente"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes [post]
func (h *pacienteHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pacient domain.Paciente
		err := c.ShouldBindJSON(&pacient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysPacient(&pacient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateDataCadastro(pacient.DataCadastro)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(pacient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete deleta um paciente
// @Summary Delete a paciente by ID
// @Tags CRUD PACIENTES
// @Description Delete a paciente by ID
// @Accept json
// @Produce json
// @Param SECRET_TOKEN header string true "SECRET_TOKEN"
// @Param id path int true "Paciente ID"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes/{id} [delete]
func (h *pacienteHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put actualiza un paciente
// @Summary Put a paciente
// @Tags CRUD PACIENTES
// @Description Put a paciente
// @Accept json
// @Produce json
// @Param SECRET_TOKEN header string true "SECRET_TOKEN"
// @Param id path int true "Paciente ID"
// @Param json body domain.Paciente true "Paciente"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes/{id} [put]
func (h *pacienteHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("pacient not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var pacient domain.Paciente
		err = c.ShouldBindJSON(&pacient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysPacient(&pacient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateDataCadastro(pacient.DataCadastro)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, pacient)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch actualiza un paciente ou algum de seus campos
// @Summary Patch a paciente
// @Tags CRUD PACIENTES
// @Description Patch a paciente
// @Accept json
// @Produce json
// @Param SECRET_TOKEN header string true "SECRET_TOKEN"
// @Param id path int true "Paciente ID"
// @Param json body domain.Paciente true "Paciente"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes/{id} [patch]
func (h *pacienteHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Nome         string `json:"nome,omitempty"`
		Sobrenome    string `json:"sobrenome,omitempty"`
		RG           int    `json:"rg,omitempty"`
		DataCadastro string `json:"data_cadastro,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("pacient not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Paciente{
			Nome:         r.Nome,
			Sobrenome:    r.Sobrenome,
			RG:           r.RG,
			DataCadastro: r.DataCadastro,
		}
		if update.DataCadastro != "" {
			valid, err := validateDataCadastro(update.DataCadastro)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
