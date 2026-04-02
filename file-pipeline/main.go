package main

import (
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
