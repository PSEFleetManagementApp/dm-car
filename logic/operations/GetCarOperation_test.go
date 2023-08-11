package operations

import (
	"car/infrastructure"
	"car/infrastructure/entities"
	"car/infrastructure/mappers"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of getting an existing car by it's Vin works
func TestGetCar(t *testing.T) {
	carRepository := infrastructure.MockCarRepository{MockDatabase: map[string]entities.CarPersistenceEntity{
		"JH4DB1561NS000565": entities.TestCarEntity,
	}}
	carOperations := NewCarOperations(&carRepository)

	carResult, err := carOperations.GetCar("JH4DB1561NS000565")
	assert.Equal(t, mappers.ConvertCarPersistenceEntityToCar(entities.TestCarEntity), carResult)
	assert.Equal(t, err, nil)
}
