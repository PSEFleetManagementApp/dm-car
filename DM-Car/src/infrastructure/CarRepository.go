package infrastructure

import (
	"car/DM-Car/src/infrastructure/entities"
	"car/DM-Car/src/infrastructure/mappers"
	"car/DM-Car/src/logic/model"
	"errors"
)

type CarRepository struct {
	Cars []entities.CarPersistenceEntity
}

func NewCarRepository() *CarRepository {
	return &CarRepository{Cars: []entities.CarPersistenceEntity{}}
}

func (repository *CarRepository) Save(car model.Car) error {
	for i, carInRepository := range repository.Cars {
		if carInRepository.Vin == car.Vin {
			repository.Cars[i] = mappers.ConvertCarToCarPersistenceEntity(car)
			return nil
		}
	}

	repository.Cars = append(repository.Cars, mappers.ConvertCarToCarPersistenceEntity(car))
	return nil
}

func (repository *CarRepository) FindAll() ([]model.Car, error) {
	var result []model.Car
	for i := range repository.Cars {
		result = append(result, *mappers.ConvertCarPersistenceEntityToCar(repository.Cars[i]))
	}
	return result, nil
}

func (repository *CarRepository) FindByVin(vin string) (*model.Car, error) {
	for i := range repository.Cars {
		if repository.Cars[i].Vin == vin {
			return mappers.ConvertCarPersistenceEntityToCar(repository.Cars[i]), nil
		}
	}

	return &model.Car{}, errors.New("there is no car with the given vin")
}
