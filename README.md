# DM-CarV2.0
The domain microservice in its Version V2.0 enables the access to basic information of a car which is identified by its VIN. The information is provided via an external system ConnectedCars. In Version 2.0, the external system is  simulated by a [CSV file](src/infrastructure/connectedcars/cars.csv).

## API Specification DM-CarV2.0

[API Diagram DM-CarV2.0](https://gitlab.kit.edu/kit/cm/teaching/carrentalapp/carrentalappv2/doccarrentalappv2/-/blob/main/pages/ad_dm-car_v2.0.md)

[API Specification DM-CarV2.0](/src/api/specification/openapi.yaml)


## Run DM-Car

1. Set the environment variable `PORT` which should contain the port on which DM-Car will be reachable
2. Execute main.go
3. Call GET localhost:PORT/cars/\<vin\> to retrieve created car. Example Vin <code>JH4DB1561NS000565</code>

## Run Tests

1. Execute `go test -v ./...`
