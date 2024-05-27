package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	CurrentRes Current `json:"current"`
}

type Current struct {
	TempCelsius      float64   `json:"temp_c"`
	WeatherCondition Condition `json:"condition"`
}

type Condition struct {
	Weather string `json:"text"`
}

func fetchWeather(l string) {
	const apiKey = "ae6a8cca1fea4de7ada212642242605"
	location := l
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%v&q=%v", apiKey, location)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	temp := responseObject.CurrentRes.TempCelsius
	weather := strings.ToLower(responseObject.CurrentRes.WeatherCondition.Weather)

	fmt.Printf("The weather at %v is %v at %v degree celsius \n", location, weather, temp)

}

func getInput() string {
	fmt.Print("Hey there, welcome to my simple weather app. What state in Nigeria would you like to know the current weather?: ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print(err)
	}

	trimmedInput := strings.TrimSpace(strings.ToLower(input))

	return trimmedInput
}

func main() {

	fetchWeather(getInput())

	fmt.Println(" ")

	fmt.Println("Would you like to know another state weather? Reply (Yes) with y or (No) with n")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print(err)
	}

	answer := strings.TrimSpace(strings.ToLower(input))

	switch answer {
	case "y":
		fetchWeather(getInput())
	case "n":
		fmt.Println("Have a nice day!!")
	default:
		fmt.Println("Not a valid option")
		fetchWeather(getInput())
	}

}
