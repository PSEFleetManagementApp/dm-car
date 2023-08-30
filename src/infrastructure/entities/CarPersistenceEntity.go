package entities

// Car model for persistence
type CarPersistenceEntity struct {
	Vin   VinPersistenceEntity
	Brand string
	Model string
}
