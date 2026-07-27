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
	reader := bufio.NewReader(os.Stdin)
	version := "1.0"

	fmt.Println("====================================")
	fmt.Println(title)
	fmt.Println(welcomeMessage)
	fmt.Println("Version:", version)
	fmt.Println("====================================")
	fmt.Println("Enter the first number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("You entered:", input)
	input = strings.TrimSpace(input)
	number1, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number. Please enter a valid integer.")
		return
	}
	fmt.Println("First number accepted:", number1)
}
