package handlers

import (
	"encoding/json"
	"final_project/pkg/contexts"
	"final_project/pkg/models"
	"final_project/utils"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	Context *contexts.BookContext
}

func NewBookHandler(context *contexts.BookContext) *BookHandler {
	return &BookHandler{Context: context}
}

func (h *BookHandler) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleCreateBook called")
	if r.Method != http.MethodPost {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid input")
		utils.LogError(err)
		return
	}

	created, ok := h.Context.CreateBook(r.Context(), book)
	if !ok {
		WriteErrorResponse(w, http.StatusBadRequest, "Failed to create book")
		utils.LogError(nil)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
	utils.LogInfo("Book created successfully")
}

func (h *BookHandler) HandleGetBooks(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleGetBooks called")
	if r.Method != http.MethodGet {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	books := h.Context.ListBooks(r.Context())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
	utils.LogInfo("Books retrieved successfully")
}

func (h *BookHandler) HandleGetBook(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleGetBook called")
	if r.Method != http.MethodGet {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) != 2 {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid path")
		utils.LogError(nil)
		return
	}

	id, err := strconv.Atoi(pathParts[1])
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid book ID")
		utils.LogError(err)
		return
	}

	book, found := h.Context.GetBook(r.Context(), id)
	if !found {
		WriteErrorResponse(w, http.StatusNotFound, "Book not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
	utils.LogInfo("Book retrieved successfully")
}

func (h *BookHandler) HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleUpdateBook called")
	if r.Method != http.MethodPut {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) != 2 {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid path")
		utils.LogError(nil)
		return
	}

	id, err := strconv.Atoi(pathParts[1])
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid book ID")
		utils.LogError(err)
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid input")
		utils.LogError(err)
		return
	}

	success := h.Context.UpdateBook(r.Context(), id, book)
	if !success {
		WriteErrorResponse(w, http.StatusNotFound, "Book not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
	utils.LogInfo("Book updated successfully")
}

func (h *BookHandler) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleDeleteBook called")
	if r.Method != http.MethodDelete {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) != 2 {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid path")
		utils.LogError(nil)
		return
	}

	id, err := strconv.Atoi(pathParts[1])
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid book ID")
		utils.LogError(err)
		return
	}

	success := h.Context.DeleteBook(r.Context(), id)
	if !success {
		WriteErrorResponse(w, http.StatusNotFound, "Book not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	utils.LogInfo("Book deleted successfully")
}
