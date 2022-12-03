package tests

import (
	"enchanting_helper/structs"
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
