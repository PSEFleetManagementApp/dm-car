package mappers

import (
	"car/infrastructure/connectedcar/entities"
	"car/logic/model"
)

func ConvertCarToConnectedCarEntity(car model.Car) entities.ConnectedCarEntity {
	return entities.ConnectedCarEntity{
		Vin:   car.Vin.Vin,
		Brand: car.Brand,
		Model: car.Model,
	}
}

func ConvertConnectedCarEntityToCar(carPersistenceEntity entities.ConnectedCarEntity) model.Car {
	return model.Car{
		Vin:   model.Vin{Vin: carPersistenceEntity.Vin},
		Brand: carPersistenceEntity.Brand,
		Model: carPersistenceEntity.Model,
	}
}
