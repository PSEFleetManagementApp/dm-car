package persistenceentities

// Car model for persistence
type CarPersistenceEntity struct {
	Vin   string `gorm:"primaryKey"`
	Brand string
	Model string
}
