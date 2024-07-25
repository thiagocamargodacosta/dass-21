package form

// Form defines a structure that comprises of all the information present in the dass-21 form
type Form struct {
	Date  string
	Email string
	Q1    string
	Q2    string
	Q3    string
	Q4    string
	Q5    string
	Q6    string
	Q7    string
	Q8    string
	Q9    string
	Q10   string
	Q11   string
	Q12   string
	Q13   string
	Q14   string
	Q15   string
	Q16   string
	Q17   string
	Q18   string
	Q19   string
	Q20   string
	Q21   string
}

// Turns a row into a Form
func CreateForm(data []string) Form {

	var f Form

	for i, field := range data {
		switch i {
		case 0:
			f.Date = field
		case 1:
			f.Email = field
		case 2:
			f.Q1 = field
		case 3:
			f.Q2 = field
		case 4:
			f.Q3 = field
		case 5:
			f.Q4 = field
		case 6:
			f.Q5 = field
		case 7:
			f.Q6 = field
		case 8:
			f.Q7 = field
		case 9:
			f.Q8 = field
		case 10:
			f.Q9 = field
		case 11:
			f.Q10 = field
		case 12:
			f.Q11 = field
		case 13:
			f.Q12 = field
		case 14:
			f.Q13 = field
		case 15:
			f.Q14 = field
		case 16:
			f.Q15 = field
		case 17:
			f.Q16 = field
		case 18:
			f.Q17 = field
		case 19:
			f.Q18 = field
		case 20:
			f.Q19 = field
		case 21:
			f.Q20 = field
		case 22:
			f.Q21 = field
		}
	}
	return f
}

// Turns a collection of rows into a collection of Form
func CreateForms(data [][]string) []Form {

	var forms []Form

	for i, line := range data {
		if i > 0 {
			forms = append(forms, CreateForm(line))
		}
	}

	return forms
}
