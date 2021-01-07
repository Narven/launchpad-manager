package tickets

type Ticket struct {
	ID            int64  `db:"id"`
	FirstName     string `db:"first_name"`
	LastName      string `db:"last_name"`
	Gender        string `db:"gender"`
	Birthday      string `db:"birthday"`
	LaunchpadID   string `db:"launchpad_id"`
	DestinationID int64  `db:"destination_id"`
	LaunchDate    string `db:"launch_date"`
}
