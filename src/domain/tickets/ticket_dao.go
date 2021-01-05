package tickets

import (
	db "github.com/Narven/launchpad-manager/src/datasources/psql/launchpadmanager"
	"github.com/Narven/launchpad-manager/src/logger"
	"github.com/Narven/launchpad-manager/src/utils/errs"
)

const (
	querySaveTicket = "INSERT INTO ticket (first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id"
)

func (ticket *Ticket) Save() *errs.RestErr {
	tx := db.Client.MustBegin()

	//stmt, err := tx.Prepare(querySaveTicket)
	//if err != nil {
	//	logger.Error("database error", err)
	//	return errs.NewBadRequestError("database error")
	//}
	//defer stmt.Close()

	var id int64
	_ = tx.QueryRowx(querySaveTicket,
		ticket.FirstName,
		ticket.LastName,
		ticket.Gender,
		ticket.Birthday,
		ticket.LaunchpadID,
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
