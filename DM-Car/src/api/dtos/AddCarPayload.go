package dtos

type AddCarPayload struct {
	Vin   string `json:"vin"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}
