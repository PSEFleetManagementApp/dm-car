package entities

// Collection of data for testing

var TestCarEntity = ConnectedCarEntity{
	Vin:   "JH4DB1561NS000565",
	Brand: "VW",
	Model: "ID2",
}

var TestCarsEntity = []ConnectedCarEntity{
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
