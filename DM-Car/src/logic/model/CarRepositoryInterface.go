package model

type CarRepositoryInterface interface {
	Save(car Car) error
	FindAll() ([]Car, error)
	FindByVin(vin string) (Car, error)
}
