package Storage

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

// Data is a struct for storing data with metadata
type Data struct {
	ID        uuid.UUID
	Data      interface{}
	WriteTime time.Time
}

type DataDb struct {
	mu    sync.RWMutex
	Store map[string]*Data
}

type TTL struct {
	key string
	Data *Data
	ttl time.Time
}

type TTLDB struct {
	mu    sync.RWMutex
	Store map[string]*TTL
}

type Db struct {

	ddb		*DataDb
	tdb		*TTLDB
	link	*LinkedList
}

func GetDb(name string) *Db {
	db := &Db{
			ddb: &DataDb{
				Store: make(map[string]*Data),
			},
			tdb: &TTLDB{
				Store: make(map[string]*TTL),
			},
			link: &LinkedList{
				Head: nil,
				Tail: nil,
			},
	}
	go db.link.Sweep(db.tdb, db)
	return db
}