package model

// Collection of data for testing

// A valid car model for testing
var TestCarModel = Car{
	Vin:   Vin{Vin: "JH4DA3350KS009715"},
	Brand: "Mercedes-Benz",
	Model: "S Klasse",
}

var TestCarsModel = Cars{
	Cars: []Car{
		{
			Vin:   Vin{Vin: "JH4DA3350KS009715"},
			Brand: "Mercedes-Benz",
			Model: "S Klasse",
		},
		{
			Vin:   Vin{Vin: "2C4GP44362R700796"},
			Brand: "Mercedes-Benz",
			Model: "E Klasse",
		},
	},
}
