package support

import (
	"car/DM-Car/src/logic/model"
	"fmt"
)

// Collection of data for testing

// A valid car model for testing
var Car = model.Car{
	Vin:   model.Vin{Vin: "JH4DA3350KS009715"},
	Brand: "Mercedes-Benz",
	Model: "S Klasse",
}

// A valid request body for creating a car
var Body = fmt.Sprintf(`
	{
		"vin": "%s",
		"brand": "%s",
		"model": "%s"
	}
	`, Car.Vin.Vin, Car.Brand, Car.Model)

// List of valid Vins according to the domain constraints
var ValidVins = []string{
	"JH4DA3350KS009715",
	"2C4GP44362R700796",
	"1C3CDZBG8DN504146",
	"1GCDC14K2LE198114",
	"1G3NF52E3XC403652",
}

// List of invalid Vins according to the domain constraints
var InvalidVins = []string{
	"JH4DA3350KS00",
	"2CIGP44362R700796",
	"1C3CDZBG8DN5O4146",
	"1gCDC14K2LE198114",
	"1G3NF52E3XC4036521",
}
