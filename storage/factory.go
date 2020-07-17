package storage

import "log"

type StorageType int

const (
	//kalau pake iota itu otomatis yang awal =1 kebawahnya auto increment ++ tipedata sama
	//nah yang storagetypemock itu otomatis tipedatanya storagetype terus valuenya 2
	StorageTypeMemory StorageType = iota
	StorageTypeMock
	StorageTypeDatabase
)

func GetStorage(t StorageType) Storage {
	var s Storage
	switch t {
	case StorageTypeMemory:
		s = newMemory()
	case StorageTypeMock:
		s = mock{}
	case StorageTypeDatabase:
		s = structDb{}
	default:
		log.Fatal("storage tipe tidak ditemukan")
	}
	return s
}
