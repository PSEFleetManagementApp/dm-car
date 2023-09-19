package persistenceentities

// Collection of data for testing

var TestCarEntity = CarPersistenceEntity{
	Vin:   "JH4DB1561NS000565",
	Brand: "VW",
	Model: "ID2",
}

var TestCarsEntity = []CarPersistenceEntity{
	{
		Vin:   "JH4DB1561NS000565",
		Brand: "VW",
		Model: "ID2",
	},
	{
		Vin:   "JN8AZ2NC5B9300256",
		Brand: "VW",
		Model: "ID2",
	},
}
