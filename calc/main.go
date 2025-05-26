package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) !=4 {
		fmt.Println("Usage: go run cal.go <num1> <operator> <num2>")
		return
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	operator := os.Args[2]
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: please provide valid values")
		return
	}

	switch operator {
	case "+":
		fmt.Println("Result: ", num1+num2)
	case "-":
		fmt.Println("Result: ", num1-num2)
	case "*":
		fmt.Println("Result: ", num1*num2)
	case "/":
		if num2==0 {
			fmt.Println("Error: Division by Zero")
		} else {
			fmt.Println("Result: ", num1/num2)
		}
	default:
		fmt.Println("Unsuported opersator")
	}
}