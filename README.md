## Depression, Anxiety, and Stress Scale (DASS-21)

This project provides an implementation for calculating the scoring
based on DASS-21

### Requirements

To run this software you will need to have `go` installed which is available at [go.dev/dl](https://go.dev/dl/)

### Running

Clone this repository and open a terminal window inside the cloned folder

Place the .csv file in the root of the cloned folder - same path as the `main.go` file - with the following name `DASS-21 (respostas) - Respostas ao formulário 1.csv`

Run the following command in the terminal:

```console
go run main.go
```

A csv file named `dass-21-scoring-results.csv` will be created and will contain the scoring results for each
entry present in the input csv file.
