package connectedcar

import (
	"car/infrastructure/connectedcar/entities"
	"car/infrastructure/connectedcar/mappers"
	"car/logic/model"
	"errors"
)

type ConnectedCarSystem struct {
	Cars map[string]entities.ConnectedCarEntity
}

func (repository *ConnectedCarSystem) GetCar(vin model.Vin) (model.Car, error) {
	car, ok := repository.Cars[vin.Vin]
	if ok {
		return mappers.ConvertConnectedCarEntityToCar(car), nil
	}
	return model.Car{}, errors.New("not found")
}
