package entities

// Collection of data for testing

var TestCarEntity = CarPersistenceEntity{
	Vin:   Vin{Vin: "JH4DB1561NS000565"},
	Brand: "VW",
	Model: "ID2",
}

var TestCarsEntity = CarsPersistenceEntity{
	Cars: []CarPersistenceEntity{
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
	},
}
