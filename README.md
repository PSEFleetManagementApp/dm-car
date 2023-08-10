# 3.MicroserviceEngineering

## API Diagram DM-Car

[API Diagram](pages/api_diagram_dm_car.md)

## API Specification DM-Car

[API Specification](DM-Car/src/api/specification/api_specification_dm_car.yaml)

## Run DM-Car

1. Execute main.go
2. Call POST localhost:8080/cars to create a new car. Example Payload:
```
{
   "vin": "2af3d31e-15ef-11ee-be56-0242ac120002",
   "brand": "Mercedes Benz",
   "model": "S Klasse"
}
```
3. Call GET localhost:8080/cars/<vin> to retrieve created car. Example vin <code>2af3d31e-15ef-11ee-be56-0242ac120002</code>
4. Call GET localhost:8080/cars to retrieve all created cars.
