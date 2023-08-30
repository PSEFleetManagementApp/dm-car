package operations

import (
	"car/infrastructure"
	"car/infrastructure/mappers"
	"car/infrastructure/persistenceentities"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of getting all existing cars works
func TestGetCars(t *testing.T) {
	carRepository := infrastructure.InMemoryRepository{Cars: map[string]persistenceentities.CarPersistenceEntity{
		"JH4DB1561NS000565": persistenceentities.TestCarEntity,
	}}
	carOperations := NewCarOperations(&carRepository)

	cars, err := carOperations.GetCars()
	assert.Contains(t, cars.Cars, mappers.ConvertCarPersistenceEntityToCar(persistenceentities.TestCarEntity))
	assert.Equal(t, err, nil)
}
