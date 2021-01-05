package tickets

import (
	"github.com/Narven/launchpad-manager/src/domain/tickets"
	"github.com/Narven/launchpad-manager/src/logger"
	"github.com/Narven/launchpad-manager/src/services"
	"github.com/Narven/launchpad-manager/src/utils/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var ticketRequest tickets.CreateTicketRequestDto
	if err := c.ShouldBindJSON(&ticketRequest); err != nil {
		logger.Error("could not bind to json", err)
		restErr := errs.NewBadRequestError("invalid payload")
		c.JSON(restErr.Status, restErr)
		return
	}

	ticket, createErr := services.TicketService.CreateTicket(ticketRequest)
	if createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}

	response := tickets.TicketResponseDto{
		ID:            ticket.ID,
		FirstName:     ticket.FirstName,
		LastName:      ticket.LastName,
		Gender:        ticket.Gender,
		Birthday:      ticket.Birthday,
		LaunchpadID:   ticket.LaunchpadID,
		DestinationID: ticket.DestinationID,
		LaunchDate:    ticket.LaunchDate,
	}

	c.JSON(http.StatusOK, response)
}

func GetAll(c *gin.Context) {
	ticketList, getErr := services.TicketService.GetTickets()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	var response = make([]tickets.TicketResponseDto, 0)

	for _, tick := range *ticketList {
		t := tickets.TicketResponseDto{
			ID:            tick.ID,
			FirstName:     tick.FirstName,
			LastName:      tick.LastName,
			Gender:        tick.Gender,
			Birthday:      tick.Birthday,
			LaunchpadID:   tick.LaunchpadID,
			DestinationID: tick.DestinationID,
			LaunchDate:    tick.LaunchDate,
		}
		response = append(response, t)
	}

	c.JSON(http.StatusOK, response)
}
