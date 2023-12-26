# Weather API CLI

This is a command-line tool that fetches and displays weather information from the [WeatherAPI](http://api.weatherapi.com).

## Prerequisites

Before running this tool, make sure you have the following prerequisites installed:

- Go (Golang): [Install Go](https://golang.org/doc/install)

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/Crvanetten15/weather-cli.git
   ```

Change the directory to the project folder:

```bash
cd weather-api-cli
```

2. Configure Your CLI

Create a configuration file named config.yaml with your WeatherAPI credentials:

```yaml
API_KEY: YOUR_API_KEY
LOCATION: YOUR_LOCATION
```

3. Build the program:

```bash
go build
```

4. Run the program:

```
weather-cli
```

## Usage

The program will fetch and display the current weather conditions for the specified location as well as the hourly forecast. It also colorizes the output based on the chance of rain.

## Example Output

```makefile
New York, NY : 63°F, Partly cloudy
2:00 PM - 64°F, 20%, Partly cloudy
3:00 PM - 64°F, 20%, Partly cloudy
4:00 PM - 63°F, 20%, Partly cloudy
5:00 PM - 63°F, 20%, Partly cloudy
6:00 PM - 62°F, 20%, Partly cloudy
7:00 PM - 61°F, 20%, Partly cloudy
8:00 PM - 60°F, 20%, Partly cloudy
9:00 PM - 59°F, 20%, Partly cloudy
10:00 PM - 59°F, 20%, Partly cloudy
11:00 PM - 58°F, 20%, Partly cloudy
12:00 AM - 58°F, 20%, Partly cloudy
```
