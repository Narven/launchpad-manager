package tickets

import (
	"errors"
	db "github.com/Narven/launchpad-manager/src/datasources/psql/launchpadmanagerdb"
)

const (
	querySaveTicket = "INSERT INTO ticket(first_name, last_name) VALUES (?,?)"
)

func (ticket *Ticket) Save() error {
	stmt, err := db.Client.Prepare(querySaveTicket)
	if err != nil {
		return errors.New("dasdasdas")
	}
	defer stmt.Close()

	_, saveErr := stmt.Exec(
		ticket.FirstName,
		ticket.LastName,
	)
	if saveErr != nil {
		return saveErr
	}

	return nil
}
