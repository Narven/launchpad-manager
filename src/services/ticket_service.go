package services

import (
	"github.com/Narven/launchpad-manager/src/domain/tickets"
	"github.com/Narven/launchpad-manager/src/utils/errs"
)

var (
	TicketService ticketServiceInterface = &ticketService{}
)

type ticketService struct {
}

type ticketServiceInterface interface {
	CreateTicket(ticket tickets.CreateTicketRequestDto) (*tickets.Ticket, *errs.RestErr)
}

func (s *ticketService) CreateTicket(ticketDto tickets.CreateTicketRequestDto) (*tickets.Ticket, *errs.RestErr) {
	// TODO add validation

	var ticket tickets.Ticket
	ticket.FirstName = ticketDto.FirstName
	ticket.LastName = ticketDto.LastName
	ticket.Birthday = ticketDto.Birthday
	ticket.LaunchpadID = ticketDto.LaunchpadID
	ticket.DestinationID = ticketDto.DestinationID
	ticket.LaunchDate = ticketDto.LaunchDate

	if err := ticket.Save(); err != nil {
		return nil, err
	}

	return &ticket, nil
}
