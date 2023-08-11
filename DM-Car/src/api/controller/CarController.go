package controller

import (
	"car/DM-Car/src/api/stubs"
	"car/DM-Car/src/logic/operations"
	"errors"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

type CarController struct {
	ops operations.CarOperationsInterface
}

func NewCarController(ops operations.CarOperations) CarController {
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

	// Check that the Vin is valid
	if !IsValidVin(*payload.Vin) {
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
	if !IsValidVin(vin) {
		return errors.New("invalid Vin")
	}

	car, err := resource.ops.GetCar(vin)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, car)
}

// Check that a Vin is valid according to the domain constraints
func IsValidVin(vin stubs.Vin) bool {
	match, err := regexp.MatchString("^[A-HJ-NPR-Z0-9]{13}[0-9]{4}$", vin)
	if err != nil {
		return false
	}
	return match
}