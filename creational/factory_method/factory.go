package main

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return nil /* newMemoryStorage() */
	case TempStorage:
		return nil /* newTempStorage() */
	case DiskStorage:
		return nil /* newDiskStorage() */
	default:
		return nil /* newTempStorage() */
	}

	// logic not implemented
	return nil
}
