package infrastructure

import (
	"car/infrastructure/mappers"
	"car/infrastructure/persistenceentities"
	"car/logic/model"
	"errors"
)

// A mocked version of CarRepository that uses a Map instead of persisting Cars to the database
type MockCarRepository struct {
	MockDatabase map[string]persistenceentities.CarPersistenceEntity
}

func (mockRepository *MockCarRepository) AddCar(car model.Car) error {
	carPersistenceEntity := mappers.ConvertCarToCarPersistenceEntity(car)
	if _, ok := mockRepository.MockDatabase[car.Vin.Vin]; ok {
		return errors.New("vin already exists")
	}
	mockRepository.MockDatabase[car.Vin.Vin] = carPersistenceEntity
	return nil
}

func (mockRepository *MockCarRepository) GetCars() (model.Cars, error) {
	cars := persistenceentities.CarsPersistenceEntity{
		Cars: []persistenceentities.CarPersistenceEntity{},
	}
	for _, value := range mockRepository.MockDatabase {
		cars.Cars = append(cars.Cars, value)
	}
	return mappers.ConvertCarsPersistenceEntityToCars(cars), nil
}

func (mockRepository *MockCarRepository) GetCar(vin model.Vin) (model.Car, error) {
	car, ok := mockRepository.MockDatabase[vin.Vin]
	if ok {
		return mappers.ConvertCarPersistenceEntityToCar(car), nil
	}
	return model.Car{}, errors.New("not found")
}
