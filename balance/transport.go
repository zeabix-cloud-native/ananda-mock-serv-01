package balance

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateAccountBalance(c *gin.Context)
	GetAccountBalance(c *gin.Context)
	GetBalanceByOwner(c *gin.Context)
	DebitAccountBalance(c *gin.Context)
}

type handler struct {
	s Service
}

type DebitRequest struct {
	Amount uint `json:"amount"`
}

func NewHandler(s Service) Handler {
	return &handler{
		s: s,
	}
}

func (h *handler) CreateAccountBalance(c *gin.Context) {
	var acc BalanceAccountDTO
	if err := c.BindJSON(&acc); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
		return
	}

	created, err := h.s.CreateBalanceAccount(&acc)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusCreated, &created)
}

func (h *handler) GetAccountBalance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
		return
	}

	find, err := h.s.Get(uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusOK, &find)
}

func (h *handler) GetBalanceByOwner(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
		return
	}

	find, err := h.s.GetByOwner(uint(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusOK, &find)
}

func (h *handler) DebitAccountBalance(c *gin.Context) {
	var req DebitRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{})
		return
	}

	dto, err := h.s.Debit(uint(id), req.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.IndentedJSON(http.StatusOK, &dto)
}
