package service

import (
	"github.com/Le0nar/calculate_xlsx/internal/portfolio"
	"github.com/google/uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CalculatePortfolio(dto portfolio.CaclulatePortfolioDto) error {
	return nil
}

func (s *Service) GetPortfolio(id uuid.UUID) (*portfolio.Portfolio, error) {
	return nil, nil
}
