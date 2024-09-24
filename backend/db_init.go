package main

import (
	"database/sql"
)

func initDB(db *sql.DB) error {
	if err := executeQuery(db, "./queries/create_polls_table.sql"); err != nil {
		return err
	}
	if err := executeQuery(db, "./queries/create_options_table.sql"); err != nil {
		return err
	}
	if err := executeQuery(db, "./queries/create_votes_table.sql"); err != nil {
		return err
	}
	return nil
}
