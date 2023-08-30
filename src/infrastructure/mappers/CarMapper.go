package mappers

import (
	"car/infrastructure/persistenceentities"
	"car/logic/model"
)

func ConvertCarToCarPersistenceEntity(car model.Car) persistenceentities.CarPersistenceEntity {
	return persistenceentities.CarPersistenceEntity{
		Vin:   persistenceentities.VinPersistenceEntity{Vin: car.Vin.Vin},
		Brand: car.Brand,
		Model: car.Model,
	}
}

func ConvertCarPersistenceEntityToCar(carPersistenceEntity persistenceentities.CarPersistenceEntity) model.Car {
	return model.Car{
		Vin:   model.Vin{Vin: carPersistenceEntity.Vin.Vin},
		Brand: carPersistenceEntity.Brand,
		Model: carPersistenceEntity.Model,
	}
}

func ConvertCarsToCarsPersistenceEntity(cars model.Cars) persistenceentities.CarsPersistenceEntity {
	carsPersistenceEntity := persistenceentities.CarsPersistenceEntity{
		Cars: []persistenceentities.CarPersistenceEntity{},
	}
	for _, car := range cars.Cars {
		carsPersistenceEntity.Cars = append(
			carsPersistenceEntity.Cars,
			ConvertCarToCarPersistenceEntity(car))
	}
	return carsPersistenceEntity
}

func ConvertCarsPersistenceEntityToCars(carsPersistenceEntity persistenceentities.CarsPersistenceEntity) model.Cars {
	cars := model.Cars{
		Cars: []model.Car{},
	}
	for _, carPersistenceEntity := range carsPersistenceEntity.Cars {
		cars.Cars = append(
			cars.Cars,
			ConvertCarPersistenceEntityToCar(carPersistenceEntity))
	}
	return cars
}
