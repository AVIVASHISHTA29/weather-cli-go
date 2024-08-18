package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

const apiKey = "<YOUR_API_KEY>"

type WeatherResponse struct {
    Name string `json:"name"`
    Main struct {
        Temp float64 `json:"temp"`
    } `json:"main"`
    Weather []struct {
        Description string `json:"description"`
    } `json:"weather"`
}

func getWeather(city string) (WeatherResponse, error) {
    var weatherData WeatherResponse
    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)

    resp, err := http.Get(url)
    if err != nil {
        return weatherData, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return weatherData, fmt.Errorf("failed to get weather data: %s", resp.Status)
    }

    err = json.NewDecoder(resp.Body).Decode(&weatherData)
    if err != nil {
        return weatherData, err
    }

    return weatherData, nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: weather-cli <city>")
        os.Exit(1)
    }

    city := os.Args[1]
    weather, err := getWeather(city)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    fmt.Printf("Weather in %s: %0.2f°C, %s\n", weather.Name, weather.Main.Temp, weather.Weather[0].Description)
}
