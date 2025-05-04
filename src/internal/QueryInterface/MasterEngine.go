package QueryInterface

import (
	"KVS/src/internal/Storage"
	"sync"
)

//master engine manages the list of all dbs

type MasterEngine struct {
	DBs  map[string]*Storage.Db
	Lock sync.RWMutex
}

func (me *MasterEngine) AddDB(name string) {
	me.Lock.Lock()
	defer me.Lock.Unlock()
	me.DBs[name] = Storage.GetDb(name)
}

func (me *MasterEngine) GetDB(name string) (*Storage.Db, bool) {
	me.Lock.RLock()
	defer me.Lock.RUnlock()
	db, ok := me.DBs[name]
	return db, ok
}

func (me *MasterEngine) DeleteDB(name string) {
	me.Lock.Lock()
	defer me.Lock.Unlock()
	delete(me.DBs, name)
}

func GetMasterEngine() *MasterEngine {
	return &MasterEngine{
		DBs:  make(map[string]*Storage.Db),
		Lock: sync.RWMutex{},
	}
}
