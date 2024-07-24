package exporter

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/thiagocamargodacosta/dass-21/reporting"
)

const (
	date_column       string = "Carimbo de data/hora"
	email_column      string = "Endereço de e-mail"
	depression_column string = "Classificação do sintoma - Depressão"
	anxiety_column    string = "Classificação do sintoma - Ansiedade"
	stress_column     string = "Classificação do sintoma - Estresse"
)

var header []string = []string{
	date_column,
	email_column,
	depression_column,
	anxiety_column,
	stress_column,
}

// Writes the scoring for each entry to `output` and returns any error that occurs during file creation
func WriteOutputCSV(output string, reports []reporting.Report) error {

	n := len(reports)
	records := make([][]string, n+1)

	records[0] = header

	for i, r := range reports {
		records[i+1] = []string{r.Date,
			r.Email,
			strconv.Itoa(r.Depression),
			strconv.Itoa(r.Anxiety),
			strconv.Itoa(r.Stress),
		}

	}

	file, err := os.Create(filepath.Clean(output))

	if err != nil {
		log.Fatalln("unable to create output file", err)
		return err
	}

	defer file.Close()

	w := csv.NewWriter(file)

	if err = w.WriteAll(records); err != nil {
		log.Fatalln(err)
		return err
	}

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv", err)
		return err
	}

	return nil
}
