package mappers

import (
	"car/DM-Car/src/infrastructure/entities"
	"car/DM-Car/src/logic/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCarToCarPersistenceEntity(t *testing.T) {
	carModel := model.Car{
		Vin:   model.Vin{Vin: "2af3d31e-15ef-11ee-be56-0242ac120005"},
		Brand: "Mercedes-Benz",
		Model: "S Klasse",
	}
	carPersistenceEntity := ConvertCarToCarPersistenceEntity(carModel)
	assert.Equal(t, carModel.Vin.Vin, carPersistenceEntity.Vin.Vin)
	assert.Equal(t, carModel.Brand, carPersistenceEntity.Brand)
	assert.Equal(t, carModel.Model, carPersistenceEntity.Model)
}

func TestConvertCarPersistenceEntityToCar(t *testing.T) {
	carPersistenceEntity := entities.CarPersistenceEntity{
		Vin:   entities.Vin{Vin: "2af3d31e-15ef-11ee-be56-0242ac120005"},
		Brand: "Mercedes-Benz",
		Model: "S Klasse",
	}
	carModel := ConvertCarPersistenceEntityToCar(carPersistenceEntity)
	assert.Equal(t, carPersistenceEntity.Vin.Vin, carModel.Vin.Vin)
	assert.Equal(t, carPersistenceEntity.Brand, carModel.Brand)
	assert.Equal(t, carPersistenceEntity.Model, carModel.Model)
}
