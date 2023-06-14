package operations

import (
	"car/DM-Car/src/logic/model"
)

func (ops CarOperations) AddCar(vin string, brand string, carModel string) (*model.Car, error) {
	car := model.Car{
		Vin:   vin,
		Brand: brand,
		Model: carModel,
	}

	err := ops.repository.Save(car)
	if err != nil {
		return &model.Car{}, err
	}

	return &car, nil
}
