package infrastructure

import (
	"car/DM-Car/src/support"
	"testing"

	"github.com/pashagolub/pgxmock"
)

func TestSave(t *testing.T) {
	mockDatabaseConnection := support.CreateMockDatabaseConnection(t)

	mockDatabaseConnection.ExpectExec(`INSERT INTO public\."Car" \(vin, brand, model\) VALUES \('.*?', '.*?', '.*?'\)`).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	carRepository := CarRepository{connection: DatabaseConnection{connection: mockDatabaseConnection}}
	if err := carRepository.Save(support.Car); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	support.ExpectExpectationsToBeMet(mockDatabaseConnection, t)
}

func TestFindAll(t *testing.T) {
	mockDatabaseConnection := support.CreateMockDatabaseConnection(t)

	rows := mockDatabaseConnection.NewRows([]string{"vin", "model", "brand"}).AddRow("JH4DA3350KS009715", "Mercedes-Benz", "S Klasse")
	mockDatabaseConnection.ExpectQuery(`SELECT \* FROM public\."Car"`).WillReturnRows(rows)

	carRepository := CarRepository{connection: DatabaseConnection{connection: mockDatabaseConnection}}
	if _, err := carRepository.FindAll(); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	support.ExpectExpectationsToBeMet(mockDatabaseConnection, t)
}

func TestFindByVin(t *testing.T) {
	mockDatabaseConnection := support.CreateMockDatabaseConnection(t)

	rows := mockDatabaseConnection.NewRows([]string{"vin", "model", "brand"}).AddRow("JH4DA3350KS009715", "Mercedes-Benz", "S Klasse")
	mockDatabaseConnection.ExpectQuery(`SELECT \* FROM public\."Car" WHERE vin LIKE '.*?'`).WillReturnRows(rows)

	carRepository := CarRepository{connection: DatabaseConnection{connection: mockDatabaseConnection}}
	if _, err := carRepository.FindByVin(support.Car.Vin.Vin); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	support.ExpectExpectationsToBeMet(mockDatabaseConnection, t)
}
