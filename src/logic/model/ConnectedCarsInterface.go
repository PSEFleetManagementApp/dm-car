package model

type ConnectedCarsInterface interface {
	GetCar(vin Vin) (Car, error)
}
