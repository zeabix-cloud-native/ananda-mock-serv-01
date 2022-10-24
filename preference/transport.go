package preference

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PreferenceDTO struct {
	ProfileID uint   `json:"profile-id"`
	Language  string `json:"language"`
}

type Handler interface {
	CreatePreference(c *gin.Context)
	GetPreference(c *gin.Context)
}

type handler struct {
	s Service
}

func NewHandler(s Service) Handler {
	return &handler{
		s: s,
	}
}

func (h *handler) CreatePreference(c *gin.Context) {
	var p PreferenceDTO
	if err := c.BindJSON(&p); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
		return
	}

	created, err := h.s.CreatePreference(&p)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusCreated, &created)
}

func (h *handler) GetPreference(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("profileId"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
		return
	}

	find, err := h.s.GetLanguagePreference(uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusOK, &find)
}
