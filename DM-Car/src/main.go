package main

import (
	"car/DM-Car/src/api/controller"
	"car/DM-Car/src/api/stubs"
	"car/DM-Car/src/infrastructure"
	"car/DM-Car/src/logic/operations"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
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
		e.Logger.Fatal("invalid port number")	
	}
	var port = flag.Int("port", portNumber, "Port for local server")
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
