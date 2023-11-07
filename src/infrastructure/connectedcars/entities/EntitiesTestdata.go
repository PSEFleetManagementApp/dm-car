package entities

// Collection of data for testing

var TestCarEntity = ConnectedCarsEntity{
	Vin:   "JH4DB1561NS000565",
	Brand: "VW",
	Model: "ID2",
}

var TestCarsEntity = []ConnectedCarsEntity{
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
