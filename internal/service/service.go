package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"runtime"
	"strconv"
	"sync"

	"github.com/Le0nar/calculate_xlsx/internal/portfolio"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

type repository interface {
	CreatePortfolio(portf *portfolio.Portfolio) error
	GetPortfolioById(userID uuid.UUID) (*portfolio.Portfolio, error)
}

type cache interface {
	CreatePortfolio(portf *portfolio.Portfolio) error
	GetPortfolioById(userID uuid.UUID) (*portfolio.Portfolio, error)
}

// TODO: add Cache field
type Service struct {
	Repository repository
	// TODO: use interface
	Cache cache
}

func NewService(repo repository) *Service {
	return &Service{
		Repository: repo,
		Cache:      NewCache(),
	}
}

func (s *Service) CreatePortfolio(id uuid.UUID, file *multipart.File) (*portfolio.Portfolio, error) {
	// 1) Look for cache
	portf, err := s.Cache.GetPortfolioById(id)
	if err != nil {
		fmt.Printf("cache read error: %s \n", err.Error())
	}

	if portf != nil {
		return portf, nil
	}

	// 2) Look for DB
	portf, err = s.Repository.GetPortfolioById(id)
	if err != nil {
		fmt.Printf("data base read error: %s \n", err.Error())
	}

	if portf != nil {
		err = s.Cache.CreatePortfolio(portf)
		if err != nil {
			fmt.Printf("write to cache error: %s \n", err.Error())
		}

		return portf, nil
	}

	// 3) Calculate
	capital, err := calculateCapital(file)

	if err != nil {
		return nil, err
	}

	portf = &portfolio.Portfolio{
		UserID:  id,
		Capital: capital,
	}

	err = s.Repository.CreatePortfolio(portf)

	if err != nil {
		fmt.Printf("data base write error: %s \n", err.Error())
	}

	err = s.Cache.CreatePortfolio(portf)
	if err != nil {
		fmt.Printf("write to cache error: %s \n", err.Error())
	}

	return portf, nil
}

func calculateCapital(file *multipart.File) (float64, error) {
	// Открываем файл Excel с помощью excelize
	f, err := excelize.OpenReader(*file)
	if err != nil {
		return 0, errors.New("ошибка при чтении файла Excel")
	}

	// Получаем все имена листов в файле
	sheetNames := f.GetSheetList()
	if len(sheetNames) == 0 {
		return 0, errors.New("листы не найдены в файле")
	}

	// Читаем первый лист
	sheet := sheetNames[0] // Берем имя первого листа
	rows, err := f.GetRows(sheet)
	if err != nil {
		return 0, errors.New("ошибка при получении строк")
	}

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

	return capital, nil
}
