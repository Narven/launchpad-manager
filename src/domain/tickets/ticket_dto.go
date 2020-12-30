package tickets

type Ticket struct {
	ID            string `db:"id"`
	FirstName     string `db:"first_name"`
	LastName      string `db:"last_name"`
	Gender        string `db:"gender"`
	Birthday      string `db:"birthday"`
	LaunchpadID   string `db:"launchpad_id"`
	DestinationID string `db:"destination_id"`
	LaunchDate    string `db:"launch_date"`
}
