package operations

import (
	model2 "car/src/logic/model"
)

type CarOperationsInterface interface {
	AddCar(vin string, brand string, model string) (model2.Car, error)
	GetCars() (model2.Cars, error)
	GetCar(vin string) (model2.Car, error)
}
