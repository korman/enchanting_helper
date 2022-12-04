package tests

import (
	"testing"

	"github.com/modood/table"
)

type OutputTable struct {
	Name  string `table:"附魔名"`
	Price string `table:"价格"`
}

func TestTableOutput(t *testing.T) {
	hs := []OutputTable{
		{"Stark", "direwolf"},
		{"Targaryen", "dragon"},
		{"Lannister", "lion"},
	}

	table.Output(hs)
}
