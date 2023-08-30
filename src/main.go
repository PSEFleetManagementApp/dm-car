package main

import (
	"car/api/controller"
	"car/api/stubs"
	"car/infrastructure"
	"car/logic/operations"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"strconv"
)

func main() {
	// Create the CarRepository
	// This also establishes the connection to the database
	carRepository := infrastructure.NewCarRepository()
	// Close the connection to the database when carRepository goes out of scope
	// This happens when the program exists
	defer carRepository.Close()

	// Create the CarOperations and the CarController
	carOperations := operations.NewCarOperations(carRepository)
	carsResource := controller.NewCarController(carOperations)

	// Register the CarController with the server for handling it's routes
	e := echo.New()
	stubs.RegisterHandlers(e, &carsResource)

	// Start the server
	var portNumber, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		e.Logger.Fatal("The port number configuration is incorrect. Did you set the environment variable PORT?")
	}
	var port = flag.Int("port", portNumber, "Port for local server")
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
