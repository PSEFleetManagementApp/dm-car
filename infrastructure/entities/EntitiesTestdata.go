package entities

// Collection of data for testing

var TestCarEntity = CarPersistenceEntity{
	Vin:   VinPersistenceEntity{Vin: "JH4DB1561NS000565"},
	Brand: "VW",
	Model: "ID2",
}

var TestCarsEntity = CarsPersistenceEntity{
	Cars: []CarPersistenceEntity{
		{
			Vin:   VinPersistenceEntity{Vin: "JH4DB1561NS000565"},
			Brand: "VW",
			Model: "ID2",
		},
		{
			Vin:   VinPersistenceEntity{Vin: "JN8AZ2NC5B9300256"},
			Brand: "VW",
			Model: "ID2",
		},
	},
}
