package mappers

import (
	entities2 "car/src/infrastructure/persistenceentities"
	model2 "car/src/logic/model"
)

func ConvertCarToCarPersistenceEntity(car model2.Car) entities2.CarPersistenceEntity {
	return entities2.CarPersistenceEntity{
		Vin:   entities2.VinPersistenceEntity{Vin: car.Vin.Vin},
		Brand: car.Brand,
		Model: car.Model,
	}
}

func ConvertCarPersistenceEntityToCar(carPersistenceEntity entities2.CarPersistenceEntity) model2.Car {
	return model2.Car{
		Vin:   model2.Vin{Vin: carPersistenceEntity.Vin.Vin},
		Brand: carPersistenceEntity.Brand,
		Model: carPersistenceEntity.Model,
	}
}

func ConvertCarsToCarsPersistenceEntity(cars model2.Cars) entities2.CarsPersistenceEntity {
	carsPersistenceEntity := entities2.CarsPersistenceEntity{
		Cars: []entities2.CarPersistenceEntity{},
	}
	for _, car := range cars.Cars {
		carsPersistenceEntity.Cars = append(
			carsPersistenceEntity.Cars,
			ConvertCarToCarPersistenceEntity(car))
	}
	return carsPersistenceEntity
}

func ConvertCarsPersistenceEntityToCars(carsPersistenceEntity entities2.CarsPersistenceEntity) model2.Cars {
	cars := model2.Cars{
		Cars: []model2.Car{},
	}
	for _, carPersistenceEntity := range carsPersistenceEntity.Cars {
		cars.Cars = append(
			cars.Cars,
			ConvertCarPersistenceEntityToCar(carPersistenceEntity))
	}
	return cars
}
