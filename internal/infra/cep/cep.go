package cep

type Place struct {
	City string
}

type PlaceProviderInterface interface {
	getByCep(cep string) (error, Place)
}
