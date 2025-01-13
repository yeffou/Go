package scheduler

import (
	"context"
	"encoding/json"
	"final_project/pkg/contexts"
	"final_project/utils"
	"log"
	"os"
	"path/filepath"
	"time"
)

func StartDailyReportJob(reportContext *contexts.ReportContext) {
	go func() {
		for {
			now := time.Now()
			nextRun := now.Truncate(24 * time.Hour).Add(24 * time.Hour)
			if now.After(nextRun) {
				nextRun = nextRun.Add(24 * time.Hour)
			}

			timeUntilNextRun := nextRun.Sub(now)
			utils.LogInfo("Waiting until next report generation: " + timeUntilNextRun.String())
			time.Sleep(timeUntilNextRun)

			from := now.Truncate(24 * time.Hour).Add(-24 * time.Hour)
			to := now.Truncate(24 * time.Hour)

			utils.LogInfo("Generating daily sales report...")
			report, err := reportContext.GenerateSalesReport(context.Background(), from, to)
			if err != nil {
				utils.LogError(err)
				continue
			}

			os.MkdirAll("output-reports", os.ModePerm)
			filename := "report_" + report.Timestamp.Format("20060102_150405") + ".json"
			filePath := filepath.Join("output-reports", filename)

			data, err := json.MarshalIndent(report, "", "  ")
			if err != nil {
				utils.LogError(err)
				continue
			}

			if err := os.WriteFile(filePath, data, 0644); err != nil {
				utils.LogError(err)
				continue
			}

			log.Println("Daily report generated successfully:", filePath)
		}
	}()
}
