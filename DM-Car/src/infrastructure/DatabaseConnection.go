package infrastructure

import (
	"car/DM-Car/src/infrastructure/entities"
	"car/DM-Car/src/support"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

// The database connection is stored as a custom interface to enable testing
// the CarRepository
type DatabaseConnection struct {
	connection support.PGXInterface
}

// Establish a connection to the database
func CreateDatabaseConnection() (DatabaseConnection, error) {
	url, err := getDatabaseURL()
	if err != nil {
		return DatabaseConnection{}, err
	}
	// Create the actual connection to the database
	connection, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return DatabaseConnection{}, err
	}
	// Ping the database to confirm that the connection works
	err = connection.Ping(context.Background())
	if err != nil {
		return DatabaseConnection{}, err
	}
	return DatabaseConnection{connection}, nil
}

// The database url is of the format:
// postgres://USER:PASSWORD@HOST:PORT/DB_NAME
func getDatabaseURL() (string, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname), nil
}

// A database connection needs to be closed as to not waste resources
// This task is delegated to the CarRepository which holds the database connection
func (connection *DatabaseConnection) Close() error {
	return connection.connection.Close(context.Background())
}

func (connection *DatabaseConnection) AddCar(car entities.CarPersistenceEntity) error {
	statement := fmt.Sprintf(`
	INSERT INTO public."Car" (vin, brand, model)
	VALUES ('%s', '%s', '%s')
`, car.Vin.Vin, car.Brand, car.Model)
	// Exec performs mutations on the database
	_, err := connection.connection.Exec(context.Background(), statement)
	return err
}

func (connection *DatabaseConnection) GetCars() ([]entities.CarPersistenceEntity, error) {
	statement := `
		SELECT *
		FROM public."Car"
	`
	// Query can return multiple rows as a result
	rows, err := connection.connection.Query(context.Background(), statement)
	if err != nil {
		return nil, err
	}
	// Rows are a resource that need to be closed so that they can be free'd from memory
	defer rows.Close()
	cars := []entities.CarPersistenceEntity{}
	for rows.Next() {
		car := entities.CarPersistenceEntity{}
		vin := entities.Vin{}
		err = rows.Scan(&vin.Vin, &car.Brand, &car.Model)
		if err != nil {
			return nil, err
		}
		car.Vin = vin
		cars = append(cars, car)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (connection *DatabaseConnection) GetCar(vin string) (entities.CarPersistenceEntity, error) {
	car := entities.CarPersistenceEntity{}
	vinObject := entities.Vin{}
	statement := fmt.Sprintf(`
	SELECT *
	FROM public."Car"
	WHERE vin LIKE '%s'
`, vin)
	// QueryRow only returns a single row as a result
	// A single row does not need to be closed
	row := connection.connection.QueryRow(context.Background(), statement)
	switch err := row.Scan(&vinObject.Vin, &car.Model, &car.Brand); err {
	case sql.ErrNoRows:
		return entities.CarPersistenceEntity{}, errors.New("no car has the specified VIN")
	case nil:
		car.Vin = vinObject
		return car, nil
	default:
		return entities.CarPersistenceEntity{}, errors.New("DatabaseConnection.FindByVin: Unknown error")
	}
}
