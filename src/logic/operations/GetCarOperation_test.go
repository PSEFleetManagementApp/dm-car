package operations

import (
	"car/infrastructure/connectedcar"
	persistenceentities2 "car/infrastructure/connectedcar/entities"
	"car/infrastructure/connectedcar/mappers"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that the operation of getting an existing car by it's Vin works
func TestGetCar(t *testing.T) {
	carRepository := connectedcar.ConnectedCarSystem{Cars: map[string]persistenceentities2.ConnectedCarEntity{
		"JH4DB1561NS000565": persistenceentities2.TestCarEntity,
	}}
	carOperations := NewCarOperations(&carRepository)

	carResult, err := carOperations.GetCar("JH4DB1561NS000565")
	assert.Equal(t, mappers.ConvertConnectedCarEntityToCar(persistenceentities2.TestCarEntity), carResult)
	assert.Equal(t, err, nil)
}
