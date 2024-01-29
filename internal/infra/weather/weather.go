package weather

type Weather struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

func (w *Weather) SetTemperatures(tempC float64, tempF float64) {
	w.TempC = tempC
	w.TempF = tempF
	w.TempK = w.TempC + 273
}

type WeatherProviderInterface interface {
	getWeather(city string) (Weather, error)
}
