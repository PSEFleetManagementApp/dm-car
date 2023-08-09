package stubs

import "regexp"

// If any part of the Car is not present in a request body
// it will be set to nil instead of an empty string
type Car struct {
	Brand *string `json:"brand,omitempty"`
	Model *string `json:"model,omitempty"`
	Vin   *Vin    `json:"vin,omitempty"`
}

type Vin = string

type PostCarJSONRequestBody = Car

// Check that a Vin is valid according to the domain constraints
func IsValidVin(vin Vin) bool {
	match, err := regexp.MatchString("^[A-HJ-NPR-Z0-9]{13}[0-9]{4}$", vin)
	if err != nil {
		return false
	}
	return match
}
