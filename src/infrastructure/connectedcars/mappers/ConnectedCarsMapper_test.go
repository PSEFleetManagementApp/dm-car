package mappers

import (
	"car/infrastructure/connectedcars/entities"
	"car/logic/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCarToConnectedCarEntity(t *testing.T) {
	carPersistenceEntity := ConvertCarToConnectedCarsEntity(model.TestCarModel)
	assert.Equal(t, model.TestCarModel.Vin.Vin, carPersistenceEntity.Vin)
	assert.Equal(t, model.TestCarModel.Brand, carPersistenceEntity.Brand)
	assert.Equal(t, model.TestCarModel.Model, carPersistenceEntity.Model)
}

func TestConvertConnectedCarEntityToCar(t *testing.T) {
	carModel := ConvertConnectedCarsEntityToCar(entities.TestCarEntity)
	assert.Equal(t, entities.TestCarEntity.Vin, carModel.Vin.Vin)
	assert.Equal(t, entities.TestCarEntity.Brand, carModel.Brand)
	assert.Equal(t, entities.TestCarEntity.Model, carModel.Model)
}
