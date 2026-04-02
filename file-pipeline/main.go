package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// this function is handling trailing and leading whitespaces
func trimSpace(s string) string {
	result := strings.Fields(s)
	return strings.Join(result, " ")
}

// this function replace TODO with ✦ ACTION
func ToDo(s string) string {

	result := strings.Fields(s)

	for r := range result {
		if result[r] == "TODO" && r == 0 {
			result[r] = strings.ReplaceAll(result[r], result[r], "✦ ACTION")
		}
	}
	return strings.Join(result, " ")
}

func CapToTitle(word string) string {
	if len(word) == 0 {
		return word
	}
	result := strings.Fields(word)
	for i := 0; i < len(result); i++ {
		if string(result[i]) == strings.ToUpper(string(result[i])) {
			result[i] = strings.ToUpper(string(result[i][0])) + strings.ToLower(result[i][1:])
		}
	}
	return strings.Join(result, " ")
}
func ClassToRed(word string) string {
	if word == "" {
		return word
	}
	result := strings.Fields(word)
	for i := 0; i < len(result); i++ {
		if result[i] == "CLASSIFIED:" && i == 0 {
			result[i] = "[REDACTED]:"
		}
	}
	return strings.Join(result, " ")
}
func LowerToUpper(s string) string {

	result := strings.Fields(s)

	for i := 0; i < len(result); i++ {
		if result[i] == strings.ToLower(result[i]) {
			result[i] = strings.ToUpper(result[i])
		}
	}
	return strings.Join(result, " ")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong Run command:..........")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	if input == output{
		fmt.Println("Command can not be the same")
		return
	}

	f, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer f.Close()

	o, er := os.Create(output)
	if er != nil {
		fmt.Println("Error creating output file: ", er)
	}
	defer o.Close()
	scanner := bufio.NewScanner(f)
	writer := bufio.NewWriter(o)

	for scanner.Scan() {
		line := scanner.Text()
		text := trimSpace(line)
		text = ToDo(text)
		text = LowerToUpper(text)
		text = CapToTitle(text)
		text = ClassToRed(text)

		_, err := writer.WriteString(text + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	writer.Flush()
	fmt.Println("File Proceesed")
}
