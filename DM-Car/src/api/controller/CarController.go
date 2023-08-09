package controller

import (
	"car/DM-Car/src/api/stubs"
	"car/DM-Car/src/logic/operations"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CarController struct {
	ops operations.CarOperationsInterface
}

func NewCarController(ops operations.CarOperations) CarController {
	return CarController{ops: ops}
}

func (resource CarController) AddCar(ctx echo.Context) error {
	var payload stubs.PostCarJSONRequestBody

	err := ctx.Bind(&payload)
	if err != nil {
		return err
	}

	if payload.Vin == nil || payload.Brand == nil || payload.Model == nil {
		return errors.New("invalid payload")
	}

	if !stubs.IsValidVin(*payload.Vin) {
		return errors.New("invalid Vin")
	}

	car, err := resource.ops.AddCar(*payload.Vin, *payload.Brand, *payload.Model)
	if err != nil {
		return err
	}

	err = ctx.JSON(http.StatusOK, car)
	if err != nil {
		return err
	}

	return nil
}

func (resource CarController) GetCars(ctx echo.Context) error {
	cars, err := resource.ops.GetCars()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, cars)
}

func (resource CarController) GetCar(ctx echo.Context, vin stubs.Vin) error {
	if !stubs.IsValidVin(vin) {
		return errors.New("invalid Vin")
	}

	car, err := resource.ops.GetCar(vin)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, car)
}
