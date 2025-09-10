package user

import "database/sql"

type Repository struct {
	Conn    *sql.DB
	Queries *db.Queries
}
