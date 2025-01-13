package contexts

import (
	"context"
	"final_project/pkg/models"
	"final_project/pkg/stores"
)

type AuthorContext struct {
	Store stores.AuthorStore
}

func NewAuthorContext(store stores.AuthorStore) *AuthorContext {
	return &AuthorContext{Store: store}
}

func (ctx *AuthorContext) CreateAuthor(ctxContext context.Context, author models.Author) models.Author {
	return ctx.Store.CreateAuthor(author)
}

func (ctx *AuthorContext) GetAuthor(ctxContext context.Context, id int) (models.Author, bool) {
	return ctx.Store.GetAuthor(id)
}

func (ctx *AuthorContext) UpdateAuthor(ctxContext context.Context, id int, author models.Author) bool {
	return ctx.Store.UpdateAuthor(id, author)
}

func (ctx *AuthorContext) DeleteAuthor(ctxContext context.Context, id int) bool {
	return ctx.Store.DeleteAuthor(id)
}

func (ctx *AuthorContext) ListAuthors(ctxContext context.Context) []models.Author {
	return ctx.Store.ListAuthors()
}
