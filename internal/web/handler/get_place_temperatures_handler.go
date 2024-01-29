package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/usecase"
)

type GetPlaceTemperaturesHandler struct {
	uc usecase.GetPlaceForecast
}

func NewGetPlaceTemperaturesHandler(uc usecase.GetPlaceForecast) GetPlaceTemperaturesHandler {
	return GetPlaceTemperaturesHandler{
		uc: uc,
	}
}

func (h *GetPlaceTemperaturesHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	output, err := h.uc.Execute(cep)

	if err != nil {
		if err == usecase.ErrPostalCodeNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err == usecase.ErrPostalCodeNotFound || err == usecase.ErrWeatherNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err == usecase.ErrPostalCodeNotFound || err == usecase.ErrWeatherNotFound {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//montar o DTO de resposta
	/*
	   ​​​Em caso de falha, caso o CEP não seja encontrado:
	   Código HTTP: 404
	   Mensagem: can not found zipcode
	   Deverá ser realizado o deploy no Google Cloud Run.

	*/
}
