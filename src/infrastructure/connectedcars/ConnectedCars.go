package connectedcars

import (
	"car/infrastructure/connectedcars/entities"
	"car/infrastructure/connectedcars/mappers"
	"car/logic/model"
	_ "embed"
	"errors"
	"github.com/gocarina/gocsv"
)

type ConnectedCars struct {
	Cars []entities.ConnectedCarsEntity
}

//go:embed ConnectedCars.csv
var connectedCarCars []byte

func NewConnectedCars() ConnectedCars {
	var cars []entities.ConnectedCarsEntity
	if err := gocsv.UnmarshalBytes(connectedCarCars, &cars); err != nil {
		panic(err)
	}
	return ConnectedCars{Cars: cars}
}

func (repository ConnectedCars) GetCar(vin model.Vin) (model.Car, error) {
	for _, car := range repository.Cars {
		if car.Vin == vin.Vin {
			return mappers.ConvertConnectedCarsEntityToCar(car), nil
		}
	}
	return model.Car{}, errors.New("not found")
}
