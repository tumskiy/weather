package calculate

import "weather/types"

func GetTempFromMap(apiWeather types.ApiWeather) []float64 {
	temperatures := []float64{apiWeather.Main.Temp}
	return temperatures
}

func ConvertToCelsiusFromKelvin(tempArr []float64) []float64 {
	for i, _ := range tempArr {
		tempArr[i] = tempArr[i] - 273.1
	}
	return tempArr
}
