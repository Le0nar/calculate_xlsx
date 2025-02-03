package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Le0nar/calculate_xlsx/internal/portfolio"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

const redisAddr = "localhost:6379"

type Cache struct {
	redisClient *redis.Client
}

func NewCache() *Cache {
	// Инициализируем Redis клиент
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr, // Например, "localhost:6379"
	})

	// Проверяем подключение к Redis
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to Redis: %v", err))
	}

	return &Cache{
		redisClient: rdb,
	}
}

func (c *Cache) CreatePortfolio(portf *portfolio.Portfolio) error {
	fmt.Printf("portf: %v\n", portf)

	// Преобразуем Portfolio в строку JSON (например, с помощью encoding/json)
	data := fmt.Sprintf(`{"user_id": "%s", "capital": %f}`, portf.UserID.String(), portf.Capital)

	fmt.Printf("data: %v\n", data)

	// Записываем портфель в Redis с ключом по user_id и сроком хранения 1 час (3600 секунд)
	err := c.redisClient.Set(portf.UserID.String(), data, 3600*time.Second).Err()
	if err != nil {
		return fmt.Errorf("failed to save portfolio to Redis: %w", err)
	}

	fmt.Printf("WRITE TO REDIS: %v\n", data)

	return nil
}

func (c *Cache) GetPortfolioById(userID uuid.UUID) (*portfolio.Portfolio, error) {
	// Ищем данные в Redis по ключу user_id
	data, err := c.redisClient.Get(userID.String()).Result()
	if err == redis.Nil {
		// Если данные не найдены не выводим ошибку, так как мы ожидаем такое поведение
		return nil, nil
		// return nil, fmt.Errorf("portfolio not found in cache for user_id: %v", userID)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get portfolio from Redis: %w", err)
	}

	var portf portfolio.Portfolio

	err = json.Unmarshal([]byte(data), &portf)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal portfolio data: %w", err)
	}

	fmt.Printf("READ FROM REDIS: %v\n", portf)

	return &portf, nil
}
