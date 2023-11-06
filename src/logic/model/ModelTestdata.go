package model

// Collection of data for testing

// A valid car model for testing
var TestCarModel = Car{
	Vin:   Vin{Vin: "JH4DB1561NS000565"},
	Brand: "VW",
	Model: "ID2",
}

var TestCarsModel = []Car{
	{
		Vin:   Vin{Vin: "JH4DB1561NS000565"},
		Brand: "VW",
		Model: "ID2",
	},
	{
		Vin:   Vin{Vin: "JN8AZ2NC5B9300256"},
		Brand: "VW",
		Model: "ID2",
	},
}
