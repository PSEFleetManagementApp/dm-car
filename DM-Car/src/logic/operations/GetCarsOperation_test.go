package operations

import (
	"car/DM-Car/src/infrastructure"
	"car/DM-Car/src/infrastructure/entities"
	"car/DM-Car/src/infrastructure/mappers"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of getting all existing cars works
func TestGetCars(t *testing.T) {
	carRepository := infrastructure.MockCarRepository{MockDatabase: map[string]entities.CarPersistenceEntity{
		"JH4DA3350KS009715": entities.TestCarEntity,
	}}
	carOperations := NewCarOperations(&carRepository)

	cars, err := carOperations.GetCars()
	assert.Contains(t, cars.Cars, mappers.ConvertCarPersistenceEntityToCar(entities.TestCarEntity))
	assert.Equal(t, err, nil)
}
