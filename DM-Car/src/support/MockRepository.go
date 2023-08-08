package support

import (
	"car/DM-Car/src/logic/model"
	"errors"
)

type MockCarRepository struct {
	MockDatabase map[string]model.Car
}

func (mockRepository *MockCarRepository) Save(car model.Car) error {
	if _, ok := mockRepository.MockDatabase[car.Vin.Vin]; ok {
		return errors.New("vin already exists")
	}
	mockRepository.MockDatabase[car.Vin.Vin] = car
	return nil
}

func (mockRepository *MockCarRepository) FindAll() ([]model.Car, error) {
	cars := []model.Car{}
	for _, value := range mockRepository.MockDatabase {
		cars = append(cars, value)
	}
	return cars, nil
}

func (mockRepository *MockCarRepository) FindByVin(vin string) (model.Car, error) {
	car, ok := mockRepository.MockDatabase[vin]
	if ok {
		return car, nil
	}
	return model.Car{}, errors.New("not found")
}
