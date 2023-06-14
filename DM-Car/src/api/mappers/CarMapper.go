package mappers

import (
	"car/DM-Car/src/api/dtos"
	"car/DM-Car/src/infrastructure/entities"
	"car/DM-Car/src/logic/model"
)

func ConvertCarToCarDto(car model.Car) dtos.CarDto {
	return dtos.CarDto{
		Vin:   car.Vin,
		Brand: car.Brand,
		Model: car.Model,
	}
}

func ConvertCarPersistenceEntityToCar(carPersistenceEntity entities.CarPersistenceEntity) model.Car {
	return model.Car{
		Vin:   carPersistenceEntity.Vin,
		Brand: carPersistenceEntity.Brand,
		Model: carPersistenceEntity.Model,
	}
}
