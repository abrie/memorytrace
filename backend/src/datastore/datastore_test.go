package datastore

import (
	"backend/db"
	"backend/models/memory"
	"log"
	"reflect"
	"testing"
)

func TestPutMemory(t *testing.T) {
	db := db.MustOpen(":memory:")
	datastore, err := New(db)
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
