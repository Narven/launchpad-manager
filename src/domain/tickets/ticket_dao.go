package tickets

import (
	db "github.com/Narven/launchpad-manager/src/datasources/psql/launchpadmanager"
	"github.com/Narven/launchpad-manager/src/logger"
	"github.com/Narven/launchpad-manager/src/utils/errs"
)

const (
	querySaveTicket = "INSERT INTO ticket (first_name, last_name, gender, birthday, destination_id, launch_date) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
)

func (ticket *Ticket) Save() *errs.RestErr {
	tx := db.Client.MustBegin()

	var id int64
	_ = tx.QueryRowx(querySaveTicket,
		ticket.FirstName,
		ticket.LastName,
		ticket.Gender,
		ticket.Birthday,
		ticket.DestinationID,
		ticket.LaunchDate,
	).Scan(&id)

	if saveErr := tx.Commit(); saveErr != nil {
		tx.Rollback()
		logger.Error("database error", saveErr)
		return errs.NewBadRequestError("database error")
	}

	ticket.ID = id

	return nil
}
