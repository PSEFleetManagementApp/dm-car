package operations

import (
	"car/infrastructure"
	"car/infrastructure/entities"
	"car/infrastructure/mappers"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of getting all existing cars works
func TestGetCars(t *testing.T) {
	carRepository := infrastructure.MockCarRepository{MockDatabase: map[string]entities.CarPersistenceEntity{
		"JH4DB1561NS000565": entities.TestCarEntity,
	}}
	carOperations := NewCarOperations(&carRepository)

	cars, err := carOperations.GetCars()
	assert.Contains(t, cars.Cars, mappers.ConvertCarPersistenceEntityToCar(entities.TestCarEntity))
	assert.Equal(t, err, nil)
}
