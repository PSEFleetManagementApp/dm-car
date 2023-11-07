package operations

import (
	"car/logic/model"
)

type CarOperations struct {
	repository model.ConnectedCarsInterface
}

func NewCarOperations(repository model.ConnectedCarsInterface) CarOperations {
	return CarOperations{repository: repository}
}
