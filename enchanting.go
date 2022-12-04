package main

import (
	"enchanting_helper/structs"
	"enchanting_helper/utils"
	"log"

	"github.com/modood/table"
)

func main() {

	prices, err := loadPrice()

	if err != nil {
		log.Fatal(err)
	}

	formulaMgr, err := utils.NewFormulaMgr("./data/formula.csv")

	if err != nil {
		log.Fatal(err)
	}

	config, err := structs.NewConfig("./data/config.json")

	if err != nil {
		log.Fatal(err)
	}

	showList, err := formulaMgr.ComputeResult(config, prices)

	if err != nil {
		log.Fatal(err)
	}

	table.Output(showList)
}

func loadPrice() (*structs.Price, error) {
	p, err := structs.NewPrice("./data/price.csv")

	if err != nil {
		return nil, err
	}

	return p, nil
}
