# DM-CarV1.1
The domain microservice DM-Car copes with the heterogeneity of car APIs from different manufacturers.

## API Specification DM-CarV1.1

[API Specification DM-CarV1.1](/src/api/specification/openapi.yaml)

## Setting up the database and the connection to it

Create the following environment variables for DM-Car:

```env
POSTGRES_HOST // Address of the database
POSTGRES_PORT // Port of the database
POSTGRES_USER // User to access the database
POSTGRES_PASSWORD // Password of the database user
POSTGRES_DATABASE // Name of the database
```

## Run DM-Car

1. Set the environment variable `PORT` which should contain the port on which DM-Car will be reachable
2. Execute main.go
3. Call POST localhost:PORT/cars to create a new car. Example Payload:
```
{
   "vin": "JH4DB1561NS000565",
   "brand": "VW",
   "model": "ID2"
}
```
4. Call GET localhost:PORT/cars/<vin> to retrieve created car. Example Vin <code>JH4DB1561NS000565</code>
5. Call GET localhost:PORT/cars to retrieve all created cars.

## Run Tests

1. Execute `go test -v ./...`
