package controller

import (
	"car/DM-Car/src/api/dtos"
	"car/DM-Car/src/api/mappers"
	"car/DM-Car/src/logic/operations"
	"encoding/json"
	"net/http"
)

type CarCollectionResource struct {
	ops operations.CarOperationsInterface
}

func NewCarCollectionResource(ops operations.CarOperations) CarCollectionResource {
	return CarCollectionResource{ops: ops}
}

func (resource CarCollectionResource) HandleAddCar(writer http.ResponseWriter, request *http.Request) {
	payload := dtos.AddCarPayload{}

	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	car, err := resource.addCar(payload.Vin, payload.Brand, payload.Model)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(writer).Encode(car)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (resource CarCollectionResource) addCar(vin string, brand string, carModel string) (dtos.CarDto, error) {
	car, err := resource.ops.AddCar(vin, brand, carModel)
	if err != nil {
		return dtos.CarDto{}, err
	}

	return mappers.ConvertCarToCarDto(*car), nil
}

func (resource CarCollectionResource) HandleGetCar(writer http.ResponseWriter, request *http.Request) {
	vin := request.URL.Path[len("/car/"):]

	car, err := resource.getCar(vin)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(writer).Encode(car)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (resource CarCollectionResource) getCar(vin string) (dtos.CarDto, error) {
	car, err := resource.ops.GetCar(vin)
	if err != nil {
		return dtos.CarDto{}, err
	}
	return mappers.ConvertCarToCarDto(*car), nil
}
