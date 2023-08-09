package model

type CarRepositoryInterface interface {
	AddCar(car Car) error
	GetCars() ([]Car, error)
	GetCar(vin string) (Car, error)
}
