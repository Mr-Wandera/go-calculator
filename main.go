package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	title := "GO CALCULATOR"
	welcomeMessage := "Welcome to my first Go project!"
	version := "1.0"
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("====================================")
	fmt.Println(title)
	fmt.Println(welcomeMessage)
	fmt.Println("Version:", version)
	fmt.Println("====================================")
	fmt.Print("Enter the first number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(input)
	number1, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number. Please enter a valid integer.")
		return
	}
	fmt.Println("First number accepted:", number1)
	fmt.Print("Select an operator (+, -, *, /): ")
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	operator := strings.TrimSpace(input)
	fmt.Println("Operator accepted:", operator)
	fmt.Print("Enter the second number: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(input)
	number2, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number. Please enter a valid integer.")
		return
	}
	fmt.Println("Second number accepted:", number2)
	switch operator {
	case "+":
		fmt.Println("Result:", number1+number2)
	case "-":
		fmt.Println("Result:", number1-number2)
	case "*":
		fmt.Println("Result:", number1*number2)
	case "/":
		if number2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		} else {
			fmt.Println("Result:", number1/number2)
		}
	default:
		fmt.Println("Invalid operator. Please select a valid operator (+, -, *, /).")
	}
}
