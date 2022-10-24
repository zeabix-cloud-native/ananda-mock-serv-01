package profile

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Handler interface {
	CreateProfile(c *gin.Context)
	GetProfile(c *gin.Context)
	GetProfileShallow(c *gin.Context)
}

type handler struct {
	s Service
}

func NewHandlers(s Service) Handler {
	return &handler{
		s: s,
	}
}

func (h *handler) CreateProfile(c *gin.Context) {
	var profile ProfileDTO
	if err := c.BindJSON(&profile); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
		return
	}

	created, err := h.s.CreateProfile(&profile)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusCreated, &created)
}

// Expect .../{id}
func (h *handler) GetProfile(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
		return
	}

	find, err := h.s.GetProfile(uint(id))
	if err == ErrProfileNotFound {
		log.Error(err.Error())
		c.IndentedJSON(http.StatusNotFound, gin.H{})
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
		log.Error(err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, &find)
}

func (h *handler) GetProfileShallow(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
		return
	}

	find, err := h.s.GetProfileShallow(uint(id))
	if err == ErrProfileNotFound {
		log.Error(err.Error())
		c.IndentedJSON(http.StatusNotFound, gin.H{})
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
		log.Error(err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, &find)
}
