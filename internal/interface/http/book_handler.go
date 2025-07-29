package http

import (
	"net/http"

	"github.com/MarceloBxD/goandk8s/internal/domain/book"
	book "github.com/MarceloBxD/goandk8s/internal/domain/book"
	"github.com/MarceloBxD/goandk8s/internal/usecase/book"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	CreateUC *book.CreateBookUseCase
}

func NewBookHandler(r *gin.Engine, uc *book.CreateBookUseCase) {
	handler := &BookHandler{CreateUC: uc}
	r.POST("/books", handler.Create)
}

func (h *BookHandler) Create(c *gin.Context) {
	var b book.Book 
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error", err.Error()})
		return 
	}

	if err:= h.CreateUC.Execute(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, b)
}