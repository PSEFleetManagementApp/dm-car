package model

type CarRepositoryInterface interface {
	GetCar(vin Vin) (Car, error)
}
