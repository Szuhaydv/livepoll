package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func insertPoll(db *sql.DB, duration string, title string) error {
	sqlBytes, err := os.ReadFile("./queries/insert_poll.sql")
	if err != nil {
		return fmt.Errorf("Error reading sql query for insert into polls. %w", err)
	}

	id := uuid.New()

	_, err = db.Exec(string(sqlBytes), id, duration, title)
	if err != nil {
		return fmt.Errorf("Error exectuing sql query for insert into polls. %w", err)

	}

	return nil
}
