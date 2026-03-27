package main

import (
	"fmt"
)

func add(a, b int) int {

	return a + b
}
func sub(a, b int) int {

	return a - b
}

func mul(a, b int) int {

	return a * b
}

func div(a, b int) int {

	return a / b
}

func main() {

	// var number1 int
	// var number2 int

	var a int
	var b int

	var Op string
	for {
		fmt.Print("Select Opperations: \n add \n sub \n mul \n div \n Quit --> To Exit \n Select operation to continue... ")

		fmt.Scanln(&Op)

		// Operation line
		switch Op {
		case "add", "Add":
			fmt.Print("Enter first Number: ")
			fmt.Scanln(&a)
			fmt.Print("Enter second Number: ")
			fmt.Scanln(&b)
			fmt.Println("result: ", add(a, b))

		case "sub", "Sub":
			fmt.Println("result: ", sub(a, b))

		case "mul", "Mul":
			fmt.Println("result: ", mul(a, b))
		case "div", "Div":
			if b == 0 {
				fmt.Println(a, "is not divisible by ", b)
				continue
			}
			if a%b != 0 {
				fmt.Println("Input number: ", a, " is not Divisible by, ", b)
				continue
			} else {
				fmt.Println("return: ", div(a, b))
			}

		}

	}
}

// 	_, err := fmt.Scanln(&Op)

// 	if err != nil {
// 		fmt.Println("invalid input")
// 		return
// 	}
// 	if Op == "quit" || Op == "Quit" {
// 		fmt.Println("Goodbye!!!")
// 		break
// 	}
// 	//Input 1 collection
// 	fmt.Print("Enter first number: ")
// 	_, err = fmt.Scanln(&a)
// if err != nil {
// 	fmt.Println("Not A valid Number!")
// 	fmt.Print("Input only Digit, Example '5', '3' , '100', '49' etc \n")
// 	break
// }

// 	// Input 2 collection
// 	fmt.Print("Enter second number: ")
// 	_, err = fmt.Scanln(&b)
// if err != nil {
// 	fmt.Println("Not A valid Number")
// 	fmt.Println("Input only Digit, Example '5','3','100', '49' etc")
// }
