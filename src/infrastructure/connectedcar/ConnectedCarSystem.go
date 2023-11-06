package connectedcar

import (
	"car/infrastructure/connectedcar/entities"
	"car/infrastructure/connectedcar/mappers"
	"car/logic/model"
	_ "embed"
	"errors"
	"fmt"
	"github.com/gocarina/gocsv"
)

type ConnectedCarSystem struct {
	Cars []entities.ConnectedCarEntity
}

//go:embed cars.csv
var connectedCarCars []byte

func NewConnectedCarSystem() ConnectedCarSystem {
	fmt.Println(connectedCarCars)
	var cars []entities.ConnectedCarEntity
	if err := gocsv.UnmarshalBytes(connectedCarCars, &cars); err != nil {
		panic(err)
	}
	return ConnectedCarSystem{Cars: cars}
}

func (repository ConnectedCarSystem) GetCar(vin model.Vin) (model.Car, error) {
	for _, car := range repository.Cars {
		if car.Vin == vin.Vin {
			return mappers.ConvertConnectedCarEntityToCar(car), nil
		}
	}
	return model.Car{}, errors.New("not found")
}
