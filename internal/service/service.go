package service

import (
	"errors"
	"mime/multipart"
	"strconv"

	"github.com/Le0nar/calculate_xlsx/internal/portfolio"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CalculatePortfolio(id uuid.UUID, file *multipart.File) (*portfolio.Portfolio, error) {
	// Открываем файл Excel с помощью excelize
	f, err := excelize.OpenReader(*file)
	if err != nil {
		return nil, errors.New("ошибка при чтении файла Excel")
	}

	// Получаем все имена листов в файле
	sheetNames := f.GetSheetList()
	if len(sheetNames) == 0 {
		return nil, errors.New("листы не найдены в файле")
	}

	// Читаем первый лист
	sheet := sheetNames[0] // Берем имя первого листа
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, errors.New("ошибка при получении строк")
	}

	portfolio := portfolio.Portfolio{
		UserID:  id,
		Capital: 0,
	}

	// TODO: добавить калькуляцию с горутинами здесь
	// Пример обработки строк
	for _, row := range rows {
		for _, cell := range row {
			value, err := strconv.ParseFloat(cell, 64)

			if err == nil {
				portfolio.Capital += value
			}
		}
	}

	return &portfolio, nil
}

func (s *Service) GetPortfolio(id uuid.UUID) (*portfolio.Portfolio, error) {
	return nil, nil
}
