package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println(name, price)

	case "t":
		tip, _ := getInput("Enter tip amount ($)", reader)
		fmt.Println(tip)
		
	case "s":

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
