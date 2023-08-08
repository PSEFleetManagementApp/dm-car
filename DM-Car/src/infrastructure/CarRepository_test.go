package infrastructure

import (
	"car/DM-Car/src/logic/model"
	"context"
	"testing"

	"github.com/pashagolub/pgxmock"
)

func TestSave(t *testing.T) {
	car := model.Car{
		Vin:   model.Vin{Vin: "2af3d31e-15ef-11ee-be56-0242ac120005"},
		Brand: "Mercedes-Benz",
		Model: "S Klasse",
	}

	mockDatabaseConnection, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDatabaseConnection.Close(context.Background())

	mockDatabaseConnection.ExpectExec(`INSERT INTO public\."Car" \(vin, brand, model\) VALUES \('.*?', '.*?', '.*?'\)`).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	carRepository := CarRepository{connection: DatabaseConnection{connection: mockDatabaseConnection}}
	if err = carRepository.Save(car); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	if err := mockDatabaseConnection.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindAll(t *testing.T) {
	mockDatabaseConnection, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDatabaseConnection.Close(context.Background())

	rows := mockDatabaseConnection.NewRows([]string{"vin", "model", "brand"}).AddRow("2af3d31e-15ef-11ee-be56-0242ac120005", "Mercedes-Benz", "S Klasse")
	mockDatabaseConnection.ExpectQuery(`SELECT \* FROM public\."Car"`).WillReturnRows(rows)

	carRepository := CarRepository{connection: DatabaseConnection{connection: mockDatabaseConnection}}
	if _, err = carRepository.FindAll(); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	if err := mockDatabaseConnection.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindByVin(t *testing.T) {
	car := model.Car{
		Vin:   model.Vin{Vin: "2af3d31e-15ef-11ee-be56-0242ac120005"},
		Brand: "Mercedes-Benz",
		Model: "S Klasse",
	}

	mockDatabaseConnection, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDatabaseConnection.Close(context.Background())

	rows := mockDatabaseConnection.NewRows([]string{"vin", "model", "brand"}).AddRow("2af3d31e-15ef-11ee-be56-0242ac120005", "Mercedes-Benz", "S Klasse")
	mockDatabaseConnection.ExpectQuery(`SELECT \* FROM public\."Car" WHERE vin LIKE '.*?'`).WillReturnRows(rows)

	carRepository := CarRepository{connection: DatabaseConnection{connection: mockDatabaseConnection}}
	if _, err = carRepository.FindByVin(car.Vin.Vin); err != nil {
		t.Errorf("did not expect error: %s", err)
	}

	if err := mockDatabaseConnection.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
