package reporting

import "fmt"

type Report struct {
	Date       string
	Email      string
	Depression int
	Anxiety    int
	Stress     int
}

const (
	date_column       string = "Carimbo de data/hora"
	email_column      string = "Endereço de e-mail"
	depression_column string = "Classificação do sintoma - Depressão"
	anxiety_column    string = "Classificação do sintoma - Ansiedade"
	stress_column     string = "Classificação do sintoma - Estresse"
)

func Print(r Report) {

	fmt.Printf("%s:\t\t\t\t%s\n", date_column, r.Date)
	fmt.Printf("%s:\t\t\t\t%s\n", email_column, r.Email)
	fmt.Printf("%s:\t\t\t%d\n", depression_column, r.Depression)
	fmt.Printf("%s:\t\t\t%d\n", anxiety_column, r.Anxiety)
	fmt.Printf("%s:\t\t\t%d\n\n", stress_column, r.Stress)
}
