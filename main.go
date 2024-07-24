package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/thiagocamargodacosta/dass-21/exporter"
	"github.com/thiagocamargodacosta/dass-21/form"
	"github.com/thiagocamargodacosta/dass-21/reporting"
	"github.com/thiagocamargodacosta/dass-21/scoring"
)

const filename = "DASS-21 (respostas) - Respostas ao formulário 1.csv"
const output = "dass-21-scoring-results.csv"

func main() {

	// Open input csv
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Read the csv file
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	// Convert csv to slice
	forms := form.CreateForms(data)

	reports := make([]reporting.Report, len(forms))

	for i, entry := range forms {

		var s scoring.Score = scoring.DASS(entry)
		var r reporting.Report = reporting.Report{
			Date:       entry.Date,
			Email:      entry.Email,
			Depression: s.Depression,
			Anxiety:    s.Anxiety,
			Stress:     s.Stress,
		}

		reporting.Print(r)

		reports[i] = r
	}

	err = exporter.WriteOutputCSV(output, reports)

	if err != nil {
		log.Fatalf("error while writing output csv", err)
	}

}
