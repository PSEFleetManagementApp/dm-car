# 3.MicroserviceEngineering

## API Diagram DM-Car

[API Diagram](pages/api_diagram_dm_car.md)

## API Specification DM-Car

[API Specification](DM-Car/src/api/specification/api_specification_dm_car.yaml)

## Guideline Generate API Controller Stubs from API Specification

[Generate API Controller Stubs from API Specification](pages/guideline_generate_go_api_controller_stubs_from_api_specification.md)

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
DB_HOST // Address of the database
DB_PORT // Port of the database
DB_USER // User to access the database
DB_PASSWORD // Password of the database user
DB_NAME // Name of the database
DB_SSLMODE // SSL Mode can be one of the following: "require" (default), "verify-full", "verify-ca", and "disable"
```

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
