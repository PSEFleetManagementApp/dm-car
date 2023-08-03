package main

import (
	"car/DM-Car/src/api/controller"
	"car/DM-Car/src/api/stubs"
	"car/DM-Car/src/infrastructure"
	"car/DM-Car/src/logic/operations"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	carOperations := operations.NewCarOperations(infrastructure.NewCarRepository())
	carsResource := controller.NewCarController(carOperations)

	e := echo.New()
	stubs.RegisterHandlers(e, &carsResource)

	var port = flag.Int("port", 8080, "Port for local server")
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
