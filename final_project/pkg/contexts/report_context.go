package contexts

import (
	"context"
	"final_project/pkg/models"
	"final_project/pkg/stores"
	"time"
)

type ReportContext struct {
	OrderStore stores.OrderStore
}

func NewReportContext(orderStore stores.OrderStore) *ReportContext {
	return &ReportContext{OrderStore: orderStore}
}

func (ctx *ReportContext) GenerateSalesReport(ctxContext context.Context, from, to time.Time) (models.SalesReport, error) {
	orders := ctx.OrderStore.ListOrders()

	totalRevenue := 0.0
	totalOrders := 0
	bookSales := make(map[int]int)
	books := make(map[int]models.Book)

	for _, order := range orders {
		if order.CreatedAt.After(from) && order.CreatedAt.Before(to) {
			totalRevenue += order.TotalPrice
			totalOrders++

			for _, item := range order.Items {
				bookSales[item.Book.ID] += item.Quantity
				books[item.Book.ID] = item.Book
			}
		}
	}

	topSellingBooks := []models.BookSales{}
	for bookID, quantity := range bookSales {
		topSellingBooks = append(topSellingBooks, models.BookSales{
			Book:     books[bookID],
			Quantity: quantity,
		})
	}

	report := models.SalesReport{
		Timestamp:       time.Now(),
		TotalRevenue:    totalRevenue,
		TotalOrders:     totalOrders,
		TopSellingBooks: topSellingBooks,
	}

	return report, nil
}
