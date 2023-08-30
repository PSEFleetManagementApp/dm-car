package persistenceentities

// Car model for persistence
type CarPersistenceEntity struct {
	Vin   VinPersistenceEntity
	Brand string
	Model string
}
