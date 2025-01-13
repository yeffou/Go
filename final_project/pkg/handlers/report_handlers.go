package handlers

import (
	"encoding/json"
	"final_project/pkg/contexts"
	"final_project/utils"
	"net/http"
	"time"
)

type ReportHandler struct {
	Context *contexts.ReportContext
}

func NewReportHandler(context *contexts.ReportContext) *ReportHandler {
	return &ReportHandler{Context: context}
}

func (h *ReportHandler) HandleGenerateReport(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("HandleGenerateReport called")
	if r.Method != http.MethodGet {
		WriteErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		utils.LogError(nil)
		return
	}

	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	if fromStr == "" || toStr == "" {
		WriteErrorResponse(w, http.StatusBadRequest, "Missing 'from' or 'to' query parameters")
		utils.LogError(nil)
		return
	}

	from, err := time.Parse(time.RFC3339, fromStr)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid 'from' date format")
		utils.LogError(err)
		return
	}

	to, err := time.Parse(time.RFC3339, toStr)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "Invalid 'to' date format")
		utils.LogError(err)
		return
	}

	report, err := h.Context.GenerateSalesReport(r.Context(), from, to)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to generate report")
		utils.LogError(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)
	utils.LogInfo("Sales report generated successfully")
}
