package infrastructure

import (
	"car/DM-Car/src/support"
	"testing"

	"github.com/pashagolub/pgxmock"
)

// Test that persisting a car works
// This assumes that only valid cars will be used as it is the controllers task
// to prevent invalid cars from entering the system
func TestAddCar(t *testing.T) {
	mockDatabaseConnection := support.CreateMockDatabaseConnection(t)

	mockDatabaseConnection.ExpectExec(`INSERT INTO public\."Car" \(vin, brand, model\) VALUES \('.*?', '.*?', '.*?'\)`).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	carRepository := CarRepository{connection: DatabaseConnection{connection: mockDatabaseConnection}}
	if err := carRepository.AddCar(support.Car); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	support.ExpectExpectationsToBeMet(mockDatabaseConnection, t)
}

// Test that retrieving all persisted cars works
// This uses the same assumptions as the test above
func TestGetCars(t *testing.T) {
	mockDatabaseConnection := support.CreateMockDatabaseConnection(t)

	rows := mockDatabaseConnection.NewRows([]string{"vin", "model", "brand"}).AddRow("JH4DA3350KS009715", "Mercedes-Benz", "S Klasse")
	mockDatabaseConnection.ExpectQuery(`SELECT \* FROM public\."Car"`).WillReturnRows(rows)

	carRepository := CarRepository{connection: DatabaseConnection{connection: mockDatabaseConnection}}
	if _, err := carRepository.GetCars(); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	support.ExpectExpectationsToBeMet(mockDatabaseConnection, t)
}

// Test that retrieving a specific persisted car works
// This uses the same assumptions as the test above
func TestGetCar(t *testing.T) {
	mockDatabaseConnection := support.CreateMockDatabaseConnection(t)

	rows := mockDatabaseConnection.NewRows([]string{"vin", "model", "brand"}).AddRow("JH4DA3350KS009715", "Mercedes-Benz", "S Klasse")
	mockDatabaseConnection.ExpectQuery(`SELECT \* FROM public\."Car" WHERE vin LIKE '.*?'`).WillReturnRows(rows)

	carRepository := CarRepository{connection: DatabaseConnection{connection: mockDatabaseConnection}}
	if _, err := carRepository.GetCar(support.Car.Vin.Vin); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	support.ExpectExpectationsToBeMet(mockDatabaseConnection, t)
}
