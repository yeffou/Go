package utils

import (
	"encoding/json"
	"final_project/pkg/models"
	"fmt"
	"os"
)

func SaveReportToFile(filePath string, report models.SalesReport) error {

	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal report: %v", err)
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open report file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write report to file: %v", err)
	}
	return nil
}
