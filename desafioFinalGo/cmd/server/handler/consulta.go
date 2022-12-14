package handler

import (
	"desafioFinal/internal/consulta"
	"desafioFinal/internal/domain"
	"desafioFinal/pkg/web"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type consultaHandler struct {
	s consulta.Service
}

// NewProductHandler crea un nuevo controller de consultas
func NewConsultHandler(s consulta.Service) *consultaHandler {
	return &consultaHandler{
		s: s,
	}
}

// validateExpiration valida que la fecha de expiracion sea valida
func validateData(exp string) (bool, error) {
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
func countIds(ids []int) (result map[int]int) {
	result = make(map[int]int)
	for _, id := range ids {
		if v, found := result[id]; found {
			result[id] = v + 1
		} else {
			result[id] = 1
		}
	}
	return
}

// GetByRg busca dentistas por rg
// @Summary Get a consult by rg
// @Tags CRUD CONSULTAS
// @Description Get a consult by rg
// @Accept json
// @Produce json
// @Param rg path int true "Consulta RG"
// @Success 200 {object} domain.ConsultaDTO
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /consultas/consulta_paciente/{rg} [get]
func (h *consultaHandler) GetByRg() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("rg")
		dados, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		consult, err := h.s.GetByRg(dados)
		if err != nil {
			web.Failure(c, 404, errors.New("consult not found"))
			return
		}
		web.Success(c, 200, consult)
	}
}

// GetByID busca consulta por id
// @Summary Get a consult by id
// @Tags CRUD CONSULTAS
// @Description Get a consult by id
// @Accept json
// @Produce json
// @Param id path int true "Consulta id"
// @Success 200 {object} domain.Consulta
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /consultas/{id} [get]
func (h *consultaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		consult, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("consult not found"))
			return
		}
		web.Success(c, 200, consult)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptysConsult(c *domain.Consulta) (bool, error) {
	switch {
	case c.Data == "" || c.Hora == "" || c.Descricao == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post cria uma consulta
// @Summary Post a consult
// @Tags CRUD CONSULTAS
// @Description Post a consult
// @Accept json
// @Produce json
// @Param json body domain.Consulta true "Consulta"
// @Success 200 {object} domain.Consulta
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /consultas [post]
func (h *consultaHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var consult domain.Consulta
		err := c.ShouldBindJSON(&consult)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysConsult(&consult)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateData(consult.Data)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(consult)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete deleta uma consulta
// @Summary Delete a consult by ID
// @Tags CRUD CONSULTAS
// @Description Delete a consult by ID
// @Accept json
// @Produce json
// @Param id path int true "Consulta ID"
// @Success 200 {object} domain.Consulta
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /consultas/{id} [delete]
func (h *consultaHandler) Delete() gin.HandlerFunc {
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

// Put actualiza uma consulta
// @Summary Put a consult
// @Tags CRUD CONSULTAS
// @Description Put a consult
// @Accept json
// @Produce json
// @Param id path int true "Consulta ID"
// @Param json body domain.Consulta true "Consulta"
// @Success 200 {object} domain.Consulta
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /consultas/{id} [put]
func (h *consultaHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("consult not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var consult domain.Consulta
		err = c.ShouldBindJSON(&consult)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysConsult(&consult)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateData(consult.Data)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, consult)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch actualiza uma consulta ou algum de seus campos
// @Summary Patch a consult
// @Tags CRUD CONSULTAS
// @Description Patch a consult
// @Accept json
// @Produce json
// @Param id path int true "Consulta ID"
// @Param json body domain.Consulta true "Consulta"
// @Success 200 {object} domain.Consulta
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /consultas/{id} [patch]
func (h *consultaHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Data       string `json:"data,omitempty"`
		Hora       string `json:"hora,omitempty"`
		Descricao  string `json:"descricao,omitempty"`
		PacienteID int    `json:"paciente_id, omitempty"`
		DentistaID int    `json:"dentista_id, omitempty"`
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
			web.Failure(c, 404, errors.New("consult not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Consulta{
			Data:      r.Data,
			Hora:      r.Hora,
			Descricao: r.Descricao,
		}
		if update.Data != "" {
			valid, err := validateDataCadastro(update.Data)
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
