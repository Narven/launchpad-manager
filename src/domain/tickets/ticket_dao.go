package tickets

import (
	db "github.com/Narven/launchpad-manager/src/datasources/psql/launchpadmanagerdb"
	"github.com/Narven/launchpad-manager/src/logger"
	"github.com/Narven/launchpad-manager/src/utils/errs"
)

const (
	querySaveTicket = "INSERT INTO ticket(first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date) VALUES (?,?,?,?,?,?,?)"
)

func (ticket *Ticket) Save() *errs.RestErr {
	stmt, err := db.Client.Prepare(querySaveTicket)
	if err != nil {
		logger.Error("database error", err)
		return errs.NewBadRequestError("dataase error")
	}
	defer stmt.Close()

	_, saveErr := stmt.Exec(
		ticket.FirstName,
		ticket.LastName,
		ticket.Gender,
		ticket.Birthday,
		ticket.LaunchpadID,
		ticket.DestinationID,
		ticket.LaunchDate,
	)
	if saveErr != nil {
		logger.Error("database error", saveErr)
		return errs.NewBadRequestError("dataase error")
	}

	return nil
}
