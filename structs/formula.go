package structs

type Formula struct {
	Materials map[string]int64
	Name      string
}

func NewFormula(data map[string]int64) (*Formula, error) {
	formula := new(Formula)
	formula.Materials = data

	return formula, nil
}
