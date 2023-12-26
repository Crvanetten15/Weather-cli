package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

// Struct to hold API Config Data
type apiConfig struct {
	Key      string `yaml:"API_KEY"`
	Location string `yaml:"LOCATION"`
}

// Method for return API Config Data and mapping it to our Struct
func setEnv() (apiConfig, error) {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("Error reading YAML file: %v", err)
		return apiConfig{}, fmt.Errorf("error reading YAML file: %v", err)
	}

	var config apiConfig

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Error unmarshaling YAML: %v", err)
		return apiConfig{}, fmt.Errorf("error reading YAML file: %v", err)
	}

	return config, nil
}

// Main Function
func main() {
	// Get Config Data
	config, err := setEnv()

	if err != nil {
		fmt.Printf("Error setting environment: %v", err)
		return
	}

	// Url formatting
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&aqi=no&alerts=no", config.Key, config.Location)

	// API Call
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		if res.StatusCode >= 400 && res.StatusCode <= 499 {
			err_mess := fmt.Sprintf("Status Code : %d\nGenerally a 400 code relates to errors in URL. Please Check API :\n%s", res.StatusCode, url)
			panic(err_mess)
		}

	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf(
		"%s, %s : %.0f F, %s\n",
		location.Name,
		location.Region,
		current.Tempf,
		current.Condition.Text)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}
		message := fmt.Sprintf(
			"%s - %.0f F, %.0f%%, %s\n",
			date.Format("3:04 PM"),
			hour.Tempf,
			hour.ChanceOfRain,
			hour.Condition.Text)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}

}

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		Region  string `json:"region"`
	} `json:"location"`
	Current struct {
		Tempf     float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				Tempf     float64 `json:"temp_f"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}
