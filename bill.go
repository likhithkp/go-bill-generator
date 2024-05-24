package main

import "fmt"

type bill struct {
	name string
	bill map[string]float64
	tip  float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		bill: map[string]float64{"pie": 2.99, "cheeseCake": 3.99},
		tip:  0,
	}

	return b
}

func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.bill {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total :", total)
	return fs
}
