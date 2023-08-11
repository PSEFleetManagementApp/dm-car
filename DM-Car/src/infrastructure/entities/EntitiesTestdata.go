package entities

// Collection of data for testing

var TestCarEntity = CarPersistenceEntity{
	Vin:   Vin{Vin: "JH4DA3350KS009715"},
	Brand: "Mercedes-Benz",
	Model: "S Klasse",
}

var TestCarsEntity = CarsPersistenceEntity{
	Cars: []CarPersistenceEntity{
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
