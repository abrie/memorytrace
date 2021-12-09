package datastore

import (
	"backend/db"
	"backend/models/memory"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMemoryHandler(t *testing.T) {
	db, err := db.Open(":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	datastore, err := Open(db)
	if err != nil {
		t.Fatalf("Failed to create Datastore: %v", err)
	}

	memories := []memory.Memory{
		{Text: "memory1", Chain: []string{"first", "second", "third"}, Timestamp: 1, Latitude: 10.5, Longitude: 9.5},
		{Text: "memory2", Chain: []string{"one", "two"}, Timestamp: 2, Latitude: 20.5, Longitude: 19.5},
	}

	w := httptest.NewRecorder()

	for idx := range memories {
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(memories[idx])
		r := httptest.NewRequest(http.MethodPost, "/", buf)

		datastore.MemoryHandler(w, r)
	}

	got, err := datastore.GetMemories()
	if err != nil {
		t.Fatalf("Failed to get memories: %v", err)
	}

	if len(got) != len(memories) {
		t.Fatalf("Expected %d memories, got: %d", len(memories), len(got))
	}

	for idx := range memories {
		if same := reflect.DeepEqual(memories[idx], got[idx]); !same {
			log.Fatalf("Expected %v, Got: %v", memories[idx], got[idx])
		}
	}

}

func TestPutMemory(t *testing.T) {
	db, err := db.Open(":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}

	datastore, err := Open(db)
	if err != nil {
		t.Fatalf("Failed to create Datastore: %v", err)
	}

	memories := []memory.Memory{
		{Text: "memory1", Timestamp: 1, Latitude: 10.5, Longitude: 9.5},
		{Text: "memory2", Timestamp: 2, Latitude: 20.5, Longitude: 19.5},
	}

	for idx := range memories {
		if err := datastore.PutMemory(memories[idx]); err != nil {
			t.Fatalf("Failed to record memory: %v", err)
		}
	}

	got, err := datastore.GetMemories()
	if err != nil {
		t.Fatalf("Failed to get memories: %v", err)
	}

	if len(got) != len(memories) {
		t.Fatalf("Expected %d memories, got: %d", len(memories), len(got))
	}

	for idx := range memories {
		if same := reflect.DeepEqual(memories[idx], got[idx]); !same {
			log.Fatalf("Expected %v, Got: %v", memories[idx], got[idx])
		}
	}
}
