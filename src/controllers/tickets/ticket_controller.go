package tickets

import (
	"fmt"
	"github.com/Narven/launchpad-manager/src/domain/tickets"
	"github.com/Narven/launchpad-manager/src/logger"
	"github.com/Narven/launchpad-manager/src/services"
	"github.com/Narven/launchpad-manager/src/utils/errs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Create(c *gin.Context) {
	var ticketRequest tickets.CreateTicketRequestDto
	if err := c.ShouldBindJSON(&ticketRequest); err != nil {
		logger.Error("could not bind to json", err)
		restErr := errs.NewBadRequestError("invalid payload")
		c.JSON(restErr.Status, restErr)
		return
	}

	t, err := time.Parse("01-02-2006", ticketRequest.LaunchDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error while handling launch date")
		return
	}
	destination, getDestinationErr := services.DestionationService.GetDestination(ticketRequest.DestinationID)
	if getDestinationErr != nil {
		c.JSON(getDestinationErr.Status, getDestinationErr)
		return
	}

	fmt.Println(time.Weekday(destination.Weekday))
	fmt.Println(t.Weekday())

	if time.Weekday(destination.Weekday) != t.Weekday() {
		// throw error (destination is not possible on that date)
		c.JSON(http.StatusBadRequest, "Destination not possible on that date")
		return
	}

	// [x] Get one landing pad (https://api.spacexdata.com/v3/launchpads/{{site_id}})
	url := fmt.Sprintf("https://api.spacexdata.com/v3/launchpads/%s", ticketRequest.LaunchpadID)
	res, getErr := http.Get(url)
	if getErr != nil {
		c.JSON(http.StatusBadRequest, "launchpad not available")
		return
	}
	defer res.Body.Close()

	// TODO [] Check agains SpaceX launchs
	url = fmt.Sprintf(
		"https://api.spacexdata.com/v3/launches/upcoming?site_id=%s&launch_date_utc=%s",
		ticketRequest.LaunchpadID,
		"2020-12-06T16:17:00.000Z",
	)
	res, getErr = http.Get(url)
	if getErr != nil {
		c.JSON(http.StatusBadRequest, "launchpad already reserved")
		return
	}
	defer res.Body.Close()

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

func Delete(c *gin.Context) {
	ticketID, paramErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if paramErr != nil {
		c.JSON(http.StatusBadRequest, "bad params")
	}

	err := services.TicketService.DeleteTicket(ticketID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
