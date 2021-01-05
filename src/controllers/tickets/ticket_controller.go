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

	c.JSON(http.StatusOK, ticket)
}
