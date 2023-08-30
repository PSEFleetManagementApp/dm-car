package operations

import (
	"car/src/infrastructure"
	entities2 "car/src/infrastructure/entities"
	"car/src/infrastructure/mappers"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of getting all existing cars works
func TestGetCars(t *testing.T) {
	carRepository := infrastructure.MockCarRepository{MockDatabase: map[string]entities2.CarPersistenceEntity{
		"JH4DB1561NS000565": entities2.TestCarEntity,
	}}
	carOperations := NewCarOperations(&carRepository)

	cars, err := carOperations.GetCars()
	assert.Contains(t, cars.Cars, mappers.ConvertCarPersistenceEntityToCar(entities2.TestCarEntity))
	assert.Equal(t, err, nil)
}
