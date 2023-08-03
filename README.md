# 3.MicroserviceEngineering

## API Specification DM-Car

[API Specification](DM-Car/src/api/specification/api_specification_dm_car.yaml)

## Guideline Generate API Controller Stubs from API Specification

[Generate API Controller Stubs from API Specification](pages/guideline_generate_go_api_controller_stubs_from_api_specification.md)

## Run DM-Car

1. Execute main.go
2. Call POST localhost:8080/car to create a new car. Example Payload:
```
{
   "vin": "2af3d31e-15ef-11ee-be56-0242ac120002",
   "brand": "Mercedes Benz",
   "model": "S Klasse"
}
```
3. Call GET localhost:8080/car/<vin> to retrieve created car. Example vin <code>2af3d31e-15ef-11ee-be56-0242ac120002</code>
4. Call GET localhost:8080/car to retrieve all created cars.
