package infrastructure

import (
	"car/DM-Car/src/infrastructure/entities"
	"car/DM-Car/src/infrastructure/mappers"
	"car/DM-Car/src/logic/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// A common interface for the real and mocked database connection
type PGXInterface interface {
	Ping(ctx context.Context) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Close(ctx context.Context) error
}

type CarRepository struct {
	databaseConnection PGXInterface
}

func NewCarRepository() *CarRepository {
	connection, err := CreateDatabaseConnection()
	if err != nil {
		panic(err)
	}
	return &CarRepository{connection}
}

// Establish a connection to the database
func CreateDatabaseConnection() (PGXInterface, error) {
	url, err := getDatabaseURL()
	if err != nil {
		return nil, err
	}
	// Create the actual connection to the database
	connection, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	// Ping the database to confirm that the connection works
	err = connection.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return connection, nil
}

// The database url is of the format:
// postgres://USER:PASSWORD@HOST:PORT/DB_NAME
func getDatabaseURL() (string, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user,
		password,
		host,
		port,
		dbname), nil
}

func (repository *CarRepository) AddCar(car model.Car) error {
	carPersistenceEntity := mappers.ConvertCarToCarPersistenceEntity(car)
	statement := fmt.Sprintf(`
		INSERT INTO public."Car" (vin, brand, model)
		VALUES ('%s', '%s', '%s')
	`,
		carPersistenceEntity.Vin.Vin,
		carPersistenceEntity.Brand,
		carPersistenceEntity.Model)
	// Exec performs mutations on the database
	_, err := repository.databaseConnection.Exec(context.Background(), statement)
	return err
}

func (repository *CarRepository) GetCars() (model.Cars, error) {
	statement := `
		SELECT *
		FROM public."Car"
	`
	// Query can return multiple rows as a result
	rows, err := repository.databaseConnection.Query(context.Background(), statement)
	if err != nil {
		return model.Cars{}, err
	}
	// Rows are a resource that need to be closed so that they can be free'd from memory
	defer rows.Close()
	cars := []entities.CarPersistenceEntity{}
	for rows.Next() {
		car := entities.CarPersistenceEntity{}
		vin := entities.Vin{}
		err = rows.Scan(&vin.Vin, &car.Brand, &car.Model)
		if err != nil {
			return model.Cars{}, err
		}
		car.Vin = vin
		cars = append(cars, car)
	}
	err = rows.Err()
	if err != nil {
		return model.Cars{}, err
	}
	var result = mappers.ConvertCarsPersistenceEntityToCars(entities.CarsPersistenceEntity{
		Cars: cars,
	})
	return result, nil
}

func (repository *CarRepository) GetCar(vin string) (model.Car, error) {
	car := entities.CarPersistenceEntity{}
	vinObject := entities.Vin{}
	statement := fmt.Sprintf(`
		SELECT *
		FROM public."Car"
		WHERE vin LIKE '%s'
	`, vin)
	// QueryRow only returns a single row as a result
	// A single row does not need to be closed
	row := repository.databaseConnection.QueryRow(context.Background(), statement)
	switch err := row.Scan(&vinObject.Vin, &car.Model, &car.Brand); err {
	case sql.ErrNoRows:
		return model.Car{}, errors.New("no car has the specified VIN")
	case nil:
		car.Vin = vinObject
		return mappers.ConvertCarPersistenceEntityToCar(car), nil
	default:
		return model.Car{}, errors.New("DatabaseConnection.FindByVin: Unknown error")
	}
}

// The CarRepository is responsible for closing the database connection
func (repository *CarRepository) Close() error {
	return repository.databaseConnection.Close(context.Background())
}
