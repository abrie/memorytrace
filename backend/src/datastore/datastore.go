package datastore

import (
	"backend/datastore/db"
	"backend/models/memory"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Datastore struct {
	db db.Interface
}

func Open(db db.Interface) (Datastore, error) {
	if err := db.CreateTables(); err != nil {
		return Datastore{}, fmt.Errorf("Datastore failed to open: %w", err)
	}

	return Datastore{
		db: db,
	}, nil
}

func (datastore Datastore) GetMemories() ([]memory.Memory, error) {
	return datastore.db.SelectMemories()
}

func (datastore Datastore) PutMemory(memory memory.Memory) error {
	return datastore.db.InsertMemory(memory)
}

func (datastore Datastore) MemoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var m memory.Memory
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			log.Printf("Bad Request: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := datastore.PutMemory(m); err != nil {
			log.Printf("Internal Server Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
