package tests

import (
	"enchanting_helper/structs"
	"enchanting_helper/utils"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestLoadJson(t *testing.T) {
	f, err := ioutil.ReadFile("../data/price.json")

	if err != nil {
		t.Fatal(err)
	}

	p := &structs.Price{}

	err = json.Unmarshal(f, p)

	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadFormula(t *testing.T) {
	_, err := utils.NewFormulaMgr("../data/formula.csv")

	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadConfig(t *testing.T) {
	c, err := structs.NewConfig("../data/config.json")

	if err != nil {
		t.Fatal(err)
	}

	println(c.CustodyFee)
}
