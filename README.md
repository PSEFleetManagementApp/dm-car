# 3.MicroserviceEngineering

## API Diagram DM-Car

[API Diagram](pages/api_diagram_dm_car.md)

## API Specification DM-Car

[API Specification](DM-Car/src/api/specification/api_specification_dm_car.yaml)

## Setting up the database and the connection to it

1. Run this inside of a Postgres database:
```sql
CREATE TABLE IF NOT EXISTS public."Car"
(
    vin character varying(255) COLLATE pg_catalog."default" NOT NULL,
    brand character varying(255) COLLATE pg_catalog."default" NOT NULL,
    model character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Car_pkey" PRIMARY KEY (vin)
)
```

2. Create the following environment variables for DM-Car:
```env
POSTGRES_HOST // Address of the database
POSTGRES_PORT // Port of the database
POSTGRES_USER // User to access the database
POSTGRES_PASSWORD // Password of the database user
POSTGRES_NAME // Name of the database
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
