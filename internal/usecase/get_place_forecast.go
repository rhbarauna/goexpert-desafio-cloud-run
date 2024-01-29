package usecase

import (
	"errors"

	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/entity"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/place"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/weather"
)

var (
	ErrPostalCodeNotFound = errors.New("cannot found zipcode")
	ErrWeatherNotFound    = errors.New("weather not found for that zipcode")
	ErrInvalidInput       = errors.New("invalid postalcode")
)

type PlaceForecastOutputDTO struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

type GetPlaceForecast struct {
	placeProvider   place.PlaceProviderInterface
	weatherProvider weather.WeatherProviderInterface
}

func NewGetPlaceForecastUseCase(placeProvider place.PlaceProviderInterface, weatherProvider weather.WeatherProviderInterface) GetPlaceForecast {
	return GetPlaceForecast{
		placeProvider:   placeProvider,
		weatherProvider: weatherProvider,
	}
}

func (uc *GetPlaceForecast) Execute(cep string) (PlaceForecastOutputDTO, error) {
	normalized, err := entity.NormalizePostalCode(cep)
	outputDTO := PlaceForecastOutputDTO{}

	if err != nil {
		return outputDTO, ErrInvalidInput
	}

	placeDetails, err := uc.placeProvider.GetByCep(normalized)

	if err != nil || placeDetails.IsValid() != nil {
		return outputDTO, ErrPostalCodeNotFound
	}

	forecast, err := uc.weatherProvider.GetWeather(placeDetails.City)

	if err != nil {
		return outputDTO, ErrWeatherNotFound
	}

	outputDTO.TempC = forecast.TempC
	outputDTO.TempF = forecast.CalculateFahrenheit()
	outputDTO.TempK = forecast.CalculateKelvin()

	return outputDTO, nil
}
