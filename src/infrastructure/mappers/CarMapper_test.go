package mappers

import (
	"car/infrastructure/persistenceentities"
	"car/logic/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCarToCarPersistenceEntity(t *testing.T) {
	carPersistenceEntity := ConvertCarToCarPersistenceEntity(model.TestCarModel)
	assert.Equal(t, model.TestCarModel.Vin.Vin, carPersistenceEntity.Vin)
	assert.Equal(t, model.TestCarModel.Brand, carPersistenceEntity.Brand)
	assert.Equal(t, model.TestCarModel.Model, carPersistenceEntity.Model)
}

func TestConvertCarPersistenceEntityToCar(t *testing.T) {
	carModel := ConvertCarPersistenceEntityToCar(persistenceentities.TestCarEntity)
	assert.Equal(t, persistenceentities.TestCarEntity.Vin, carModel.Vin.Vin)
	assert.Equal(t, persistenceentities.TestCarEntity.Brand, carModel.Brand)
	assert.Equal(t, persistenceentities.TestCarEntity.Model, carModel.Model)
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
	carsModel := ConvertCarPersistenceEntitiesToCars(persistenceentities.TestCarsEntity)
	for index, value := range persistenceentities.TestCarsEntity {
		assert.Equal(t, value.Vin, carsModel.Cars[index].Vin.Vin)
		assert.Equal(t, value.Brand, carsModel.Cars[index].Brand)
		assert.Equal(t, value.Model, carsModel.Cars[index].Model)
	}
}
