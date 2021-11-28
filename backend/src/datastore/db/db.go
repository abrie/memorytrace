package db

import "backend/models/memory"

type Interface interface {
	CreateTables() error
	SelectMemories() ([]memory.Memory, error)
	InsertMemory(memory.Memory) error
}
