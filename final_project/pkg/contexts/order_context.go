package contexts

import (
	"context"
	"final_project/pkg/models"
	"final_project/pkg/stores"
	"time"
)

type OrderContext struct {
	Store stores.OrderStore
}

func NewOrderContext(store stores.OrderStore) *OrderContext {
	return &OrderContext{Store: store}
}

func (ctx *OrderContext) CreateOrder(ctxContext context.Context, order models.Order) (models.Order, bool) {
	return ctx.Store.CreateOrder(order)
}

func (ctx *OrderContext) GetOrder(ctxContext context.Context, id int) (models.Order, bool) {
	return ctx.Store.GetOrder(id)
}

func (ctx *OrderContext) UpdateOrder(ctxContext context.Context, id int, order models.Order) bool {
	return ctx.Store.UpdateOrder(id, order)
}

func (ctx *OrderContext) DeleteOrder(ctxContext context.Context, id int) bool {
	return ctx.Store.DeleteOrder(id)
}

func (ctx *OrderContext) ListOrders(ctxContext context.Context) []models.Order {
	return ctx.Store.ListOrders()
}

func (ctx *OrderContext) GetOrdersByDateRange(ctxContext context.Context, from, to time.Time) []models.Order {
	return ctx.Store.GetOrdersByDateRange(from, to)
}
