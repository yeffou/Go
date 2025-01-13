package stores

import (
	"encoding/json"
	"final_project/pkg/models"
	"fmt"
	"os"
	"sync"
)

type bookStore struct {
	mu       sync.Mutex
	books    map[int]models.Book
	nextID   int
	filePath string
}

func NewBookStore(filePath string) BookStore {
	store := &bookStore{
		books:    make(map[int]models.Book),
		nextID:   1,
		filePath: filePath,
	}
	if err := store.LoadFromFile(filePath); err != nil {
		fmt.Printf("Warning: Could not load books from file: %v\n", err)
	}
	return store
}

func (s *bookStore) CreateBook(book models.Book) models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()

	book.ID = s.nextID
	s.books[s.nextID] = book
	s.nextID++

	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving books in CreateBook: %v\n", err)
	}
	return book
}

func (s *bookStore) GetBook(id int) (models.Book, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	book, exists := s.books[id]
	return book, exists
}

func (s *bookStore) UpdateBook(id int, book models.Book) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[id]; !exists {
		return false
	}
	book.ID = id
	s.books[id] = book

	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving books in UpdateBook: %v\n", err)
	}
	return true
}

func (s *bookStore) DeleteBook(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[id]; !exists {
		return false
	}
	delete(s.books, id)

	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving books in DeleteBook: %v\n", err)
	}
	return true
}

func (s *bookStore) ListBooks() []models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()

	books := []models.Book{}
	for _, book := range s.books {
		books = append(books, book)
	}
	return books
}

func (s *bookStore) SaveToFile(filePath string) error {

	data, err := json.MarshalIndent(s.books, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal books: %v", err)
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write books to file: %v", err)
	}
	return nil
}

func (s *bookStore) LoadFromFile(filePath string) error {

	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		s.books = make(map[int]models.Book)
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.books); err != nil {
		return fmt.Errorf("failed to decode books: %v", err)
	}

	for id := range s.books {
		if id >= s.nextID {
			s.nextID = id + 1
		}
	}
	return nil
}
