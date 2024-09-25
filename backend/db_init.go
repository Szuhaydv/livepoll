package main

func tablesInit() error {
	if err := executeQuery("./queries/create_polls_table.sql", true); err != nil {
		return err
	}
	if err := executeQuery("./queries/create_options_table.sql", true); err != nil {
		return err
	}
	if err := executeQuery("./queries/create_votes_table.sql", true); err != nil {
		return err
	}
	return nil
}
