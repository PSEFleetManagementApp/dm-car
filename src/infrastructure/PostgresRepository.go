package infrastructure

import (
	"car/infrastructure/mappers"
	"car/infrastructure/persistenceentities"
	"car/logic/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	databaseConnection *gorm.DB
}

func NewPostgresRepository() *PostgresRepository {
	databaseConnection, err := createDatabaseConnection()
	if err != nil {
		panic(err)
	}

	databaseConnection.AutoMigrate(
		&persistenceentities.CarPersistenceEntity{},
	)

	return &PostgresRepository{databaseConnection}
}

// Establish a connection to the database
func createDatabaseConnection() (*gorm.DB, error) {
	dsn := getDatabaseConnectionString()
	databaseConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return databaseConnection, nil
}

// The DSN string for GORM is of the format:
// host=HOST user=USER password=PASSWORD dbname=DB_NAME port=PORT sslmode=disable
func getDatabaseConnectionString() string {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")

	gormDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port)
	return gormDSN
}

func (repository *PostgresRepository) AddCar(car model.Car) error {
	carPersistenceEntity := mappers.ConvertCarToCarPersistenceEntity(car)
	result := repository.databaseConnection.Create(&carPersistenceEntity)
	return result.Error
}

func (repository *PostgresRepository) GetCars() (model.Cars, error) {
	cars := []persistenceentities.CarPersistenceEntity{}
	result := repository.databaseConnection.Find(&cars)
	if result.Error != nil {
		return model.Cars{}, result.Error
	}
	return mappers.ConvertCarPersistenceEntitiesToCars(cars), nil
}

func (repository *PostgresRepository) GetCar(vin model.Vin) (model.Car, error) {
	car := persistenceentities.CarPersistenceEntity{}
	result := repository.databaseConnection.Where(&persistenceentities.CarPersistenceEntity{
		Vin: vin.Vin,
	}).First(&car)
	return mappers.ConvertCarPersistenceEntityToCar(car), result.Error
}
