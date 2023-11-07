package entities

// Car model for persistence
type ConnectedCarsEntity struct {
	Vin   string `csv:"vin"`
	Brand string `csv:"brand"`
	Model string `csv:"model"`
}
