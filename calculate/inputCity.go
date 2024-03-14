package calculate

import (
	"fmt"
	"math"
	"strconv"
)

func InputCity() {
	var cityName string
	fmt.Print("Enter city name: ")
	fmt.Scan(&cityName)
	lat, lon, err := GetLocation(cityName)
	if err != nil {
		fmt.Println("ERROR")
	}
	t1 := ConvertToCelsiusFromKelvin(GetTempFromMap(GetWeather(lat, lon)))
	t := sumFloat64Slice(t1)
	t = math.Round(t*100) / 100
	tempStr := strconv.FormatFloat(t, 'f', -1, 64)
	fmt.Println(`Температура сейчас = ` + tempStr + "°")
}

func sumFloat64Slice(arr []float64) float64 {
	var sum float64
	for _, v := range arr {
		sum += v
	}
	return sum
}
