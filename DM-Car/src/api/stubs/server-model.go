package stubs

import "regexp"

type Car struct {
	Brand *string `json:"brand,omitempty"`
	Model *string `json:"model,omitempty"`
	Vin   *Vin    `json:"vin,omitempty"`
}

type Vin = string

type PostCarJSONRequestBody = Car

func IsValidVin(vin Vin) bool {
	match, err := regexp.MatchString("^[A-HJ-NPR-Z0-9]{13}[0-9]{4}$", vin)
	if err != nil {
		return false
	}
	return match
}
