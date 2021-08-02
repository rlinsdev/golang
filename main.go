package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	
	name, _:= getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Create the Bill: ", b.name)

	return b
}

func prompOptions (b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - Add Item, s - save bill, t - add tip) ",reader)
	
	switch opt {

	case "a":
		name, _ := getInput("Item Name: ", reader)
		price, _ := getInput("item Price", reader)
		
		p, err := strconv.ParseFloat(price, 64)
		if err!=nil {
			fmt.Println("The prince must be a number")
			prompOptions(b)
		}

		b.addItem(name, p)
		fmt.Println("Item Added - ", name, price)
		prompOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount ($)", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err!=nil {
			fmt.Println("The tip must be a number")
			prompOptions(b)
		}

		b.updateTip(t)
		fmt.Println("Tip added: ", tip)

	case "s":
		fmt.Println("You choose to save the bill", b)

	default:
		fmt.Println("That is not a valid option...")
		prompOptions(b)

	}

}

func main () {
	mybill := createBill()
	prompOptions(mybill)

	fmt.Println(mybill)
}
