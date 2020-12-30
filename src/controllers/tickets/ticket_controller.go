package tickets

import (
	"errors"
	"github.com/Narven/launchpad-manager/src/domain/tickets"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var ticketRequest tickets.CreateTicketRequestDto
	if err := c.ShouldBindJSON(&ticketRequest); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid payload"))
		return
	}

	c.JSON(http.StatusOK, errors.New("all ok"))
}
