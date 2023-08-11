package mappers

import (
	"car/infrastructure/entities"
	"car/logic/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCarToCarPersistenceEntity(t *testing.T) {
	carPersistenceEntity := ConvertCarToCarPersistenceEntity(model.TestCarModel)
	assert.Equal(t, model.TestCarModel.Vin.Vin, carPersistenceEntity.Vin.Vin)
	assert.Equal(t, model.TestCarModel.Brand, carPersistenceEntity.Brand)
	assert.Equal(t, model.TestCarModel.Model, carPersistenceEntity.Model)
}

func TestConvertCarPersistenceEntityToCar(t *testing.T) {
	carModel := ConvertCarPersistenceEntityToCar(entities.TestCarEntity)
	assert.Equal(t, entities.TestCarEntity.Vin.Vin, carModel.Vin.Vin)
	assert.Equal(t, entities.TestCarEntity.Brand, carModel.Brand)
	assert.Equal(t, entities.TestCarEntity.Model, carModel.Model)
}

func TestConvertCarsToCarsPersistenceEntity(t *testing.T) {
	carsPersistenceEntity := ConvertCarsToCarsPersistenceEntity(model.TestCarsModel)
	for index, value := range model.TestCarsModel.Cars {
		assert.Equal(t, value.Vin.Vin, carsPersistenceEntity.Cars[index].Vin.Vin)
		assert.Equal(t, value.Brand, carsPersistenceEntity.Cars[index].Brand)
		assert.Equal(t, value.Model, carsPersistenceEntity.Cars[index].Model)
	}
}

func TestConvertCarsPersistenceEntityToCars(t *testing.T) {
	carsModel := ConvertCarsPersistenceEntityToCars(entities.TestCarsEntity)
	for index, value := range entities.TestCarsEntity.Cars {
		assert.Equal(t, value.Vin.Vin, carsModel.Cars[index].Vin.Vin)
		assert.Equal(t, value.Brand, carsModel.Cars[index].Brand)
		assert.Equal(t, value.Model, carsModel.Cars[index].Model)
	}
}
