package main

import (
	"final_project/pkg/contexts"
	"final_project/pkg/handlers"
	"final_project/pkg/scheduler"
	"final_project/pkg/stores"
	"final_project/utils"
	"log"
	"net/http"
)

func main() {
	utils.Init()

	authorFilePath := "authors.json"
	bookFilePath := "books.json"
	customerFilePath := "customers.json"
	orderFilePath := "orders.json"

	authorStore := stores.NewAuthorStore(authorFilePath)
	bookStore := stores.NewBookStore(bookFilePath)
	customerStore := stores.NewCustomerStore(customerFilePath)
	orderStore := stores.NewOrderStore(orderFilePath)

	authorContext := contexts.NewAuthorContext(authorStore)
	bookContext := contexts.NewBookContext(bookStore)
	customerContext := contexts.NewCustomerContext(customerStore)
	orderContext := contexts.NewOrderContext(orderStore)
	reportContext := contexts.NewReportContext(orderStore)

	authorHandler := handlers.NewAuthorHandler(authorContext)
	bookHandler := handlers.NewBookHandler(bookContext)
	customerHandler := handlers.NewCustomerHandler(customerContext)
	orderHandler := handlers.NewOrderHandler(orderContext)
	reportHandler := handlers.NewReportHandler(reportContext)

	scheduler.StartDailyReportJob(reportContext)

	// Author routes
	http.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			authorHandler.HandleCreateAuthor(w, r)
		case http.MethodGet:
			authorHandler.HandleGetAuthors(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/authors/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			authorHandler.HandleGetAuthor(w, r)
		case http.MethodPut:
			authorHandler.HandleUpdateAuthor(w, r)
		case http.MethodDelete:
			authorHandler.HandleDeleteAuthor(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Book routes
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			bookHandler.HandleCreateBook(w, r)
		case http.MethodGet:
			bookHandler.HandleGetBooks(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			bookHandler.HandleGetBook(w, r)
		case http.MethodPut:
			bookHandler.HandleUpdateBook(w, r)
		case http.MethodDelete:
			bookHandler.HandleDeleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Customer routes
	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			customerHandler.HandleCreateCustomer(w, r)
		case http.MethodGet:
			customerHandler.HandleGetCustomers(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/customers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			customerHandler.HandleGetCustomer(w, r)
		case http.MethodPut:
			customerHandler.HandleUpdateCustomer(w, r)
		case http.MethodDelete:
			customerHandler.HandleDeleteCustomer(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Order routes
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			orderHandler.HandleCreateOrder(w, r)
		case http.MethodGet:
			orderHandler.HandleGetOrders(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			orderHandler.HandleGetOrder(w, r)
		case http.MethodPut:
			orderHandler.HandleUpdateOrder(w, r)
		case http.MethodDelete:
			orderHandler.HandleDeleteOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Report route
	http.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		reportHandler.HandleGenerateReport(w, r)
	})

	port := "8080"
	log.Printf("Server is running on port :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
