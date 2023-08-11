package operations

import (
	"car/DM-Car/src/infrastructure"
	"car/DM-Car/src/infrastructure/entities"
	"car/DM-Car/src/infrastructure/mappers"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of getting an existing car by it's Vin works
func TestGetCar(t *testing.T) {
	carRepository := infrastructure.MockCarRepository{MockDatabase: map[string]entities.CarPersistenceEntity{
		"JH4DA3350KS009715": entities.TestCarEntity,
	}}
	carOperations := NewCarOperations(&carRepository)

	carResult, err := carOperations.GetCar("JH4DA3350KS009715")
	assert.Equal(t, mappers.ConvertCarPersistenceEntityToCar(entities.TestCarEntity), carResult)
	assert.Equal(t, err, nil)
}
