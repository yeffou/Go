package contexts

import (
	"context"
	"final_project/pkg/models"
	"final_project/pkg/stores"
)

type CustomerContext struct {
	Store stores.CustomerStore
}

func NewCustomerContext(store stores.CustomerStore) *CustomerContext {
	return &CustomerContext{Store: store}
}

func (ctx *CustomerContext) CreateCustomer(ctxContext context.Context, customer models.Customer) models.Customer {
	return ctx.Store.CreateCustomer(customer)
}

func (ctx *CustomerContext) GetCustomer(ctxContext context.Context, id int) (models.Customer, bool) {
	return ctx.Store.GetCustomer(id)
}

func (ctx *CustomerContext) UpdateCustomer(ctxContext context.Context, id int, customer models.Customer) bool {
	return ctx.Store.UpdateCustomer(id, customer)
}

func (ctx *CustomerContext) DeleteCustomer(ctxContext context.Context, id int) bool {
	return ctx.Store.DeleteCustomer(id)
}

func (ctx *CustomerContext) ListCustomers(ctxContext context.Context) []models.Customer {
	return ctx.Store.ListCustomers()
}
