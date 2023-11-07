package mappers

import (
	"car/infrastructure/connectedcars/entities"
	"car/logic/model"
)

func ConvertCarToConnectedCarsEntity(car model.Car) entities.ConnectedCarsEntity {
	return entities.ConnectedCarsEntity{
		Vin:   car.Vin.Vin,
		Brand: car.Brand,
		Model: car.Model,
	}
}

func ConvertConnectedCarsEntityToCar(carPersistenceEntity entities.ConnectedCarsEntity) model.Car {
	return model.Car{
		Vin:   model.Vin{Vin: carPersistenceEntity.Vin},
		Brand: carPersistenceEntity.Brand,
		Model: carPersistenceEntity.Model,
	}
}
