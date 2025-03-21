package models

import (
	"database/sql"
	"errors"
	"time"
)

// This Climb type holds the data for an individual climb.
// Wew use sql.NullString types incase we get a null value in our database.
// Null cannot be converted to a string and an error will be thrown.
// Using sql.NullString the each string has a Valid bool. In the case of null the string
// will be empty with the Valid bool set to false
// Ideally our database would never have a null. Only allowing valid entries and proper
// Default values where needed
type Climb struct {
	ID       int
	Title    sql.NullString
	Category sql.NullString
	Grade    sql.NullString
	Setter   sql.NullString
	Created  time.Time
}

// Define a Connection type which wraps a sql.DB connection pool.
type Connection struct {
	DB *sql.DB
}

// This function will insert a new climb into the database.
func (conn *Connection) Insert(title string, category string, grade string, setter string) (int, error) {

	transaction, err := conn.DB.Begin()
	if err != nil {
		return 0, err
	}

	defer transaction.Rollback()

	// Create the query with placeholders so we can pass it our values
	query := `INSERT INTO climbs (title, category, grade, setter, created)
    VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING ID`

	// Execute the query passing in our values then checking for errors
	// The LastInsertId() function in the sql.Result is unsupported by out driver.
	// So we use Query row instead of .Exec because this query returns the ID for the row of the inserted data
	// Using .Scan to get the data from the returned Row type and copy it to our insertId variable
	// The go database/sql lib usused prepared statements under the hood for .Exec .Query .QueryRow so they don't
	// have to be done manually
	insertId := 0
	err = transaction.QueryRow(query, title, category, grade, setter).Scan(&insertId)
	if err != nil {
		return 0, err
	}

	err = transaction.Commit()
	if err != nil {
		return 0, err
	}

	return int(insertId), nil
}

// This will return a specific snippet based on its id.
func (conn *Connection) Get(id int) (Climb, error) {

	// Create the query with placeholders so we can pass it our values
	query := `SELECT id, title, category, grade, setter, created FROM climbs WHERE id = $1`

	// Execute the query
	row := conn.DB.QueryRow(query, id)

	// Create  a new clib struct to hold the data from our Row
	var climb Climb

	//Scan the data into our climb struct and check for errors
	err := row.Scan(&climb.ID, &climb.Title, &climb.Category, &climb.Grade, &climb.Setter, &climb.Created)
	if err != nil {
		// If the query doesn't return a row, row.Scan will return a sql.ErrNoRows err
		// Check for sql.ErrNoRows and return the according error
		if errors.Is(err, sql.ErrNoRows) {
			return Climb{}, ErrNoRecord
		} else {
			return Climb{}, err
		}
	}

	return climb, nil
}

// This will return the 10 most recently created snippets.
func (conn *Connection) Latest() ([]Climb, error) {

	query := `Select id, title, category, grade, setter, created FROM climbs order by ID DESC LIMIT 10`

	// Execute the query and check for errors
	rows, err := conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	// Defer our rows to close at the end of the function. Without this the pool connection would stay open
	// This can lead to all the connections filling up
	defer rows.Close()

	var climbs []Climb

	// Iterate through the rows. Reading in the data and building the array of Climb structs
	for rows.Next() {
		var climb Climb

		err = rows.Scan(&climb.ID, &climb.Title, &climb.Category, &climb.Grade, &climb.Setter, &climb.Created)
		if err != nil {
			return nil, err
		}

		climbs = append(climbs, climb)
	}

	// Check if the rows.Next() loop closed abnormally instead of finishing
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	// The above can also be written in one line
	// if err = rows.Err(); err != nil {}

	return climbs, nil
}

// This will return the 10 most recently created snippets.
func (conn *Connection) JsonRequest(quantity int) ([]Climb, error) {

	query := `Select id, title, category, grade, setter, created FROM climbs order by ID DESC LIMIT $1`

	// Execute the query and check for errors
	rows, err := conn.DB.Query(query, quantity)
	if err != nil {
		return nil, err
	}
	// Defer our rows to close at the end of the function. Without this the pool connection would stay open
	// This can lead to all the connections filling up
	defer rows.Close()

	var climbs []Climb

	// Iterate through the rows. Reading in the data and building the array of Climb structs
	for rows.Next() {
		var climb Climb

		err = rows.Scan(&climb.ID, &climb.Title, &climb.Category, &climb.Grade, &climb.Setter, &climb.Created)
		if err != nil {
			return nil, err
		}

		climbs = append(climbs, climb)
	}

	// Check if the rows.Next() loop closed abnormally instead of finishing
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	// The above can also be written in one line
	// if err = rows.Err(); err != nil {}

	return climbs, nil
}

func (conn *Connection) ExampleTransaction() error {

	// Start the transaction and check for errors
	transaction, err := conn.DB.Begin()
	if err != nil {
		return err
	}

	// Defer a transaction rollback until the function close
	// If we don't have any errors/issues and do a transaction.Commit()
	// This deferd rollback won't do anything. However if we exit the function abnoramlly
	// after encountering errors our transaction will be rolled back
	defer transaction.Rollback()

	// Transactions must always Rollback or Commit before the function returns or the
	// connection will stay open leading to hitting the maximum connection limit

	// Call the Exec function on the transaction instead of the connection pool and check for errors
	_, err = transaction.Exec("INSERT INTO ...")
	if err != nil {
		return err
	}

	// Multiple querys can be done within the same transaction
	_, err = transaction.Exec("DELETE ...")
	if err != nil {
		return err
	}

	// Commit the transaction returning any errors
	err = transaction.Commit()
	return err

}
