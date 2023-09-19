package operations

import (
	"car/logic/model"
	"errors"
)

func (ops CarOperations) GetCar(vin string) (model.Car, error) {
	vinObject := model.Vin{Vin: vin}

	// Check that the Vin is valid
	if !model.IsValidVin(vinObject) {
		return model.Car{}, errors.New("invalid Vin provided to GetCar")
	}

	car, err := ops.repository.GetCar(vinObject)
	if err != nil {
		return model.Car{}, err
	}

	return car, nil
}
