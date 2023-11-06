package controller

import (
	"car/logic/operations"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CarController struct {
	ops operations.CarOperationsInterface
}

func NewCarController(ops operations.CarOperations) CarController {
	return CarController{ops: ops}
}

// Route: GET /cars/:vin
func (resource CarController) GetCar(ctx echo.Context, vin Vin) error {
	car, err := resource.ops.GetCar(vin)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, car)
}
