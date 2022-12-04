package structs

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Price struct {
	Header []string
	Data   map[string]float64
}

func NewPrice(path string) (*Price, error) {
	csvFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	price := new(Price)
	data := make(map[string]float64)
	header := make([]string, 0)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var isFirst bool = false

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if !isFirst {
			isFirst = true

			for _, v := range line {
				header = append(header, v)
			}

			price.Header = header

			continue
		}

		for k, v := range line {
			key := price.Header[k]
			data[key], err = strconv.ParseFloat(v, 64)

			if err != nil {
				log.Panic(err)
			}
		}

		price.Data = data
	}

	return price, nil
}

func (m *Price) GetPrice(key string) (float64, error) {
	var err error = nil
	var price float64 = 0.0

	price = m.Data[key]

	return price, err
}
