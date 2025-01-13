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

type CustomerHandler struct {
	Context *contexts.CustomerContext
}

func NewCustomerHandler(context *contexts.CustomerContext) *CustomerHandler {
	return &CustomerHandler{Context: context}
}

func (h *CustomerHandler) HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleCreateCustomer called")
	if r.Method != http.MethodPost {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid input")
		utils.LogError(err)
		return
	}

	created := h.Context.CreateCustomer(r.Context(), customer)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
	utils.LogInfo("Customer created successfully")
}

func (h *CustomerHandler) HandleGetCustomers(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleGetCustomers called")
	if r.Method != http.MethodGet {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	customers := h.Context.ListCustomers(r.Context())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
	utils.LogInfo("Customers retrieved successfully")
}

func (h *CustomerHandler) HandleGetCustomer(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleGetCustomer called")
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
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid customer ID")
		utils.LogError(err)
		return
	}

	customer, found := h.Context.GetCustomer(r.Context(), id)
	if !found {
		WriteErrorResponse(w, http.StatusNotFound, "Customer not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
	utils.LogInfo("Customer retrieved successfully")
}

func (h *CustomerHandler) HandleUpdateCustomer(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleUpdateCustomer called")
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
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid customer ID")
		utils.LogError(err)
		return
	}

	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid input")
		utils.LogError(err)
		return
	}

	success := h.Context.UpdateCustomer(r.Context(), id, customer)
	if !success {
		WriteErrorResponse(w, http.StatusNotFound, "Customer not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
	utils.LogInfo("Customer updated successfully")
}

func (h *CustomerHandler) HandleDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleDeleteCustomer called")
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
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid customer ID")
		utils.LogError(err)
		return
	}

	success := h.Context.DeleteCustomer(r.Context(), id)
	if !success {
		WriteErrorResponse(w, http.StatusNotFound, "Customer not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	utils.LogInfo("Customer deleted successfully")
}
