package operations

import "car/DM-Car/src/logic/model"

func (ops CarOperations) GetCars() (model.Cars, error) {
	cars, err := ops.repository.GetCars()
	if err != nil {
		return model.Cars{}, err
	}

	return cars, nil
}
