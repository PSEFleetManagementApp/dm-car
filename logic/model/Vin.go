package model

import "regexp"

// The model of a Vin that is used by all internal operations
// This corresponds to the API Diagram
type Vin struct {
	Vin string
}

// Check that a Vin is valid according to the domain constraints
func IsValidVin(vin Vin) bool {
	match, err := regexp.MatchString("^[A-HJ-NPR-Z0-9]{13}[0-9]{4}$", vin.Vin)
	if err != nil {
		return false
	}
	return match
}


