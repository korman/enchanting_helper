package utils

import (
	"bufio"
	"enchanting_helper/structs"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type FormulaMgr struct {
	Data []*structs.Formula
}

func NewFormulaMgr(path string) (*FormulaMgr, error) {
	f := new(FormulaMgr)
	dataList := make([]*structs.Formula, 0)

	csvFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]int64)
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

			continue
		}

		var name string = ""

		for k, v := range line {
			key := header[k]

			if key == "Name" {
				name = v
				continue
			}

			data[key], err = strconv.ParseInt(v, 10, 64)

			if err != nil {
				log.Panic(err)
			}
		}

		formula := new(structs.Formula)
		formula.Materials = data
		formula.Name = name

		dataList = append(dataList, formula)
	}

	f.Data = dataList

	return f, nil
}

func (m *FormulaMgr) ComputeResult(config *structs.Config, priceList *structs.Price) ([]*structs.ShowTable, error) {
	return nil, nil
}
