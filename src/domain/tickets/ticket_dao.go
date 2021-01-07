package tickets

import (
	"fmt"
	db "github.com/Narven/launchpad-manager/src/datasources/psql/launchpadmanager"
	"github.com/Narven/launchpad-manager/src/logger"
	"github.com/Narven/launchpad-manager/src/utils/errs"
)

const (
	querySaveTicket    = "INSERT INTO ticket (first_name, last_name, gender, birthday, destination_id, launchpad_id, launch_date) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id"
	queryGetAllTickets = "SELECT * FROM ticket"
	queryDeleteTicket  = "DELETE FROM ticket WHERE id=$1"
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
		ticket.LaunchpadID,
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

func (ticket *Ticket) GetAll() (*[]Ticket, *errs.RestErr) {
	var tickets []Ticket
	err := db.Client.Select(&tickets, queryGetAllTickets)
	if err != nil {
		fmt.Println(err)
		return nil, errs.NewBadRequestError("something wrong while gettting tickets")
	}

	return &tickets, nil
}

func (ticket *Ticket) Delete() *errs.RestErr {
	_, err := db.Client.Exec(queryDeleteTicket, ticket.ID)
	if err != nil {
		fmt.Println(err)
		return errs.NewBadRequestError("something wrong while deleting ticket")
	}

	return nil
}
