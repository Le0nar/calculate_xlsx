package service

import (
	"errors"
	"mime/multipart"
	"runtime"
	"strconv"
	"sync"

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

	capital := calculateCapital(rows)

	portfolio := portfolio.Portfolio{
		UserID:  id,
		Capital: capital,
	}

	return &portfolio, nil
}

func calculateCapital(rows [][]string) float64 {
	var capital float64
	var mu sync.Mutex
	var wg sync.WaitGroup

	numGoroutines := runtime.NumCPU() - 1
	rowsPerGoroutines := len(rows) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			firstIndex := i * rowsPerGoroutines
			lastIndex := firstIndex + rowsPerGoroutines

			if i == numGoroutines-1 {
				lastIndex = len(rows)
			}

			var localCouner float64

			for _, row := range rows[firstIndex:lastIndex] {
				for _, cell := range row {
					value, err := strconv.ParseFloat(cell, 64)
					if err == nil {
						localCouner += value
					}

				}
			}

			mu.Lock()
			capital += localCouner
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	return capital
}

func (s *Service) GetPortfolio(id uuid.UUID) (*portfolio.Portfolio, error) {
	return nil, nil
}
