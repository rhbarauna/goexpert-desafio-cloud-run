package weatherapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/weather"
)

type WeatherApiResponse struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdated string  `json:"last_updated"`
		TempC       float64 `json:"temp_c"`
		TempF       float64 `json:"temp_f"`
		IsDay       int     `json:"is_day"`
		Condition   struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

type WeatherApi struct {
	httpClient http.Client
	token      string
}

func NewWeatherAPI(tkn string) *WeatherApi {
	return &WeatherApi{
		httpClient: http.Client{},
		token:      tkn,
	}
}

func (w *WeatherApi) getWeather(city string) (weather.Weather, error) {
	req_str := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", w.token, city)
	req, err := http.NewRequest(http.MethodGet, req_str, nil)
	weather := weather.Weather{}

	if err != nil {
		log.Printf("Falha ao montar a requisição WeatherApi. %s\n", err.Error())
		return weather, err
	}

	resp, err := v.httpClient.Do(req)

	if err != nil {
		log.Printf("Falha ao executar a requisição WeatherApi. %s\n", err.Error())
		return weather, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Erro na resposta da API:", resp.Status)
		return weather, errors.New("Erro na resposta da API WeatherApi")
	}

	var weatherApiResp WeatherApiResponse

	err = json.NewDecoder(resp.Body).Decode(&weatherApiResp)
	if err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		return weather, err
	}

	weather.SetTemperatures(
		weatherApiResp.Current.TempC,
		weatherApiResp.Current.TempF,
	)

	return weather, nil
}
