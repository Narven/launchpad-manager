package tickets

import "time"

type CreateTicketRequestDto struct {
	FirstName     string    `json:"first_name" binding:"required"`
	LastName      string    `json:"last_name" binding:"required"`
	Gender        string    `json:"gender" binding:"required"`
	Birthday      time.Time `json:"birthday" binding:"required"`
	LaunchpadID   string    `json:"launchpad_id" binding:"required"`
	DestinationID string    `json:"destination_id" binding:"required"`
	LaunchDate    time.Time `json:"launch_date" binding:"required"`
}
