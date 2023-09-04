package persistenceentities

// The model of a collection of cars that is used by all internal operations
// This corresponds to the API Diagram
type CarsPersistenceEntity struct {
	Cars []CarPersistenceEntity
}
