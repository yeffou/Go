package stores

import (
	"encoding/json"
	"final_project/pkg/models"
	"fmt"
	"os"
	"sync"
)

type customerStore struct {
	mu        sync.Mutex
	customers map[int]models.Customer
	nextID    int
	filePath  string
}

func NewCustomerStore(filePath string) CustomerStore {
	store := &customerStore{
		customers: make(map[int]models.Customer),
		nextID:    1,
		filePath:  filePath,
	}
	if err := store.LoadFromFile(filePath); err != nil {
		fmt.Printf("Warning: Could not load customers from file: %v\n", err)
	}
	return store
}

func (s *customerStore) CreateCustomer(customer models.Customer) models.Customer {
	s.mu.Lock()
	defer s.mu.Unlock()

	customer.ID = s.nextID
	s.customers[s.nextID] = customer
	s.nextID++

	// Save changes
	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving customers in CreateCustomer: %v\n", err)
	}
	return customer
}

func (s *customerStore) GetCustomer(id int) (models.Customer, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	customer, exists := s.customers[id]
	return customer, exists
}

func (s *customerStore) UpdateCustomer(id int, customer models.Customer) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.customers[id]; !exists {
		return false
	}
	customer.ID = id
	s.customers[id] = customer

	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving customers in UpdateCustomer: %v\n", err)
	}
	return true
}

func (s *customerStore) DeleteCustomer(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.customers[id]; !exists {
		return false
	}
	delete(s.customers, id)

	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving customers in DeleteCustomer: %v\n", err)
	}
	return true
}

func (s *customerStore) ListCustomers() []models.Customer {
	s.mu.Lock()
	defer s.mu.Unlock()

	customers := []models.Customer{}
	for _, customer := range s.customers {
		customers = append(customers, customer)
	}
	return customers
}

func (s *customerStore) SaveToFile(filePath string) error {

	data, err := json.MarshalIndent(s.customers, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal customers: %v", err)
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write customers to file: %v", err)
	}
	return nil
}

func (s *customerStore) LoadFromFile(filePath string) error {

	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		s.customers = make(map[int]models.Customer)
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.customers); err != nil {
		return fmt.Errorf("failed to decode customers: %v", err)
	}

	for id := range s.customers {
		if id >= s.nextID {
			s.nextID = id + 1
		}
	}
	return nil
}
