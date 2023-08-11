package model

type CarRepositoryInterface interface {
	AddCar(car Car) error
	GetCars() (Cars, error)
	GetCar(vin string) (Car, error)
}
