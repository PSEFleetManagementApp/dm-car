package operations

import "car/DM-Car/src/logic/model"

func (ops CarOperations) GetCars() ([]model.Car, error) {
	cars, err := ops.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return cars, nil
}
