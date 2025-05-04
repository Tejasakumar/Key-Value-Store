package Storage

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"time"
)

func (db *Db) Get(key string) (*Data, error) {
	db.ddb.mu.RLock()
	data, ok := db.ddb.Store[key]
	db.ddb.mu.RUnlock()
	if !ok {
		return nil, errors.New("data not found")
	}
	return data, nil
}

func (db *Db) GetAll() (map[string]*Data, error) {
	return db.ddb.Store, nil
}

func (db *Db) Put(key string, value interface{}) (string, interface{}) {
	db.ddb.mu.Lock()
	defer db.ddb.mu.Unlock()

	db.ddb.Store[key] = &Data{
		ID:        uuid.New(),
		Data:      value,
		WriteTime: time.Now(),
	}
	return key, value
}

func (db *Db) Delete(key string) error {
	db.ddb.mu.Lock()
	_, ok := db.ddb.Store[key]
	if !ok {
		return errors.New("data not found")
	}
	delete(db.ddb.Store, key)
	db.ddb.mu.Unlock()
	return nil
}

func (db *Db) Dropdata() {
	db.ddb.mu.RLock()
	defer db.ddb.mu.RUnlock()
	for key, value := range db.ddb.Store {
		value.Data = nil
		value = nil
		db.ddb.Store[key] = value
		delete(db.ddb.Store, key)
	}

}

func (db *Db) Setttl(key string, ttl int) error {
	db.tdb.mu.Lock()
	defer db.tdb.mu.Unlock()

	data, err := db.Get(key)
	if err != nil {
		fmt.Println("Data not found" + err.Error())
		return err
	}
	ttlData := &TTL{
		key:  key,
		Data: data,
		ttl:  time.Now().Add(time.Duration(ttl) * time.Second),
	}
	db.tdb.Store[key] = ttlData
	db.link.Add(ttlData)
	return nil
}

func (db *Db) Getttl(key string) (int, error) {
	db.tdb.mu.RLock()
	defer db.tdb.mu.RUnlock()

	data, ok := db.tdb.Store[key]
	if !ok {
		return -1, errors.New("data not found")
	}
	return int(time.Until(data.ttl).Seconds()), nil
}

func (db *Db) Rmttl(key string) error {
	db.tdb.mu.Lock()
	defer db.tdb.mu.Unlock()

	_, ok := db.tdb.Store[key]
	if !ok {
		return errors.New("data not found")
	}
	delete(db.tdb.Store, key)
	fmt.Println("Deleted key from ttl db: " + key)
	return nil
}

func (db *Db) Updatettldb(key string, ttl int) string {
	db.tdb.mu.Lock()
	defer db.tdb.mu.Unlock()
	db.tdb.Store[key].ttl = time.Now().Add(time.Duration(ttl) * time.Second)
	db.link.Delete(key)
	db.link.Add(db.tdb.Store[key])
	return "TTL updated for key: " + key + " to " + strconv.Itoa(ttl)

}

func (db *Db) RemoveTTL(key string) string {
	fmt.Println("Removing ttl for key: " + key)
	db.Rmttl(key)
	db.link.Delete(key)
	return "TTL removed for key: " + key
}

func (db *Db) DropDb() {
	go db.Dropdata()
	db.tdb.Store = make(map[string]*TTL)
	db.link.Head = nil
	db.link.Tail = nil
}

func (db *Db) Keys() []string {
	result := make([]string, 0)
	for key := range db.ddb.Store {
		result = append(result, key)
	}
	return result
}
