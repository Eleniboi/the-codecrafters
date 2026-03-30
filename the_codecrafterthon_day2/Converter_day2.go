package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "strconv"
	// "strings"
)
func isBinary(s string)bool{
	for _, ch := range s{
		if ch != '0' && ch != '1'{
			return false
		}
	}
	return true
}

func isHex(s string) bool{

	for _, ch := range s{
		if ch >= '0' || ch <= '9' && ch >= 'a' || ch <= 'f'{
			return true
		}
	}
	return false
}

func isDec(s string)bool{
	if s == ""{
		return false
	}
	for i, ch := range s{
		if i == 0 && ch == '-'{
			continue
		}
		if ch >= '0' && ch <= '9'{
			return true
		}
	}
	return false
}

func Help() {
	fmt.Println("Guide line For Smooth Conversion")
	fmt.Println("Use the following as a Guide...")
	fmt.Println("convert 1E hex to convert --> Decimal")
	fmt.Println("convert 10 bin to convert --> Decimal")
	fmt.Println("convert 255 dec to convert--> Binary and Hex")
	fmt.Println("help ---> for this menu")
	fmt.Println("quit ---> To Exit")
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("THE CLI BASE CONVERTER🪄")
	fmt.Println("Enter Help For Guide Or Quit To Exit")

	for {
		fmt.Print("\n>")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		input = strings.ToLower(input)
		
		if len(input) == 0{
			fmt.Println("Empty Input. Enter help for Guide")
			continue
		}
		if input == "quit"{
			fmt.Println("GOODBYE!!")
			break
		}
		if input =="help" {
			Help()
			continue
		}

		part :=strings.Fields(input)
		
		if len(part) == 0{
			fmt.Println("Invalid Input")
			continue
		}

		command :=part[0]

		if command != "convert"{
			fmt.Println("Invalid command. use 'convert' or 'help'")
			continue
		}
		if len(part) != 3{
			fmt.Println("Not enough input, usage: convert value base ")
			continue
		}
		value := part[1]
		base := part[2]

		if base != "hex" && base != "bin" && base != "dec"{
			fmt.Println("Invalid base. Accepted base, 'bin','hex','dec'")
			continue
		}

		switch base {
		case "hex" :
			if !isHex(value){
				fmt.Println("Invalid Hex input")
				continue
			}
		case "bin" :
			if !isBinary(value){
				fmt.Println("Invalid Bin Input")
				continue
			}
		case "dec" :
			if !isDec(value){
				fmt.Println("Invalid Dec number")
				continue
			}
		}

		var num int64 
		var err error 

		switch base {
		case "bin" :
			num, err = strconv.ParseInt(value, 2, 64)
			if err != nil {
				fmt.Println("Invalid Conversion")
				continue
			}
			fmt.Println("Result: ", num)
		}

	}
}
