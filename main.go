package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getUserInput(r *bufio.Reader) string {
	input, err := r.ReadString('\n')

	if err != nil {
		fmt.Print(err)
	}

	trimmedInput := strings.TrimSpace(strings.ToLower(input))

	return trimmedInput
}

func action() {
	reader := bufio.NewReader(os.Stdin)
	answer := getUserInput(reader)

	switch answer {
	case "y":
		begin()
	case "n":
		fmt.Println("Have a nice day!!")
	default:
		fmt.Println("Not a valid option")
		begin()
	}
}

func begin() {
	fmt.Print("Hey there, welcome to my simple weather app. What state in Nigeria would you like to know the current weather?: ")
	reader := bufio.NewReader(os.Stdin)
	answer := getUserInput(reader)
	fetchWeather(answer)
	fmt.Println(" ")
	fmt.Println("Would you like to know another state weather? Reply (Yes) with y or (No) with n")
	action()
}

func main() {
	begin()
}
