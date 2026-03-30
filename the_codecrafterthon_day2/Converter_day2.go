package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Help(){

	fmt.Println("\033[42;3m -FOR SMOOTH EXPERIENCE USE THE FOLLOWING GUIDE LINE\033[0m")
	fmt.Println("convert 1E hex")
	fmt.Println("convert  1010 bin ")
	fmt.Println("convert  255 dec")
	fmt.Println("quit --> Exit")
	fmt.Println("Help --> For This Same Menu")
}
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\033[42;1m -WELCOME TO THE CONVERTER\033[0m")
	fmt.Println("Enter help for Guide line, Or Quit to exit the program")
	for {

		fmt.Print("\n> ")

		scanner.Scan()

		input := scanner.Text()

		input = strings.TrimSpace(strings.ToLower(input))

		if input == "quit" {
			fmt.Println("GOODBYE!!")
			break
		}
		if input == "help"{
			Help()
			continue
		}
		if len(input) == 0{
			fmt.Println("Empty input, Enter help for Guide line")
			continue
		}

		part := strings.Fields(input) 
		
		if len(part) > 3 {
			fmt.Println("To Many command, Enter Help for Guide line")
			continue
		}
		if len(part) != 3 {
			fmt.Println("Not Enough Input, Enter Help for Guide line")
			continue
		}
		

		command :=part[0]

		if command != "convert" {
			fmt.Println("Invalid Input, Accepted input E.g 'convert' or Enter help For Guide")
			continue
		}
		Str :=part[1]
		base := part[2]

		switch base {
		case "hex" :
			n, err := strconv.ParseInt(Str, 16, 64)
			if err != nil{
				fmt.Println("Invalid hex number")
				continue
			}
			fmt.Println("Decimal: ", n)
		case "bin" : 
			n, err := strconv.ParseInt(Str, 2 , 64)
			if err != nil {
				fmt.Println("Invalid binary number")
				continue
			}
			fmt.Println("Decimal: ", n)
		case "dec" :
			n, err := strconv.ParseInt(Str, 10, 64)
			if err != nil {
				fmt.Println("Invalid Decimal number")
				continue
			}

			fmt.Printf("Binary: %b\n", n)
			fmt.Printf("Hex: %X", n)
		default :
			fmt.Println("Invalid base input, enter help for guide line")
			continue
		}
	}
}
