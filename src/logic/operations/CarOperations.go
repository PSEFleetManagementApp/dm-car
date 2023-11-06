package operations

import (
	"car/logic/model"
)

type CarOperations struct {
	repository model.ConnectedCarInterface
}

func NewCarOperations(repository model.ConnectedCarInterface) CarOperations {
	return CarOperations{repository: repository}
}
