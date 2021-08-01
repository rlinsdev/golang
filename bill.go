package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// Format the bill
func (b bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	// list item
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// add tip
	fs+= fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// total
	fs+= fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	
	return fs
}

// Update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add Item to the bill
func (b *bill) addItem (name string, price float64) {
	b.items[name] = price
}