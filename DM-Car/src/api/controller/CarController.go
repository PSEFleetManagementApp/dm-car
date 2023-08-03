package controller

import (
	"car/DM-Car/src/api/stubs"
	"car/DM-Car/src/logic/operations"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CarController struct {
	ops operations.CarOperationsInterface
}

func NewCarController(ops operations.CarOperations) CarController {
	return CarController{ops: ops}
}

func (resource CarController) PostCar(ctx echo.Context) error {
	var payload stubs.PostCarJSONRequestBody

	err := ctx.Bind(&payload)
	if err != nil {
		return err
	}

	car, err := resource.ops.AddCar(payload.Vin.String(), *payload.Brand, *payload.Model)
	if err != nil {
		return err
	}

	err = ctx.JSON(http.StatusOK, car)
	if err != nil {
		return err
	}

	return nil
}

func (resource CarController) GetCar(ctx echo.Context) error {
	cars, err := resource.ops.GetCars()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, cars)
}

func (resource CarController) GetCarVin(ctx echo.Context, vin stubs.Vin) error {
	car, err := resource.ops.GetCar(vin.String())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, car)
}
