package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/lib/pq"
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

func getPoll(db *sql.DB, pollID uuid.UUID) (Poll, error) {
	var poll Poll

	queryInBytes1, err := os.ReadFile("./queries/select_poll_metadata.sql")
	if err != nil {
		return Poll{}, fmt.Errorf("Error reading sql query for select poll title. %w", err)
	}

	err = db.QueryRow(string(queryInBytes1), pollID).Scan(&poll.Title, &poll.Duration, &poll.CreatedAt)
	if err != nil {
		return Poll{}, fmt.Errorf("Error executing sql query for selecet poll title. %w", err)
	}

	queryInBytes2, err := os.ReadFile("./queries/select_options.sql")
	if err != nil {
		return Poll{}, fmt.Errorf("Error reading sql query for select from options. %w", err)
	}

	rows, execErr := db.Query(string(queryInBytes2), pollID)
	if execErr != nil {
		return Poll{}, fmt.Errorf("Error executing sql query for select from options. %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var option Option
		err = rows.Scan(&option.ID, &option.Name)
		if err != nil {
			return Poll{}, fmt.Errorf("Error scanning row from selected options. %w", err)
		}
		poll.Options = append(poll.Options, option)
	}

	return poll, nil
}

func updateVotes(db *sql.DB, vote Vote) error {

	queryInBytes1, err := os.ReadFile("./queries/insert_vote.sql")
	if err != nil {
		return fmt.Errorf("Error reading sql query for inserting vote. %w", err)
	}

	_, err = db.Exec(string(queryInBytes1), vote.VoterID, vote.OptionID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return fmt.Errorf("Vote from voterID %s already exists: %w", vote.VoterID, err)
		}
		return fmt.Errorf("Error executing sql query for inserting vote: %w", err)
	}

	queryInBytes2, err := os.ReadFile("./queries/update_num_of_votes.sql")
	if err != nil {
		return fmt.Errorf("Error reading sql query for updating number of votes. %w", err)
	}

	_, err = db.Exec(string(queryInBytes2), vote.OptionID)
	if err != nil {
		return fmt.Errorf("Error executing sql query for updating number of votes. %w", err)
	}

	return nil
}
