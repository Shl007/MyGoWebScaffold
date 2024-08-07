package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	S    *dataStore
)

type dataStore struct {
	db *gorm.DB
}

type StoreItf interface {
	DB() *gorm.DB
}

func NewStore(db *gorm.DB) *dataStore {
	once.Do(func() {
		S = &dataStore{db}
	})
	return S
}

func (ds *dataStore) DB() *gorm.DB {
	return ds.db
}
