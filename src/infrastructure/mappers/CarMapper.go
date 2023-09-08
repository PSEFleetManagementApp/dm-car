package mappers

import (
	"car/infrastructure/persistenceentities"
	"car/logic/model"
)

func ConvertCarToCarPersistenceEntity(car model.Car) persistenceentities.CarPersistenceEntity {
	return persistenceentities.CarPersistenceEntity{
		Vin:   car.Vin.Vin,
		Brand: car.Brand,
		Model: car.Model,
	}
}

func ConvertCarPersistenceEntityToCar(carPersistenceEntity persistenceentities.CarPersistenceEntity) model.Car {
	return model.Car{
		Vin:   model.Vin{Vin: carPersistenceEntity.Vin},
		Brand: carPersistenceEntity.Brand,
		Model: carPersistenceEntity.Model,
	}
}

func ConvertCarsToCarPersistenceEntities(cars model.Cars) []persistenceentities.CarPersistenceEntity {
	carsPersistenceEntity := []persistenceentities.CarPersistenceEntity{}
	for _, car := range cars.Cars {
		carsPersistenceEntity = append(
			carsPersistenceEntity,
			ConvertCarToCarPersistenceEntity(car))
	}
	return carsPersistenceEntity
}

func ConvertCarPersistenceEntitiesToCars(carsPersistenceEntity []persistenceentities.CarPersistenceEntity) model.Cars {
	cars := model.Cars{
		Cars: []model.Car{},
	}
	for _, carPersistenceEntity := range carsPersistenceEntity {
		cars.Cars = append(
			cars.Cars,
			ConvertCarPersistenceEntityToCar(carPersistenceEntity))
	}
	return cars
}
