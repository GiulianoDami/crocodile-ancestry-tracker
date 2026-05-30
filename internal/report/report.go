package report

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Result represents the analysis result for a specimen
type Result struct {
	SpecimenID   string
	Species      string
	Confidence   float64
	Lineage      string
	Migration    string
	Notes        string
}

// GenerateReport generates a lineage report from analyzed data
func GenerateReport(results []Result, outputPath string) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create report file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Specimen ID", "Species", "Confidence (%)", "Lineage", "Migration Pattern", "Notes"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write headers: %w", err)
	}

	for _, result := range results {
		row := []string{
			result.SpecimenID,
			result.Species,
			fmt.Sprintf("%.2f", result.Confidence),
			result.Lineage,
			result.Migration,
			strings.ReplaceAll(result.Notes, "\n", " "),
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write result row: %w", err)
		}
	}

	return nil
}