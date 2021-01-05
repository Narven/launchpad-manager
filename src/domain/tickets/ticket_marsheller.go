package tickets

type CreateTicketRequestDto struct {
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Gender        string `json:"gender" binding:"required"`
	Birthday      string `json:"birthday" binding:"required"`
	LaunchpadID   int64  `json:"launchpad_id" binding:"required"`
	DestinationID int64  `json:"destination_id" binding:"required"`
	LaunchDate    string `json:"launch_date" binding:"required"`
}

type TicketResponseDto struct {
	ID            int64  `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Gender        string `json:"gender"`
	Birthday      string `json:"birthday"`
	LaunchpadID   int64  `json:"launchpad_id"`
	DestinationID int64  `json:"destination_id"`
	LaunchDate    string `json:"launch_date"`
}
