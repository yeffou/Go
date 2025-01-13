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

type AuthorHandler struct {
	Context *contexts.AuthorContext
}

func NewAuthorHandler(context *contexts.AuthorContext) *AuthorHandler {
	return &AuthorHandler{Context: context}
}

func (h *AuthorHandler) HandleCreateAuthor(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleCreateAuthor called")
	if r.Method != http.MethodPost {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid input")
		utils.LogError(err)
		return
	}

	created := h.Context.CreateAuthor(r.Context(), author)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
	utils.LogInfo("Author created successfully")
}

func (h *AuthorHandler) HandleGetAuthors(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleGetAuthors called")
	if r.Method != http.MethodGet {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	authors := h.Context.ListAuthors(r.Context())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authors)
	utils.LogInfo("Authors retrieved successfully")
}

func (h *AuthorHandler) HandleGetAuthor(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleGetAuthor called")
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
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid author ID")
		utils.LogError(err)
		return
	}

	author, found := h.Context.GetAuthor(r.Context(), id)
	if !found {
		WriteErrorResponse(w, http.StatusNotFound, "Author not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
	utils.LogInfo("Author retrieved successfully")
}

func (h *AuthorHandler) HandleUpdateAuthor(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleUpdateAuthor called")
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
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid author ID")
		utils.LogError(err)
		return
	}

	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid input")
		utils.LogError(err)
		return
	}

	success := h.Context.UpdateAuthor(r.Context(), id, author)
	if !success {
		WriteErrorResponse(w, http.StatusNotFound, "Author not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)
	utils.LogInfo("Author updated successfully")
}

func (h *AuthorHandler) HandleDeleteAuthor(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleDeleteAuthor called")
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
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid author ID")
		utils.LogError(err)
		return
	}

	success := h.Context.DeleteAuthor(r.Context(), id)
	if !success {
		WriteErrorResponse(w, http.StatusNotFound, "Author not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	utils.LogInfo("Author deleted successfully")
}
