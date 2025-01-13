package contexts

import (
	"context"
	"final_project/pkg/models"
	"final_project/pkg/stores"
)

type BookContext struct {
	Store stores.BookStore
}

func NewBookContext(store stores.BookStore) *BookContext {
	return &BookContext{Store: store}
}

func (ctx *BookContext) CreateBook(ctxContext context.Context, book models.Book) (models.Book, bool) {
	createdBook := ctx.Store.CreateBook(book)
	if createdBook.ID == 0 {
		return models.Book{}, false
	}
	return createdBook, true
}

func (ctx *BookContext) GetBook(ctxContext context.Context, id int) (models.Book, bool) {
	return ctx.Store.GetBook(id)
}

func (ctx *BookContext) UpdateBook(ctxContext context.Context, id int, book models.Book) bool {
	return ctx.Store.UpdateBook(id, book)
}

func (ctx *BookContext) DeleteBook(ctxContext context.Context, id int) bool {
	return ctx.Store.DeleteBook(id)
}

func (ctx *BookContext) ListBooks(ctxContext context.Context) []models.Book {
	return ctx.Store.ListBooks()
}
