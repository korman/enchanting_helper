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

	header := make([]string, 0)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var isFirst bool = false

	for {
		data := make(map[string]int64)
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
	showList := make([]*structs.ShowTable, 0)

	for _, f := range m.Data {
		show := new(structs.ShowTable)
		show.Name = f.Name

		var total float64 = 0.0
		for k, v := range f.Materials {
			unitPrice, err := priceList.GetPrice(k)

			if err != nil {
				log.Fatal(err)
			}

			total += float64(v) * unitPrice
		}

		show.CustodyFee = N2S(float64(config.CustodyFee))
		show.AuctionFee = N2S(total * float64(config.Percentage))
		show.MaterialCost = N2S(total)
		show.Total = N2S(total + total*float64(config.Percentage) + float64(config.CustodyFee))

		showList = append(showList, show)
	}

	return showList, nil
}
