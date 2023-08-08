package operations

import (
	"car/DM-Car/src/logic/model"
)

type CarOperationsInterface interface {
	AddCar(vin string, brand string, model string) (model.Car, error)
	GetCars() ([]model.Car, error)
	GetCar(vin string) (model.Car, error)
}
