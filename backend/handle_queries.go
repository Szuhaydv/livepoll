package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func insertPoll(db *sql.DB, poll Poll) (uuid.UUID, error) {
	sqlBytes, err := os.ReadFile("./queries/insert_poll.sql")
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error reading sql query for insert into polls. %w", err)
	}

	id := uuid.New()

	_, err = db.Exec(string(sqlBytes), id, poll.Duration, poll.Title)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Error executing sql query for insert into polls. %w", err)

	}

	return id, nil
}

func insertOptions(db *sql.DB, pollID uuid.UUID, options []Option) error {
	for _, option := range options {
		sqlBytes, err := os.ReadFile("./queries/insert_option.sql")
		if err != nil {
			return fmt.Errorf("Error reading sql query for insert into options. %w", err)
		}

		_, err = db.Exec(string(sqlBytes), pollID, option.Name)
		if err != nil {
			return fmt.Errorf("Error executing sql query for insert into options. %w", err)
		}
	}
	return nil
}
