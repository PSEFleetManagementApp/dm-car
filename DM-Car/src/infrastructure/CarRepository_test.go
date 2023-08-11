package infrastructure

import (
	"car/DM-Car/src/logic/model"
	"context"
	"testing"

	"github.com/pashagolub/pgxmock"
)

func CreateMockDatabaseConnection(t *testing.T) pgxmock.PgxConnIface {
	mockDatabaseConnection, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Contrary to the real database connection, the mocked database connection
	// is responsible for closing itself
	defer mockDatabaseConnection.Close(context.Background())
	return mockDatabaseConnection
}

// Helper function that validates that all expected SQL statements have been executed
func ExpectExpectationsToBeMet(mockDatabaseConnection pgxmock.PgxConnIface, t *testing.T) {
	if err := mockDatabaseConnection.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Test that persisting a car works
// This assumes that only valid cars will be used as it is the controllers task
// to prevent invalid cars from entering the system
func TestAddCar(t *testing.T) {
	mockDatabaseConnection := CreateMockDatabaseConnection(t)

	mockDatabaseConnection.ExpectExec(`INSERT INTO public\."Car" \(vin, brand, model\) VALUES \('.*?', '.*?', '.*?'\)`).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	carRepository := CarRepository{databaseConnection: mockDatabaseConnection}
	if err := carRepository.AddCar(model.TestCarModel); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	ExpectExpectationsToBeMet(mockDatabaseConnection, t)
}

// Test that retrieving all persisted cars works
// This uses the same assumptions as the test above
func TestGetCars(t *testing.T) {
	mockDatabaseConnection := CreateMockDatabaseConnection(t)

	rows := mockDatabaseConnection.NewRows([]string{"vin", "model", "brand"}).AddRow("JH4DA3350KS009715", "Mercedes-Benz", "S Klasse")
	mockDatabaseConnection.ExpectQuery(`SELECT \* FROM public\."Car"`).WillReturnRows(rows)

	carRepository := CarRepository{databaseConnection: mockDatabaseConnection}
	if _, err := carRepository.GetCars(); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	ExpectExpectationsToBeMet(mockDatabaseConnection, t)
}

// Test that retrieving a specific persisted car works
// This uses the same assumptions as the test above
func TestGetCar(t *testing.T) {
	mockDatabaseConnection := CreateMockDatabaseConnection(t)

	rows := mockDatabaseConnection.NewRows([]string{"vin", "model", "brand"}).AddRow("JH4DA3350KS009715", "Mercedes-Benz", "S Klasse")
	mockDatabaseConnection.ExpectQuery(`SELECT \* FROM public\."Car" WHERE vin LIKE '.*?'`).WillReturnRows(rows)

	carRepository := CarRepository{databaseConnection: mockDatabaseConnection}
	if _, err := carRepository.GetCar(model.TestCarModel.Vin.Vin); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	ExpectExpectationsToBeMet(mockDatabaseConnection, t)
}
