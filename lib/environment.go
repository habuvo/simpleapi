package lib

import "database/sql"

//Env storage different cross-package values
var Env struct {
	DB *sql.DB
}
