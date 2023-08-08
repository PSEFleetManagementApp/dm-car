package mappers

import (
	"car/DM-Car/src/infrastructure/entities"
	"car/DM-Car/src/support"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertCarToCarPersistenceEntity(t *testing.T) {
	carPersistenceEntity := ConvertCarToCarPersistenceEntity(support.Car)
	assert.Equal(t, support.Car.Vin.Vin, carPersistenceEntity.Vin.Vin)
	assert.Equal(t, support.Car.Brand, carPersistenceEntity.Brand)
	assert.Equal(t, support.Car.Model, carPersistenceEntity.Model)
}

func TestConvertCarPersistenceEntityToCar(t *testing.T) {
	carPersistenceEntity := entities.CarPersistenceEntity{
		Vin:   entities.Vin{Vin: "JH4DA3350KS009715"},
		Brand: "Mercedes-Benz",
		Model: "S Klasse",
	}
	carModel := ConvertCarPersistenceEntityToCar(carPersistenceEntity)
	assert.Equal(t, carPersistenceEntity.Vin.Vin, carModel.Vin.Vin)
	assert.Equal(t, carPersistenceEntity.Brand, carModel.Brand)
	assert.Equal(t, carPersistenceEntity.Model, carModel.Model)
}
