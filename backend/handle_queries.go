package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func executeQuery(filepath string, isFatal bool, args ...interface{}) error {
	query, err := os.ReadFile(filepath)
	if err != nil {
		if isFatal {
			log.Fatal(err)
		} else {
			return fmt.Errorf("Error reading sql. %w", err)
		}
	}

	_, err = db.Exec(string(query), args...)
	if err != nil {
		if isFatal {
			log.Fatal(err)
		} else {
			return fmt.Errorf("Error executing sql. %w", err)
		}
	}
	return nil
}

func insertPoll(poll Poll) (uuid.UUID, error) {
	id := uuid.New()
	err := executeQuery("./queries/insert_poll.sql", false, id, poll.Duration, poll.Title)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func insertOptions(pollID uuid.UUID, options []Option) error {
	for _, option := range options {
		err := executeQuery("./queries/insert_option.sql", false, pollID, option.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

// could refactor
func getPoll(pollID uuid.UUID) (Poll, error) {
	var poll Poll

	queryInBytes1, err := os.ReadFile("./queries/select_poll_metadata.sql")
	if err != nil {
		return Poll{}, fmt.Errorf("Error reading sql query for select poll title. %w", err)
	}

	err = db.QueryRow(string(queryInBytes1), pollID).Scan(&poll.Title, &poll.Duration, &poll.CreatedAt)
	if err != nil {
		return Poll{}, fmt.Errorf("Error executing sql query for select poll title. %w", err)
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
		err = rows.Scan(&option.ID, &option.Name, &option.Votes)
		if err != nil {
			return Poll{}, fmt.Errorf("Error scanning row from selected options. %w", err)
		}
		poll.Options = append(poll.Options, option)
	}

	return poll, nil
}

// could refactor
func updateVotes(vote Vote) error {

	err := executeQuery("./queries/insert_vote.sql", false, vote.VoterID, vote.OptionID, vote.PollID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Constraint == "unique_vote_per_poll" && pqErr.Code == "23505" {
				return fmt.Errorf("Vote from this user already exists on poll %v: %w", vote.PollID, err)
			}
		}
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Error beginning transaction: %w", err)
	}
	defer tx.Rollback()

	queryInBytes2, err := os.ReadFile("./queries/select_votes_for_update.sql")
	if err != nil {
		return fmt.Errorf("Error reading sql query selecting for update. %w", err)
	}

	_, err = tx.Exec(string(queryInBytes2), vote.OptionID)
	if err != nil {
		return fmt.Errorf("Error executing sql query select for update. %w", err)
	}

	queryInBytes3, err := os.ReadFile("./queries/update_num_of_votes.sql")
	if err != nil {
		return fmt.Errorf("Error reading sql query for updating number of votes. %w", err)
	}

	_, err = tx.Exec(string(queryInBytes3), vote.OptionID)
	if err != nil {
		return fmt.Errorf("Error executing sql query for updating number of votes. %w", err)
	}

	return tx.Commit()
}
