package infrastructure

import (
	"car/infrastructure/mappers"
	"car/infrastructure/persistenceentities"
	"car/logic/model"
	"errors"
)

// A mocked version of PostgresRepository that uses a Map instead of persisting Cars to the database
type InMemoryRepository struct {
	Cars map[string]persistenceentities.CarPersistenceEntity
}

func (repository *InMemoryRepository) AddCar(car model.Car) error {
	carPersistenceEntity := mappers.ConvertCarToCarPersistenceEntity(car)
	if _, ok := repository.Cars[car.Vin.Vin]; ok {
		return errors.New("vin already exists")
	}
	repository.Cars[car.Vin.Vin] = carPersistenceEntity
	return nil
}

func (repository *InMemoryRepository) GetCars() (model.Cars, error) {
	cars := []persistenceentities.CarPersistenceEntity{}
	for _, value := range repository.Cars {
		cars = append(cars, value)
	}
	return mappers.ConvertCarsPersistenceEntityToCars(cars), nil
}

func (repository *InMemoryRepository) GetCar(vin model.Vin) (model.Car, error) {
	car, ok := repository.Cars[vin.Vin]
	if ok {
		return mappers.ConvertCarPersistenceEntityToCar(car), nil
	}
	return model.Car{}, errors.New("not found")
}
