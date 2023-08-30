package operations

import (
	model2 "car/src/logic/model"
	"errors"
)

func (ops CarOperations) AddCar(vin string, brand string, carModel string) (model2.Car, error) {
	car := model2.Car{
		Vin:   model2.Vin{Vin: vin},
		Brand: brand,
		Model: carModel,
	}

	// Check that the Vin is valid
	if !model2.IsValidVin(car.Vin) {
		return model2.Car{}, errors.New("invalid Vin provided to AddCar")
	}

	err := ops.repository.AddCar(car)
	if err != nil {
		return model2.Car{}, err
	}

	return car, nil
}
