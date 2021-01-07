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
	GetTickets() (*[]tickets.Ticket, *errs.RestErr)
	DeleteTicket(id int64) *errs.RestErr
}

func (s *ticketService) CreateTicket(ticketDto tickets.CreateTicketRequestDto) (*tickets.Ticket, *errs.RestErr) {
	var ticket tickets.Ticket
	ticket.FirstName = ticketDto.FirstName
	ticket.LastName = ticketDto.LastName
	ticket.Birthday = ticketDto.Birthday
	ticket.Gender = ticketDto.Gender
	ticket.LaunchpadID = ticketDto.LaunchpadID
	ticket.DestinationID = ticketDto.DestinationID
	ticket.LaunchDate = ticketDto.LaunchDate

	if err := ticket.Save(); err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (s *ticketService) GetTickets() (*[]tickets.Ticket, *errs.RestErr) {
	var ticket tickets.Ticket
	ticks, err := ticket.GetAll()
	if err != nil {
		return nil, err
	}

	return ticks, nil
}

func (s *ticketService) DeleteTicket(id int64) *errs.RestErr {
	ticket := tickets.Ticket{ID: id}
	err := ticket.Delete()
	if err != nil {
		return err
	}

	return nil
}
