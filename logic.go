package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API key not set")
	}

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
