package main

import "io"

// Factory method creational design pattern allows creating objects without having to specify the exact type of the object that will be created.

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
