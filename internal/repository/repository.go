package repository

import (
	"database/sql"
	"fmt"

	"github.com/Le0nar/calculate_xlsx/internal/portfolio"
	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreatePortfolio(portf *portfolio.Portfolio) error {
	// Строка запроса для вставки данных
	query := `INSERT INTO portfolio (user_id, capital)  VALUES ($1, $2)`

	// Выполняем запрос с параметрами: user_id и capital
	_, err := r.db.Exec(query, portf.UserID, portf.Capital)
	if err != nil {
		return fmt.Errorf("failed to insert portfolio: %w", err)
	}

	fmt.Printf("WRITE TO DB: %v\n", portf)

	return nil
}
func (r *Repository) GetPortfolioById(id uuid.UUID) (*portfolio.Portfolio, error) {
	query := `
        SELECT user_id, capital
        FROM portfolio
        WHERE user_id = $1
        LIMIT 1
    `
	row := r.db.QueryRow(query, id)

	var portfiolio portfolio.Portfolio
	if err := row.Scan(&portfiolio.UserID, &portfiolio.Capital); err != nil {
		if err == sql.ErrNoRows {
			// Не возвращаем ошибку, так как это стандартный сценарий при первой отправке user id
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query portfolio: %w", err)
	}

	fmt.Printf("READ FROM DB: %v\n", portfiolio)

	return &portfiolio, nil
}
