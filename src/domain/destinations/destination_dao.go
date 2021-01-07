package destinations

import (
	db "github.com/Narven/launchpad-manager/src/datasources/psql/launchpadmanager"
	"github.com/Narven/launchpad-manager/src/utils/errs"
)

const (
	queryGetDestination = "SELECT id, name, weekday FROM destination WHERE id=$1"
)

func (destination *Destination) Get() *errs.RestErr {
	stmt, err := db.Client.Prepare(queryGetDestination)
	if err != nil {
		return errs.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(destination.ID)
	if getErr := result.Scan(
		&destination.ID,
		&destination.Name,
		&destination.Weekday,
	); getErr != nil {
		return errs.NewInternalServerError("database error")
	}

	return nil
}
