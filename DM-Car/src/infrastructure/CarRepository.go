package infrastructure

import (
	"car/DM-Car/src/infrastructure/mappers"
	"car/DM-Car/src/logic/model"
)

type CarRepository struct {
	connection DatabaseConnection
}

func NewCarRepository() *CarRepository {
	connection, err := CreateDatabaseConnection()
	if err != nil {
		panic(err)
	}
	return &CarRepository{connection}
}

func (repository *CarRepository) Save(car model.Car) error {
	carPersistenceEntity := mappers.ConvertCarToCarPersistenceEntity(car)
	return repository.connection.Save(carPersistenceEntity)
}

func (repository *CarRepository) FindAll() ([]model.Car, error) {
	cars, err := repository.connection.FindAll()
	if err != nil {
		return nil, err
	}
	var result []model.Car
	for i := range cars {
		result = append(result, mappers.ConvertCarPersistenceEntityToCar(cars[i]))
	}
	return result, nil
}

func (repository *CarRepository) FindByVin(vin string) (model.Car, error) {
	car, err := repository.connection.FindByVin(vin)
	if err != nil {
		return model.Car{}, err
	}
	return mappers.ConvertCarPersistenceEntityToCar(car), nil
}

func (repository *CarRepository) Close() error {
	return repository.connection.Close()
}
