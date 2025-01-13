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

type OrderHandler struct {
	Context *contexts.OrderContext
}

func NewOrderHandler(context *contexts.OrderContext) *OrderHandler {
	return &OrderHandler{Context: context}
}

func (h *OrderHandler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleCreateOrder called")
	if r.Method != http.MethodPost {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid input")
		utils.LogError(err)
		return
	}

	created, ok := h.Context.CreateOrder(r.Context(), order)
	if !ok {
		WriteErrorResponse(w, http.StatusBadRequest, "Failed to create order")
		utils.LogError(nil)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
	utils.LogInfo("Order created successfully")
}

func (h *OrderHandler) HandleGetOrders(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleGetOrders called")
	if r.Method != http.MethodGet {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	orders := h.Context.ListOrders(r.Context())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
	utils.LogInfo("Orders retrieved successfully")
}

func (h *OrderHandler) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleGetOrder called")
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
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid order ID")
		utils.LogError(err)
		return
	}

	order, found := h.Context.GetOrder(r.Context(), id)
	if !found {
		WriteErrorResponse(w, http.StatusNotFound, "Order not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
	utils.LogInfo("Order retrieved successfully")
}

func (h *OrderHandler) HandleUpdateOrder(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleUpdateOrder called")
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
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid order ID")
		utils.LogError(err)
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid input")
		utils.LogError(err)
		return
	}

	success := h.Context.UpdateOrder(r.Context(), id, order)
	if !success {
		WriteErrorResponse(w, http.StatusNotFound, "Order not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
	utils.LogInfo("Order updated successfully")
}

func (h *OrderHandler) HandleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleDeleteOrder called")
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
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid order ID")
		utils.LogError(err)
		return
	}

	success := h.Context.DeleteOrder(r.Context(), id)
	if !success {
		WriteErrorResponse(w, http.StatusNotFound, "Order not found")
		utils.LogError(nil)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	utils.LogInfo("Order deleted successfully")
}
