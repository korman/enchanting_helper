package tests

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"testing"
)

func TestCsvLoad(t *testing.T) {
	csvFile, err := os.Open("../data/price.csv")

	if err != nil {
		t.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var isFirst bool = false

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			t.Error(err)
		}

		if !isFirst {
			isFirst = true
			continue
		}

		println(line)
	}
}
