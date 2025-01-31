package handler

import (
	"fmt"
	"net/http"

	"github.com/Le0nar/calculate_xlsx/internal/portfolio"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

type service interface {
	CalculatePortfolio(dto portfolio.CaclulatePortfolioDto) error
	GetPortfolio(id uuid.UUID) (*portfolio.Portfolio, error)
}

type Handler struct {
	service service
}

func NewHandler(s service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) calculatePortfolio(c *gin.Context) {
	// Получаем файл из формы
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при загрузке файла"})
		return
	}

	// Открываем файл Excel с помощью excelize
	f, err := excelize.OpenReader(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при чтении файла Excel"})
		return
	}

	// Получаем все имена листов в файле
	sheetNames := f.GetSheetList()
	if len(sheetNames) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Листы не найдены в файле"})
		return
	}

	// Читаем первый лист
	sheet := sheetNames[0] // Берем имя первого листа
	rows, err := f.GetRows(sheet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении строк"})
		return
	}

	// TODO: добавить калькуляцию с горутинами здесь
	// Пример обработки строк
	for _, row := range rows {
		for _, cell := range row {
			fmt.Println(cell)
		}
	}

	// Ответ
	c.JSON(http.StatusOK, gin.H{"message": "Файл успешно обработан"})

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
