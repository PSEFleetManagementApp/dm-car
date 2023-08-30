package operations

import (
	model2 "car/src/logic/model"
	"errors"
)

func (ops CarOperations) GetCar(vin string) (model2.Car, error) {
	vinObject := model2.Vin{Vin: vin}

	// Check that the Vin is valid
	if !model2.IsValidVin(vinObject) {
		return model2.Car{}, errors.New("invalid Vin provided to GetCar")
	}

	car, err := ops.repository.GetCar(vinObject)
	if err != nil {
		return model2.Car{}, err
	}

	return car, nil
}
