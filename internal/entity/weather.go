package entity

type Weather struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

func (w *Weather) CalculateKelvin() float64 {
	w.TempK = w.TempC + 273
	return w.TempK
}
