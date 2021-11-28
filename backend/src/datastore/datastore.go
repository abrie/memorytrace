package datastore

import (
	"backend/datastore/db"
	"backend/models/memory"
	"net/http"
)

type Datastore struct {
	db db.Interface
}

func New(db db.Interface) (Datastore, error) {
	if err := db.CreateTables(); err != nil {
		return Datastore{}, err
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

func (datastore Datastore) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			datastore.GetHandler(w, r)
			return
		}
		if r.Method == "POST" {
			datastore.PostHandler(w, r)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (datastore Datastore) GetHandler(w http.ResponseWriter, r *http.Request) {
}

func (datastore Datastore) PostHandler(w http.ResponseWriter, r *http.Request) {
}
