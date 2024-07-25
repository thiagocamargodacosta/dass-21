package scoring

import (
	"github.com/thiagocamargodacosta/dass-21/form"
)

// Score stores the classification of each subscale of the DASS-21 scale
type Score struct {
	Depression int
	Anxiety    int
	Stress     int
}

// These constants store the response pattern for each multiple choice item in the form
const (
	Zero  = "0 - Não se aplicou de maneira alguma"
	One   = "1 - Aplicou-se em algum grau, ou por pouco de tempo"
	Two   = "2 - Aplicou-se em um grau considerável, ou por uma boa parte do tempo"
	Three = "3 - Aplicou-se muito, ou na maioria do tempo"
)

func DASS(f form.Form) Score {

	var s Score = Score{
		Depression: Depression(f),
		Anxiety:    Anxiety(f),
		Stress:     Stress(f),
	}

	return s
}

func Depression(f form.Form) int {

	// Start by retrieving each of the responses related to this subscale
	var responses []string = []string{f.Q3, f.Q5, f.Q10, f.Q13, f.Q16, f.Q17, f.Q21}
	var score int

	for _, response := range responses {
		score += ToScale(response)
	}

	return score * 2
}

func Anxiety(f form.Form) int {

	// Start by retrieving each of the responses related to this subscale
	var responses []string = []string{f.Q2, f.Q4, f.Q7, f.Q9, f.Q15, f.Q19, f.Q20}
	var score int

	for _, response := range responses {
		score += ToScale(response)
	}

	return score * 2
}

func Stress(f form.Form) int {

	// Start by retrieving each of the responses related to this subscale
	var responses []string = []string{f.Q1, f.Q6, f.Q8, f.Q11, f.Q12, f.Q14, f.Q18}
	var score int

	for _, response := range responses {
		score += ToScale(response)
	}

	return score * 2
}

func ToScale(response string) int {

	var score int

	switch response {
	case Zero:
		score = 0
	case One:
		score = 1
	case Two:
		score = 2
	case Three:
		score = 3
	}

	return score
}
