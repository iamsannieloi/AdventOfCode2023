package main

import (
	"fmt"
	"strconv"
)

func isDigit(lookup string) bool {
	switch lookup {
	case
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"0":
		return true
	}
	return false
}

func getFirstAndLastDigit(text string) int {
	startIndex := 0
	endIndex := len(text) - 1
	num1 := ""
	num2 := ""
	for num1 == "" || num2 == "" {
		if num1 == "" {
			startChar := string(text[startIndex])
			if isDigit(startChar) {
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
