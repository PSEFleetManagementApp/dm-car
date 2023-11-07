package operations

import (
	"car/infrastructure/connectedcars"
	persistenceentities2 "car/infrastructure/connectedcars/entities"
	"car/infrastructure/connectedcars/mappers"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of getting an existing car by it's Vin works
func TestGetCar(t *testing.T) {
	carRepository := connectedcars.ConnectedCars{Cars: []persistenceentities2.ConnectedCarsEntity{
		0: persistenceentities2.TestCarEntity,
	}}
	carOperations := NewCarOperations(&carRepository)

	carResult, err := carOperations.GetCar("JH4DB1561NS000565")
	assert.Equal(t, mappers.ConvertConnectedCarsEntityToCar(persistenceentities2.TestCarEntity), carResult)
	assert.Equal(t, err, nil)
}
