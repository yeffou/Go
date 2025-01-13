package stores

import (
	"encoding/json"
	"final_project/pkg/models"
	"fmt"
	"os"
	"sync"
)

type authorStore struct {
	mu       sync.Mutex
	authors  map[int]models.Author
	nextID   int
	filePath string
}

func NewAuthorStore(filePath string) AuthorStore {
	store := &authorStore{
		authors:  make(map[int]models.Author),
		nextID:   1,
		filePath: filePath,
	}
	if err := store.LoadFromFile(filePath); err != nil {
		fmt.Printf("Warning: Could not load authors from file: %v\n", err)
	}
	return store
}

func (s *authorStore) CreateAuthor(author models.Author) models.Author {
	s.mu.Lock()

	author.ID = s.nextID
	s.authors[s.nextID] = author
	s.nextID++
	defer s.mu.Unlock()

	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving to file after CreateAuthor: %v\n", err)
	}
	return author
}

func (s *authorStore) GetAuthor(id int) (models.Author, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	author, exists := s.authors[id]
	return author, exists
}

func (s *authorStore) UpdateAuthor(id int, author models.Author) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.authors[id]; !exists {
		return false
	}
	author.ID = id
	s.authors[id] = author

	// Persist data after update
	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving to file after UpdateAuthor: %v\n", err)
	}
	return true
}

func (s *authorStore) DeleteAuthor(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.authors[id]; !exists {
		return false
	}
	delete(s.authors, id)

	if err := s.SaveToFile(s.filePath); err != nil {
		fmt.Printf("Error saving to file after DeleteAuthor: %v\n", err)
	}
	return true
}

func (s *authorStore) ListAuthors() []models.Author {
	s.mu.Lock()
	defer s.mu.Unlock()

	authors := []models.Author{}
	for _, author := range s.authors {
		authors = append(authors, author)
	}
	return authors
}

func (s *authorStore) SaveToFile(filePath string) error {

	data, err := json.MarshalIndent(s.authors, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal authors: %v", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write authors to file: %v", err)
	}

	return nil
}

func (s *authorStore) LoadFromFile(filePath string) error {

	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		s.authors = make(map[int]models.Author)
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.authors); err != nil {
		return fmt.Errorf("failed to decode authors: %v", err)
	}

	for id := range s.authors {
		if id >= s.nextID {
			s.nextID = id + 1
		}
	}
	return nil
}
