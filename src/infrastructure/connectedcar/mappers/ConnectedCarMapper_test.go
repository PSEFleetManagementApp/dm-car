package mappers

import (
	"car/infrastructure/connectedcar/entities"
	"car/logic/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCarToCarPersistenceEntity(t *testing.T) {
	carPersistenceEntity := ConvertCarToConnectedCarEntity(model.TestCarModel)
	assert.Equal(t, model.TestCarModel.Vin.Vin, carPersistenceEntity.Vin)
	assert.Equal(t, model.TestCarModel.Brand, carPersistenceEntity.Brand)
	assert.Equal(t, model.TestCarModel.Model, carPersistenceEntity.Model)
}

func TestConvertCarPersistenceEntityToCar(t *testing.T) {
	carModel := ConvertConnectedCarEntityToCar(entities.TestCarEntity)
	assert.Equal(t, entities.TestCarEntity.Vin, carModel.Vin.Vin)
	assert.Equal(t, entities.TestCarEntity.Brand, carModel.Brand)
	assert.Equal(t, entities.TestCarEntity.Model, carModel.Model)
}

func TestConvertCarsToCarPersistenceEntities(t *testing.T) {
	carsPersistenceEntity := ConvertCarsToCarPersistenceEntities(model.TestCarsModel)
	for index, value := range model.TestCarsModel.Cars {
		assert.Equal(t, value.Vin.Vin, carsPersistenceEntity[index].Vin)
		assert.Equal(t, value.Brand, carsPersistenceEntity[index].Brand)
		assert.Equal(t, value.Model, carsPersistenceEntity[index].Model)
	}
}

func TestConvertCarPersistenceEntitiesToCars(t *testing.T) {
	carsModel := ConvertCarPersistenceEntitiesToCars(entities.TestCarsEntity)
	for index, value := range entities.TestCarsEntity {
		assert.Equal(t, value.Vin, carsModel.Cars[index].Vin.Vin)
		assert.Equal(t, value.Brand, carsModel.Cars[index].Brand)
		assert.Equal(t, value.Model, carsModel.Cars[index].Model)
	}
}
