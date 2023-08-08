package mappers

import (
	"car/DM-Car/src/infrastructure/entities"
	"car/DM-Car/src/logic/model"
)

func ConvertCarToCarPersistenceEntity(car model.Car) entities.CarPersistenceEntity {
	return entities.CarPersistenceEntity{
		Vin:   entities.Vin{Vin: car.Vin.Vin},
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
