package handler

import (
	"mime/multipart"
	"net/http"

	"github.com/Le0nar/calculate_xlsx/internal/portfolio"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service interface {
	CalculatePortfolio(id uuid.UUID, file *multipart.File) (*portfolio.Portfolio, error)
	GetPortfolio(id uuid.UUID) (*portfolio.Portfolio, error)
}

type Handler struct {
	service service
}

func NewHandler(s service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) calculatePortfolio(c *gin.Context) {
	stringedId := c.Param("id")

	id, err := uuid.Parse(stringedId)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при загрузке файла"})
		return
	}

	portfolio, err := h.service.CalculatePortfolio(id, &file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"portfolio": portfolio})
}

func (h *Handler) getPortfolio(c *gin.Context) {

}

func (h *Handler) InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		portfolio := api.Group("/portfolio")
		{
			portfolio.POST("/:id", h.calculatePortfolio)
			portfolio.GET("/:id", h.getPortfolio)
		}
	}

	return r
}
