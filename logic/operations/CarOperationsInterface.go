package operations

import (
	"car/logic/model"
)

type CarOperationsInterface interface {
	AddCar(vin string, brand string, model string) (model.Car, error)
	GetCars() (model.Cars, error)
	GetCar(vin string) (model.Car, error)
}
