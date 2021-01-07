package tickets

import (
	"encoding/json"
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

type SpaceXLaunch struct {
	FlightID      string `json:"flight_id"`
	LaunchDateUTC string `json:"launch_date_utc"`
	SiteID        string `json:"site_id"`
}

const (
	spaceXApi = "https://api.spacexdata.com/v3"
)

func Create(c *gin.Context) {
	var ticketRequest tickets.CreateTicketRequestDto
	if err := c.ShouldBindJSON(&ticketRequest); err != nil {
		logger.Error("could not bind to json", err)
		restErr := errs.NewBadRequestError("invalid payload")
		c.JSON(restErr.Status, restErr)
		return
	}

	t, err := time.Parse("2006-01-02T15:04:000Z", ticketRequest.LaunchDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error while handling launch date")
		return
	}
	destination, getDestinationErr := services.DestionationService.GetDestination(ticketRequest.DestinationID)
	if getDestinationErr != nil {
		c.JSON(getDestinationErr.Status, getDestinationErr)
		return
	}

	if time.Weekday(destination.Weekday) != t.Weekday() {
		// throw error (destination is not possible on that date)
		c.JSON(http.StatusBadRequest, "Destination not possible on that date")
		return
	}

	// Get landing pad
	url := fmt.Sprintf(
		"%s/launchpads/%s",
		spaceXApi,
		ticketRequest.LaunchpadID,
	)
	res, getErr := http.Get(url)
	if getErr != nil {
		logger.Info(fmt.Sprintf("launchpad is not available: %s", ticketRequest.LaunchpadID))
		c.JSON(http.StatusBadRequest, "launchpad not available")
		return
	}
	defer res.Body.Close()

	// Check agains SpaceX upcomming launchs
	url = fmt.Sprintf("%s/launches/upcoming?site_id=%s&launch_date_utc=%s",
		spaceXApi,
		ticketRequest.LaunchpadID,
		ticketRequest.LaunchDate,
	)

	logger.Info(url)

	// if we get and error, is because we did not found an SpaceX launch
	// so we can book the ticket
	resLaunches, getErr := http.Get(url)
	if getErr != nil {
		// TODO what do we want to do in case we cannot get SpaceX information?
		//  I choose to fail to prevent using the same launchpads
		c.JSON(http.StatusInternalServerError, "spacex api error")
		return
	}
	defer resLaunches.Body.Close()

	// we need to check the result because this endpoint always returns 200
	// we need to check the body
	launchs := make([]SpaceXLaunch, 0)
	marshallErr := json.NewDecoder(resLaunches.Body).Decode(&launchs)
	if marshallErr != nil {
		c.JSON(http.StatusInternalServerError, "spacex api error")
		return
	}

	// if there is launchs we cannot create the ticket
	if len(launchs) != 0 {
		c.JSON(http.StatusBadRequest, "launchpad not available")
		return
	}

	// All should be fine to create the ticket

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
