package db

import (
	"backend/models/memory"
	"database/sql"
	"encoding/json"
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
	stmt, err := db.db.Prepare("CREATE TABLE IF NOT EXISTS memory(text TEXT, chain TEXT, timestamp INTEGER, geolocation_status TEXT, latitude REAL, longitude REAL)")
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
	chainText, err := json.Marshal(memory.Chain)
	if err != nil {
		return fmt.Errorf("DB failed to insert memory: unable to marshal Chain into JSON: %w", err)
	}

	stmt, err := db.db.Prepare("INSERT INTO memory (text, chain, timestamp, geolocation_status, latitude, longitude) VALUES(?,?,?,?,?,?) ")
	if err != nil {
		return fmt.Errorf("DB failed to insert memory: %w", err)
	}

	res, err := stmt.Exec(memory.Text, string(chainText), memory.Timestamp, memory.GeolocationStatus, memory.Latitude, memory.Longitude)
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
	rows, err := db.db.Query("SELECT text, chain, timestamp, geolocation_status, latitude, longitude FROM memory")
	defer rows.Close()
	if err != nil {
		return []memory.Memory{},
			fmt.Errorf("DB failed to select memories: %w", err)

	}

	results := make([]memory.Memory, 0)

	for rows.Next() {
		m := memory.Memory{}
		var chainText string

		if err := rows.Scan(&m.Text, &chainText, &m.Timestamp, &m.GeolocationStatus, &m.Latitude, &m.Longitude); err != nil {
			return results, fmt.Errorf("DB failed to select memories: %w", err)

		}

		if err := json.Unmarshal([]byte(chainText), &m.Chain); err != nil {
			return results, fmt.Errorf("DB failed to select memories: Unable to unmarshal chain text: %w", err)
		}

		results = append(results, m)
	}
	return results, nil
}

func (db DB) Close() {
	db.db.Close()
}
