package db

import (
	"backend/models/memory"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func Open(dsnURI string) (DB, error) {
	if db, err := sql.Open("sqlite3", dsnURI); err == nil {
		return DB{
			db: db,
		}, nil
	} else {
		return DB{}, fmt.Errorf("DB Failed to open: %w", err)
	}
}

func (db DB) CreateTables() error {
	stmt, err := db.db.Prepare("CREATE TABLE IF NOT EXISTS memory(text TEXT, timestamp INTEGER, geolocation_status TEXT, latitude REAL, longitude REAL)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (db DB) InsertMemory(memory memory.Memory) error {
	stmt, err := db.db.Prepare("INSERT INTO memory (text, timestamp, geolocation_status, latitude, longitude) VALUES(?,?,?,?,?) ")
	if err != nil {
		return fmt.Errorf("DB failed to insert memory: %w", err)
	}

	res, err := stmt.Exec(memory.Text, memory.Timestamp, memory.GeolocationStatus, memory.Latitude, memory.Longitude)
	if err != nil {
		return fmt.Errorf("DB failed to insert memory: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("DB failed to insert memory: %w", err)
	}

	log.Printf("Inserted memory: %d", id)

	return nil
}

func (db DB) SelectMemories() ([]memory.Memory, error) {
	rows, err := db.db.Query("SELECT text, timestamp, geolocation_status, latitude, longitude FROM memory")
	defer rows.Close()
	if err != nil {
		return []memory.Memory{},
			fmt.Errorf("DB failed to select memories: %w", err)

	}

	results := make([]memory.Memory, 0)

	for rows.Next() {
		m := memory.Memory{}
		if err := rows.Scan(&m.Text, &m.Timestamp, &m.GeolocationStatus, &m.Latitude, &m.Longitude); err != nil {
			return results, fmt.Errorf("DB failed to select memories: %w", err)

		}
		results = append(results, m)
	}
	return results, nil
}

func (db DB) Close() {
	db.db.Close()
}
