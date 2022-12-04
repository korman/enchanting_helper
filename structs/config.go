package structs

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Percentage float32 `json:"percentage"`
	CustodyFee float32 `json:"custody_fee"`
}

func NewConfig(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	c := new(Config)

	err = json.Unmarshal(f, c)

	if err != nil {
		log.Fatal(err)
	}

	return c, nil
}
