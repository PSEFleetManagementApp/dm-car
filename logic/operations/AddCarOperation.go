package operations

import (
	"car/logic/model"
	"errors"
)

func (ops CarOperations) AddCar(vin string, brand string, carModel string) (model.Car, error) {
	car := model.Car{
		Vin:   model.Vin{Vin: vin},
		Brand: brand,
		Model: carModel,
	}

	// Check that the Vin is valid
	if !model.IsValidVin(car.Vin) {
		return model.Car{}, errors.New("invalid Vin provided to AddCar")
	}

	err := ops.repository.AddCar(car)
	if err != nil {
		return model.Car{}, err
	}

	return car, nil
}
