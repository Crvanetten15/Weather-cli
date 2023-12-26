package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type apiConfig struct {
	Key      string `yaml:"API_KEY"`
	Location string `yaml:"LOCATION"`
}

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

func main() {
	config, err := setEnv()
	if err != nil {
		fmt.Printf("Error setting environment: %v", err)
		return
	}

	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&aqi=no&alerts=no", config.Key, config.Location)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
