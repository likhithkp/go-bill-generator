package main

type bill struct {
	name string
	bill map[string]float64
	tip  float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		bill: map[string]float64{},
		tip:  0,
	}

	return b
}
