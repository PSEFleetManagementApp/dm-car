package controller

import (
	"car/src/api/stubs"
	operations2 "car/src/logic/operations"
	"errors"
	"net/http"
)

type CarController struct {
	ops operations2.CarOperationsInterface
}

func NewCarController(ops operations2.CarOperations) CarController {
	return CarController{ops: ops}
}

// Route: POST /cars
func (resource CarController) AddCar(ctx echo.Context) error {
	var payload stubs.AddCarJSONRequestBody

	err := ctx.Bind(&payload)
	if err != nil {
		return err
	}

	// Check that all fields have been set
	if payload.Vin == nil || payload.Brand == nil || payload.Model == nil {
		return errors.New("invalid payload")
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

// Route: GET /cars
func (resource CarController) GetCars(ctx echo.Context) error {
	cars, err := resource.ops.GetCars()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, cars)
}

// Route: GET /cars/:vin
func (resource CarController) GetCar(ctx echo.Context, vin stubs.Vin) error {
	car, err := resource.ops.GetCar(vin)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, car)
}
