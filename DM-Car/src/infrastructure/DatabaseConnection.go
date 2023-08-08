package infrastructure

import (
	"car/DM-Car/src/infrastructure/entities"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type PGXInterface interface {
	Ping(ctx context.Context) error
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Close(ctx context.Context) error
}

type DatabaseConnection struct {
	connection PGXInterface
}

func CreateDatabaseConnection() (DatabaseConnection, error) {
	url, err := getDatabaseURL()
	if err != nil {
		return DatabaseConnection{}, err
	}
	connection, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return DatabaseConnection{}, err
	}
	err = connection.Ping(context.Background())
	if err != nil {
		return DatabaseConnection{}, err
	}
	return DatabaseConnection{connection}, nil
}

func getDatabaseURL() (string, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname), nil
}

func (connection *DatabaseConnection) Close() error {
	return connection.connection.Close(context.Background())
}

func (connection *DatabaseConnection) Save(car entities.CarPersistenceEntity) error {
	statement := fmt.Sprintf(`
	INSERT INTO public."Car" (vin, brand, model)
	VALUES ('%s', '%s', '%s')
`, car.Vin.Vin, car.Brand, car.Model)
	_, err := connection.connection.Exec(context.Background(), statement)
	return err
}

func (connection *DatabaseConnection) FindAll() ([]entities.CarPersistenceEntity, error) {
	statement := `
		SELECT *
		FROM public."Car"
	`
	rows, err := connection.connection.Query(context.Background(), statement)
	if err != nil {
		return nil, err
	}
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

func (connection *DatabaseConnection) FindByVin(vin string) (entities.CarPersistenceEntity, error) {
	car := entities.CarPersistenceEntity{}
	vinObject := entities.Vin{}
	statement := fmt.Sprintf(`
	SELECT *
	FROM public."Car"
	WHERE vin LIKE '%s'
`, vin)
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
