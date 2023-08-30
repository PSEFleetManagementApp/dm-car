package infrastructure

import (
	"car/src/infrastructure/mappers"
	entities2 "car/src/infrastructure/persistenceentities"
	model2 "car/src/logic/model"
	"errors"
)

// A mocked version of CarRepository that uses a Map instead of persisting Cars to the database
type MockCarRepository struct {
	MockDatabase map[string]entities2.CarPersistenceEntity
}

func (mockRepository *MockCarRepository) AddCar(car model2.Car) error {
	carPersistenceEntity := mappers.ConvertCarToCarPersistenceEntity(car)
	if _, ok := mockRepository.MockDatabase[car.Vin.Vin]; ok {
		return errors.New("vin already exists")
	}
	mockRepository.MockDatabase[car.Vin.Vin] = carPersistenceEntity
	return nil
}

func (mockRepository *MockCarRepository) GetCars() (model2.Cars, error) {
	cars := entities2.CarsPersistenceEntity{
		Cars: []entities2.CarPersistenceEntity{},
	}
	for _, value := range mockRepository.MockDatabase {
		cars.Cars = append(cars.Cars, value)
	}
	return mappers.ConvertCarsPersistenceEntityToCars(cars), nil
}

func (mockRepository *MockCarRepository) GetCar(vin model2.Vin) (model2.Car, error) {
	car, ok := mockRepository.MockDatabase[vin.Vin]
	if ok {
		return mappers.ConvertCarPersistenceEntityToCar(car), nil
	}
	return model2.Car{}, errors.New("not found")
}
