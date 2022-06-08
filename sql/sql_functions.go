package sql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type SQLDatabase struct {
	db *sql.DB
}

type PersonTable struct {
	// Has to match table Schema as rows.Scan errors
	id   int
	name string
}

func NewSqlDatabase() *SQLDatabase {
	path, _ := os.Getwd()
	databaseFilename := "godatabase.db"

	dbPath := filepath.Join(path, databaseFilename)

	if _, err := os.Stat(dbPath); err == nil {
		fmt.Println("File exists")

	} else if errors.Is(err, os.ErrNotExist) {
		file, _ := os.Create(dbPath)
		file.Close()
	} else {
		fmt.Println(err)
	}

	db, _ := sql.Open("sqlite3", dbPath)

	return &SQLDatabase{
		db: db,
	}
}

func (database *SQLDatabase) CreateTable() {
	fmt.Println("Creating Table")
	query := `CREATE TABLE IF NOT EXISTS Person(id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255))`
	_, err := database.db.Exec(query)

	if err != nil {
		fmt.Println(err)
	}
}

func (database *SQLDatabase) InsertRow() {
	query := "INSERT INTO Person (name) VALUES ('PersonsName')"
	_, err := database.db.Exec(query)

	if err != nil {
		fmt.Println(err)
	}
}

func (database *SQLDatabase) QueryTable() ([]PersonTable, error) {
	query := "SELECT * FROM Person"
	rows, err := database.db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var all []PersonTable
	for rows.Next() {
		var personRow PersonTable
		if errs := rows.Scan(&personRow.id, &personRow.name); errs != nil {
			return nil, errs
		}
		all = append(all, personRow)
	}
	return all, nil
}

func (database *SQLDatabase) DeleteRow() {
	query := "Delete from Person where id = 1"
	_, err := database.db.Exec(query)

	if err != nil {
		fmt.Println(err)
	}
}

func (database *SQLDatabase) AddTableColumn() {
	query := `ALTER TABLE Person ADD COLUMN nickname VARCHAR(255)`
	_, err := database.db.Exec(query)

	if err != nil {
		fmt.Println(err)
	}
}

func (database *SQLDatabase) DropTableColumn() {
	query := `ALTER TABLE Person DROP COLUMN nickname`
	_, err := database.db.Exec(query)

	if err != nil {
		fmt.Println(err)
	}
}
