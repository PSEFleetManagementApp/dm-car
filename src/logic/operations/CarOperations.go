package operations

import (
	"car/src/logic/model"
)

type CarOperations struct {
	repository model.CarRepositoryInterface
}

func NewCarOperations(repository model.CarRepositoryInterface) CarOperations {
	return CarOperations{repository: repository}
}
