package model

type CarRepositoryInterface interface {
	Save(car Car) error
	FindByVin(vin string) (*Car, error)
}
