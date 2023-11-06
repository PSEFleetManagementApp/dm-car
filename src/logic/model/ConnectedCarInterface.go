package model

type ConnectedCarInterface interface {
	GetCar(vin Vin) (Car, error)
}
