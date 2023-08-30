package operations

import (
	"car/infrastructure"
	"car/infrastructure/mappers"
	"car/infrastructure/persistenceentities"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of getting an existing car by it's Vin works
func TestGetCar(t *testing.T) {
	carRepository := infrastructure.MockCarRepository{MockDatabase: map[string]persistenceentities.CarPersistenceEntity{
		"JH4DB1561NS000565": persistenceentities.TestCarEntity,
	}}
	carOperations := NewCarOperations(&carRepository)

	carResult, err := carOperations.GetCar("JH4DB1561NS000565")
	assert.Equal(t, mappers.ConvertCarPersistenceEntityToCar(persistenceentities.TestCarEntity), carResult)
	assert.Equal(t, err, nil)
}
