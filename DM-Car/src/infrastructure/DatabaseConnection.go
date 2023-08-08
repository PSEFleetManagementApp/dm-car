package infrastructure

import (
	"car/DM-Car/src/infrastructure/entities"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type DatabaseConnectionInfo struct {
	host     string
	port     uint16
	user     string
	password string
	dbname   string
	sslmode  string
}

type DatabaseConnection struct {
	db *sql.DB
}

func CreateDatabaseConnection() (DatabaseConnection, error) {
	info, err := getDatabaseConnectionInfo()
	if err != nil {
		return DatabaseConnection{}, err
	}
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		info.host, info.port, info.user, info.password, info.dbname, info.sslmode)
	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return DatabaseConnection{}, err
	}
	err = db.Ping()
	if err != nil {
		return DatabaseConnection{}, err
	}
	return DatabaseConnection{db}, nil
}

func getDatabaseConnectionInfo() (DatabaseConnectionInfo, error) {
	host := os.Getenv("DB_HOST")
	portValue, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 64)
	if err != nil {
		return DatabaseConnectionInfo{}, err
	}
	port := uint16(portValue)
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	return DatabaseConnectionInfo{
		host,
		port,
		user,
		password,
		dbname,
		sslmode,
	}, nil
}

func (connection *DatabaseConnection) Close() error {
	return connection.db.Close()
}

func (connection *DatabaseConnection) Save(car entities.CarPersistenceEntity) error {
	statement := fmt.Sprintf(`
	INSERT INTO public."Car" (vin, brand, model)
	VALUES ('%s', '%s', '%s')
`, car.Vin, car.Brand, car.Model)
	_, err := connection.db.Exec(statement)
	return err
}

func (connection *DatabaseConnection) FindAll() ([]entities.CarPersistenceEntity, error) {
	statement := `
		SELECT *
		FROM public."Car"
	`
	rows, err := connection.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cars := []entities.CarPersistenceEntity{}
	for rows.Next() {
		car := entities.CarPersistenceEntity{}
		err = rows.Scan(&car.Vin, &car.Brand, &car.Model)
		if err != nil {
			return nil, err
		}
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
	statement := fmt.Sprintf(`
	SELECT *
	FROM public."Car"
	WHERE vin LIKE '%s'
`, vin)
	row := connection.db.QueryRow(statement)
	switch err := row.Scan(&car.Vin, &car.Model, &car.Brand); err {
	case sql.ErrNoRows:
		return entities.CarPersistenceEntity{}, errors.New("no car has the specified VIN")
	case nil:
		return car, nil
	default:
		return entities.CarPersistenceEntity{}, errors.New("DatabaseConnection.FindByVin: Unknown error")
	}
}
