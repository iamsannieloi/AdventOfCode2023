package main

import (
	"fmt"
	"strconv"
)

func isWordDigit(lookup string) bool {
	switch lookup {
	case
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine":
		return true
	}
	return false
}

func getFirstAndLastWordDigit(text string) int {
	startIndex := 0
	endIndex := len(text) - 1
	num1 := ""
	num2 := ""
	for num1 == "" || num2 == "" {
		if num1 == "" {
			startChar := string(text[startIndex])
			if isWordDigit(startChar) {
				num1 = startChar
			} else {
				startIndex += 1
			}
		}
		if num2 == "" {
			endChar := string(text[endIndex])
			if isDigit(endChar) {
				num2 = endChar
			} else {
				endIndex -= 1
			}
		}
	}
	result := num1 + num2
	i, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
