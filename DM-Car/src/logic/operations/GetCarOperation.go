package operations

import "car/DM-Car/src/logic/model"

func (ops CarOperations) GetCar(vin string) (model.Car, error) {
	car, err := ops.repository.FindByVin(vin)
	if err != nil {
		return model.Car{}, err
	}

	return car, nil
}
