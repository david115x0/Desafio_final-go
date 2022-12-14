package handler

import (
	"desafioFinal/internal/dentista"
	"desafioFinal/internal/domain"
	"desafioFinal/pkg/web"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type dentistaHandler struct {
	s dentista.Service
}

// NewProductHandler crea un nuevo controller de dentista
func NewProductHandler(s dentista.Service) *dentistaHandler {
	return &dentistaHandler{
		s: s,
	}
}

// GetAll busca todos os dentistas
// @Summary List dentistas
// @Tags CRUD DENTISTAS
// @Description getall dentistas
// @Accept json
// @Produce json
// @Success 200 {object} domain.Dentista
// @Router /dentistas [get]
func (h *dentistaHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		dentists, _ := h.s.GetAll()
		web.Success(c, http.StatusOK, dentists)
	}
}

// GetByID busca dentistas por id
// @Summary Get a dentista by ID
// @Tags CRUD DENTISTAS
// @Description Get a dentist by ID
// @Accept json
// @Produce json
// @Param id path int true "Dentista ID"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/{id} [get]
func (h *dentistaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		product, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		web.Success(c, 200, product)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(d *domain.Dentista) (bool, error) {
	switch {
	case d.Nome == "" || d.Sobrenome == "" || d.Matricula == 0:
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post cria um dentista
// @Summary Post a dentist
// @Tags CRUD DENTISTAS
// @Description Post a dentist
// @Accept json
// @Produce json
// @Param json body domain.Dentista true "Dentista"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas [post]
func (h *dentistaHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var d domain.Dentista
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		err := c.ShouldBindJSON(&d)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&d)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		//valid, err = validateExpiration(dentista.Matricula)
		//if !valid {
		//	web.Failure(c, 400, err)
		//	return
		//}
		p, err := h.s.Create(d)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete deleta um dentista
// @Summary Delete a dentist by ID
// @Tags CRUD DENTISTAS
// @Description Delete a dentist by ID
// @Accept json
// @Produce json
// @Param id path int true "Dentista ID"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/{id} [delete]
func (h *dentistaHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
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

// Put actualiza un dentista
// @Summary Put a dentist
// @Tags CRUD DENTISTAS
// @Description Put a dentist
// @Accept json
// @Produce json
// @Param id path int true "Dentista ID"
// @Param json body domain.Dentista true "Dentista"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/{id} [put]
func (h *dentistaHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var dentist domain.Dentista
		err = c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		//valid, err = validateExpiration(dentist.Matricula)
		//if !valid {
		//	web.Failure(c, 400, err)
		//	return
		//}
		p, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch actualiza un dentista ou algum de seus campos
// @Summary Patch a dentist
// @Tags CRUD DENTISTAS
// @Description Patch a dentist
// @Accept json
// @Produce json
// @Param id path int true "Dentista ID"
// @Param json body domain.Dentista true "Dentista"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/{id} [patch]
func (h *dentistaHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Nome      string `json:"nome,omitempty"`
		Sobrenome string `json:"sobrenome,omitempty"`
		Matricula int    `json:"matricula,omitempty"`
	}
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentista{
			Nome:      r.Nome,
			Sobrenome: r.Sobrenome,
			Matricula: r.Matricula,
		}
		//if update.Matricula != "" {
		//	valid, err := validateExpiration(update.Matricula)
		//	if !valid {
		//		web.Failure(c, 400, err)
		//		return
		//	}
		//}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
