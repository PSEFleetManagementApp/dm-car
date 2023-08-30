package operations

import (
	"car/src/infrastructure"
	"car/src/infrastructure/mappers"
	entities2 "car/src/infrastructure/persistenceentities"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of getting an existing car by it's Vin works
func TestGetCar(t *testing.T) {
	carRepository := infrastructure.MockCarRepository{MockDatabase: map[string]entities2.CarPersistenceEntity{
		"JH4DB1561NS000565": entities2.TestCarEntity,
	}}
	carOperations := NewCarOperations(&carRepository)

	carResult, err := carOperations.GetCar("JH4DB1561NS000565")
	assert.Equal(t, mappers.ConvertCarPersistenceEntityToCar(entities2.TestCarEntity), carResult)
	assert.Equal(t, err, nil)
}
