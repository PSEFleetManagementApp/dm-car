package operations

import (
	"car/DM-Car/src/logic/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockCarRepository struct {
	mockDatabase map[string]model.Car
}

func (mockRepository *MockCarRepository) Save(car model.Car) error {
	mockRepository.mockDatabase[car.Vin.Vin] = car
	return nil
}

func (mockRepository *MockCarRepository) FindAll() ([]model.Car, error) {
	cars := []model.Car{}
	for _, value := range mockRepository.mockDatabase {
		cars = append(cars, value)
	}
	return cars, nil
}

func (mockRepository *MockCarRepository) FindByVin(vin string) (model.Car, error) {
	car, ok := mockRepository.mockDatabase[vin]
	if ok {
		return car, nil
	}
	return model.Car{}, errors.New("Not found")
}

func TestAddCar(t *testing.T) {
	car := model.Car{
		Vin:   model.Vin{Vin: "2af3d31e-15ef-11ee-be56-0242ac120005"},
		Brand: "Mercedes-Benz",
		Model: "S Klasse",
	}

	carRepository := MockCarRepository{mockDatabase: map[string]model.Car{}}
	carOperations := NewCarOperations(&carRepository)

	carResult, err := carOperations.AddCar(car.Vin.Vin, car.Brand, car.Model)
	assert.Equal(t, car, carResult)
	assert.Equal(t, err, nil)
}