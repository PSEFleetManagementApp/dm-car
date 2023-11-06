package operations

import (
	"car/logic/model"
)

type CarOperationsInterface interface {
	GetCar(vin string) (model.Car, error)
}
