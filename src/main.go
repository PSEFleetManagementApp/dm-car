package main

import (
	"car/api/controller"
	"car/infrastructure/connectedcars"
	"car/logic/operations"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"strconv"
)

func main() {
	// Create the PostgresRepository
	// This also establishes the connection to the database
	connectedCars := connectedcars.NewConnectedCars()

	// Create the CarOperations and the CarController
	carOperations := operations.NewCarOperations(connectedCars)
	carsResource := controller.NewCarController(carOperations)

	// Register the CarController with the server for handling it's routes
	e := echo.New()
	controller.RegisterHandlers(e, &carsResource)

	// Start the server
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		portEnv = "80"
	}

	var portNumber, err = strconv.Atoi(portEnv)
	if err != nil {
		e.Logger.Fatal("The port number configuration is incorrect. Did you set the environment variable PORT?")
	}
	var port = flag.Int("port", portNumber, "Port for local server")
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
