package stores

import (
	"encoding/json"
	"final_project/pkg/models"
	"fmt"
	"os"
	"sync"
	"time"
)

type orderStore struct {
	mu       sync.Mutex
	orders   map[int]models.Order
	nextID   int
	filePath string
}

func NewOrderStore(filePath string) OrderStore {
	store := &orderStore{
		orders:   make(map[int]models.Order),
		nextID:   1,
		filePath: filePath,
	}
	if err := store.LoadFromFile(filePath); err != nil {
		fmt.Printf("Warning: Could not load orders from file: %v\n", err)
	}
	return store
}

func (s *orderStore) CreateOrder(order models.Order) (models.Order, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order.ID = s.nextID
	order.CreatedAt = time.Now()
	s.orders[s.nextID] = order
	s.nextID++

	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving orders in CreateOrder: %v\n", err)
		return models.Order{}, false
	}
	return order, true
}

func (s *orderStore) GetOrder(id int) (models.Order, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, exists := s.orders[id]
	return order, exists
}

func (s *orderStore) UpdateOrder(id int, order models.Order) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.orders[id]; !exists {
		return false
	}
	order.ID = id
	s.orders[id] = order

	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving orders in UpdateOrder: %v\n", err)
	}
	return true
}

func (s *orderStore) DeleteOrder(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.orders[id]; !exists {
		return false
	}
	delete(s.orders, id)

	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving orders in DeleteOrder: %v\n", err)
	}
	return true
}

func (s *orderStore) ListOrders() []models.Order {
	s.mu.Lock()
	defer s.mu.Unlock()

	orders := []models.Order{}
	for _, order := range s.orders {
		orders = append(orders, order)
	}
	return orders
}

func (s *orderStore) GetOrdersByDateRange(from, to time.Time) []models.Order {
	s.mu.Lock()
	defer s.mu.Unlock()

	orders := []models.Order{}
	for _, order := range s.orders {
		if order.CreatedAt.After(from) && order.CreatedAt.Before(to) {
			orders = append(orders, order)
		}
	}
	return orders
}

func (s *orderStore) SaveToFile(filePath string) error {

	data, err := json.MarshalIndent(s.orders, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal orders: %v", err)
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write orders to file: %v", err)
	}
	return nil
}

func (s *orderStore) LoadFromFile(filePath string) error {

	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		s.orders = make(map[int]models.Order)
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.orders); err != nil {
		return fmt.Errorf("failed to decode orders: %v", err)
	}

	for id := range s.orders {
		if id >= s.nextID {
			s.nextID = id + 1
		}
	}
	return nil
}
