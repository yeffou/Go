# Final Project Documentation

## Project Overview

This project implements a book-ordering API with the following features:

1. **CRUD Operations** for:
   - Authors
   - Books
   - Customers
   - Orders
2. **Search and Filtering**:
   - Retrieve books using search criteria.
   - Retrieve orders based on a date range.
3. **Periodic Sales Report Generation**:
   - Background task generates reports every 30 seconds.
   - Reports include total revenue, number of orders, books sold, and top-selling books.
4. **Error Handling**:
   - Consistent error responses in JSON format.
5. **Logging**:
   - Logs API events and errors to the console.
6. **API Documentation** using Swagger/OpenAPI.

---

## How to Run the Project

### Prerequisites

- Go installed on your machine.
- A tool like Postman or `curl` for testing endpoints.
- (Optional) `swagger-ui` for visualizing API documentation.

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/yeffou/GO
   cd final_project
   ```

2. Start the server:

   ```bash
   go run main.go
   ```

3. The server will run at `http://localhost:8080`.

---

## API Endpoints

### Authors

| Method | Endpoint        | Description                 |
| ------ | --------------- | --------------------------- |
| POST   | `/authors`      | Create a new author.        |
| GET    | `/authors`      | Retrieve all authors.       |
| GET    | `/authors/{id}` | Retrieve a specific author. |
| PUT    | `/authors/{id}` | Update a specific author.   |
| DELETE | `/authors/{id}` | Delete a specific author.   |

#### Example Request (Create Author):

```json
POST /authors
{
  "first_name": "Jaafar",
  "last_name": "Yeffou",
  "bio": "Computer Science Student"
}
```

---

### Books

| Method | Endpoint      | Description               |
| ------ | ------------- | ------------------------- |
| POST   | `/books`      | Create a new book.        |
| GET    | `/books`      | Retrieve all books.       |
| GET    | `/books/{id}` | Retrieve a specific book. |
| PUT    | `/books/{id}` | Update a specific book.   |
| DELETE | `/books/{id}` | Delete a specific book.   |

### Customers

| Method | Endpoint          | Description                   |
| ------ | ----------------- | ----------------------------- |
| POST   | `/customers`      | Create a new customer.        |
| GET    | `/customers`      | Retrieve all customers.       |
| GET    | `/customers/{id}` | Retrieve a specific customer. |
| PUT    | `/customers/{id}` | Update a specific customer.   |
| DELETE | `/customers/{id}` | Delete a specific customer.   |

---

### Orders

| Method | Endpoint           | Description                          |
| ------ | ------------------ | ------------------------------------ |
| POST   | `/orders`          | Create a new order.                  |
| GET    | `/orders`          | Retrieve all orders.                 |
| GET    | `/orders/{id}`     | Retrieve a specific order.           |
| PUT    | `/orders/{id}`     | Update a specific order.             |
| DELETE | `/orders/{id}`     | Delete a specific order.             |
| GET    | `/orders?from=to=` | Retrieve orders within a date range. |

### Reports

| Method | Endpoint            | Description                       |
| ------ | ------------------- | --------------------------------- |
| GET    | `/reports?from=to=` | Retrieve generated sales reports. |

#### Example Request (Retrieve Reports):

```json
GET "http://localhost:8080/report?from=2025-01-12T00:00:00Z&to=2025-01-13T02:00:00Z"
```

---

## Periodic Report Generation

- The periodic report job runs every 24 hours.
- Reports are saved in the `output-reports/` directory with filenames based on the timestamp.

---

## Logs

- Logs are printed to the console with the following format:
  - **INFO**: Logs successful events.
  - **ERROR**: Logs errors during execution.

---

## Testing

1. Use Postman or `curl` to test endpoints.
2. Verify the periodic report generation by checking the `output-reports/` directory.
3. Ensure proper responses for both valid and invalid requests.

---

## Future Enhancements

- Add authentication and authorization.
- Optimize search and filtering.
- Extend periodic reporting to support weekly/monthly summaries.

---

## Contributors

- **Jaafar YEFFOU**: [jaafar.yeffou@um6p.ma]

---
