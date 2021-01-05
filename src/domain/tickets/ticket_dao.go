package tickets

import (
	"fmt"
	db "github.com/Narven/launchpad-manager/src/datasources/psql/launchpadmanager"
	"github.com/Narven/launchpad-manager/src/logger"
	"github.com/Narven/launchpad-manager/src/utils/errs"
)

const (
	querySaveTicket = "INSERT INTO ticket (first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date) VALUES ($1,$2,$3,$4,$5,$6,$7)"
)

func (ticket *Ticket) Save() *errs.RestErr {
	fmt.Println(db.Client)
	stmt, err := db.Client.Prepare(querySaveTicket)
	if err != nil {
		logger.Error("database error", err)
		return errs.NewBadRequestError("database error")
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
