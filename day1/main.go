package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(intarr []int) int {
	total := 0
	for _, v := range intarr {
		total += v
	}
	return total
}

func main() {
	// open file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	a := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trebuchet := getFirstAndLastDigit(scanner.Text())
		a = append(a, trebuchet)
	}

	total := sum(a)
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
