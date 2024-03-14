package calculate

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"weather/types"
)

const ApiKey string = "1"

func GetWeather(lat, lon float64) types.ApiWeather {
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lonStr := strconv.FormatFloat(lon, 'f', -1, 64)
	buildResp := "https://api.openweathermap.org/data/2.5/weather?lat=" + latStr + "&lon=" + lonStr + "&appid=" + ApiKey
	fmt.Println(buildResp)
	response, err := http.Get(buildResp)
	if err != nil {
		log.Fatal("error response weather")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	var weatherResponse types.ApiWeather
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	return weatherResponse
}

func GetLocation(name string) (lat, lon float64, err error) {
	lower := strings.ToLower(name)
	buildResp := "https://api.openweathermap.org/geo/1.0/direct?q=" + lower + "&limit=5&&appid=" + ApiKey
	fmt.Println(buildResp)
	response, err := http.Get(buildResp)
	if err != nil {
		log.Fatal("error response city")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	var coordinates types.City
	if err := json.Unmarshal(body, &coordinates); err != nil {
		fmt.Println("Can not unmarshal JSON")
		return 0, 0, err
	}
	for _, rec := range coordinates {
		return rec.Lat, rec.Lon, nil
	}
	return
}
