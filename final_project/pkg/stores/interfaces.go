package stores

import (
	"final_project/pkg/models"
	"time"
)

type AuthorStore interface {
	CreateAuthor(author models.Author) models.Author
	GetAuthor(id int) (models.Author, bool)
	UpdateAuthor(id int, author models.Author) bool
	DeleteAuthor(id int) bool
	ListAuthors() []models.Author
	SaveToFile(filePath string) error
	LoadFromFile(filePath string) error
}

type BookStore interface {
	CreateBook(models.Book) models.Book
	GetBook(int) (models.Book, bool)
	UpdateBook(int, models.Book) bool
	DeleteBook(int) bool
	ListBooks() []models.Book
	SaveToFile(filePath string) error
	LoadFromFile(filePath string) error
	SearchBooksByCriteria(criteria models.SearchCriteria) []models.Book
}

type CustomerStore interface {
	CreateCustomer(models.Customer) models.Customer
	GetCustomer(int) (models.Customer, bool)
	UpdateCustomer(int, models.Customer) bool
	DeleteCustomer(int) bool
	ListCustomers() []models.Customer
	SaveToFile(filePath string) error
	LoadFromFile(filePath string) error
}

type OrderStore interface {
	CreateOrder(order models.Order) (models.Order, bool)
	GetOrder(id int) (models.Order, bool)
	UpdateOrder(id int, order models.Order) bool
	DeleteOrder(id int) bool
	ListOrders() []models.Order
	GetOrdersByDateRange(from, to time.Time) []models.Order
	SaveToFile(filePath string) error
	LoadFromFile(filePath string) error
}
