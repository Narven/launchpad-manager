package services

import (
	"github.com/Narven/launchpad-manager/src/domain/destinations"
	"github.com/Narven/launchpad-manager/src/utils/errs"
)

var (
	DestionationService destinationsServiceInterface = &destinationService{}
)

type destinationService struct {
}

type destinationsServiceInterface interface {
	GetDestination(destinationId int64) (*destinations.Destination, *errs.RestErr)
}

func (s *destinationService) GetDestination(destinationId int64) (*destinations.Destination, *errs.RestErr) {
	destination := &destinations.Destination{ID: destinationId}
	if err := destination.Get(); err != nil {
		return nil, err
	}

	return destination, nil
}
