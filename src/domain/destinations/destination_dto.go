package destinations

type Destination struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
