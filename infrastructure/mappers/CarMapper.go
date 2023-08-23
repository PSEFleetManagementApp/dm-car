package mappers

import (
	"car/infrastructure/entities"
	"car/logic/model"
)

func ConvertCarToCarPersistenceEntity(car model.Car) entities.CarPersistenceEntity {
	return entities.CarPersistenceEntity{
		Vin:   entities.VinPersistenceEntity{Vin: car.Vin.Vin},
		Brand: car.Brand,
		Model: car.Model,
	}
}

func ConvertCarPersistenceEntityToCar(carPersistenceEntity entities.CarPersistenceEntity) model.Car {
	return model.Car{
		Vin:   model.Vin{Vin: carPersistenceEntity.Vin.Vin},
		Brand: carPersistenceEntity.Brand,
		Model: carPersistenceEntity.Model,
	}
}

func ConvertCarsToCarsPersistenceEntity(cars model.Cars) entities.CarsPersistenceEntity {
	carsPersistenceEntity := entities.CarsPersistenceEntity{
		Cars: []entities.CarPersistenceEntity{},
	}
	for _, car := range cars.Cars {
		carsPersistenceEntity.Cars = append(
			carsPersistenceEntity.Cars,
			ConvertCarToCarPersistenceEntity(car))
	}
	return carsPersistenceEntity
}

func ConvertCarsPersistenceEntityToCars(carsPersistenceEntity entities.CarsPersistenceEntity) model.Cars {
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
