package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type WeatherData struct {
	QueryCost       int     `json:"queryCost"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
	Address         string  `json:"address"`
	Timezone        string  `json:"timezone"`
	Tzoffset        float64 `json:"tzoffset"`
	Description     string  `json:"description"`
	Days            []struct {
		Datetime       string      `json:"datetime"`
		DatetimeEpoch  int         `json:"datetimeEpoch"`
		Tempmax        float64     `json:"tempmax"`
		Tempmin        float64     `json:"tempmin"`
		Temp           float64     `json:"temp"`
		Feelslikemax   float64     `json:"feelslikemax"`
		Feelslikemin   float64     `json:"feelslikemin"`
		Feelslike      float64     `json:"feelslike"`
		Dew            float64     `json:"dew"`
		Humidity       float64     `json:"humidity"`
		Precip         float64     `json:"precip"`
		Precipprob     float64     `json:"precipprob"`
		Precipcover    float64     `json:"precipcover"`
		Preciptype     interface{} `json:"preciptype"`
		Snow           float64     `json:"snow"`
		Snowdepth      float64     `json:"snowdepth"`
		Windgust       float64     `json:"windgust"`
		Windspeed      float64     `json:"windspeed"`
		Winddir        float64     `json:"winddir"`
		Pressure       float64     `json:"pressure"`
		Cloudcover     float64     `json:"cloudcover"`
		Visibility     float64     `json:"visibility"`
		Solarradiation float64     `json:"solarradiation"`
		Solarenergy    float64     `json:"solarenergy"`
		Uvindex        float64     `json:"uvindex"`
		Severerisk     float64     `json:"severerisk"`
		Sunrise        string      `json:"sunrise"`
		SunriseEpoch   int         `json:"sunriseEpoch"`
		Sunset         string      `json:"sunset"`
		SunsetEpoch    int         `json:"sunsetEpoch"`
		Moonphase      float64     `json:"moonphase"`
		Conditions     string      `json:"conditions"`
		Description    string      `json:"description"`
		Icon           string      `json:"icon"`
		Stations       []string    `json:"stations"`
		Source         string      `json:"source"`
		Hours          []struct {
			Datetime       string      `json:"datetime"`
			DatetimeEpoch  int         `json:"datetimeEpoch"`
			Temp           float64     `json:"temp"`
			Feelslike      float64     `json:"feelslike"`
			Humidity       float64     `json:"humidity"`
			Dew            float64     `json:"dew"`
			Precip         float64     `json:"precip"`
			Precipprob     float64     `json:"precipprob"`
			Snow           float64     `json:"snow"`
			Snowdepth      float64     `json:"snowdepth"`
			Preciptype     interface{} `json:"preciptype"`
			Windgust       float64     `json:"windgust"`
			Windspeed      float64     `json:"windspeed"`
			Winddir        float64     `json:"winddir"`
			Pressure       float64     `json:"pressure"`
			Visibility     float64     `json:"visibility"`
			Cloudcover     float64     `json:"cloudcover"`
			Solarradiation float64     `json:"solarradiation"`
			Solarenergy    float64     `json:"solarenergy"`
			Uvindex        float64     `json:"uvindex"`
			Severerisk     float64     `json:"severerisk"`
			Conditions     string      `json:"conditions"`
			Icon           string      `json:"icon"`
			Stations       []string    `json:"stations"`
			Source         string      `json:"source"`
		} `json:"hours"`
	} `json:"days"`
	Alerts []struct {
		Event       string `json:"event"`
		Headline    string `json:"headline"`
		Ends        string `json:"ends"`
		EndsEpoch   int    `json:"endsEpoch"`
		Onset       string `json:"onset"`
		OnsetEpoch  int    `json:"onsetEpoch"`
		ID          string `json:"id"`
		Language    string `json:"language"`
		Link        string `json:"link"`
		Description string `json:"description"`
	} `json:"alerts"`
	Stations struct {
		AU447 struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"AU447"`
		C6421 struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"C6421"`
		EDDB struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"EDDB"`
		E2835 struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"E2835"`
	} `json:"stations"`
	CurrentConditions struct {
		Datetime       string      `json:"datetime"`
		DatetimeEpoch  int         `json:"datetimeEpoch"`
		Temp           float64     `json:"temp"`
		Feelslike      float64     `json:"feelslike"`
		Humidity       float64     `json:"humidity"`
		Dew            float64     `json:"dew"`
		Precip         float64     `json:"precip"`
		Precipprob     float64     `json:"precipprob"`
		Snow           float64     `json:"snow"`
		Snowdepth      float64     `json:"snowdepth"`
		Preciptype     interface{} `json:"preciptype"`
		Windgust       float64     `json:"windgust"`
		Windspeed      float64     `json:"windspeed"`
		Winddir        float64     `json:"winddir"`
		Pressure       float64     `json:"pressure"`
		Visibility     float64     `json:"visibility"`
		Cloudcover     float64     `json:"cloudcover"`
		Solarradiation float64     `json:"solarradiation"`
		Solarenergy    float64     `json:"solarenergy"`
		Uvindex        float64     `json:"uvindex"`
		Conditions     string      `json:"conditions"`
		Icon           string      `json:"icon"`
		Stations       []string    `json:"stations"`
		Source         string      `json:"source"`
		Sunrise        string      `json:"sunrise"`
		SunriseEpoch   int         `json:"sunriseEpoch"`
		Sunset         string      `json:"sunset"`
		SunsetEpoch    int         `json:"sunsetEpoch"`
		Moonphase      float64     `json:"moonphase"`
	} `json:"currentConditions"`
}

func main() {

	const addition = ":00:00"

	var address string
	fmt.Print("Enter your address: ")
	address = readLine()
	address = url.QueryEscape(address)

	var locationun string
	fmt.Print("Enter your timezone (https://data.iana.org/time-zones/tzdb-2021a/zone1970.tab): ")
	locationun = readLine()

	var key string
	fmt.Print("Enter your key: ")
	key = readLine()
	fmt.Print("\n")

	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json", address, key)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while requesting!")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading response!")
		return
	}

	var weatherdata WeatherData
	json.Unmarshal([]byte(body), &weatherdata)

	location, err := time.LoadLocation(locationun)
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	if location != nil {
		fmt.Print("")
	}

	// Get the current time in the specified location
	currentTime := time.Now().In(location).Hour()
	prepared_time := strconv.Itoa(currentTime)
	if currentTime < 10 {
		prepared_time = "0" + prepared_time + addition
	} else {
		prepared_time = prepared_time + addition
	}
	for _, day := range weatherdata.Days {
		for _, hour := range day.Hours {
			realhour := hour.Datetime
			if prepared_time == realhour {
				fmt.Printf("Temp: %.1f\n", hour.Temp)
				break
			}
		}
		break
	}
	fmt.Printf("Address: %v\n", weatherdata.ResolvedAddress)

}

func readLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
