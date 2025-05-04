package QueryInterface

import (
	"KVS/src/internal/Storage"
	"fmt"
	"strings"
)

// QueryExecutionEngine handles the execution of parsed queries
type QueryExecutionEngine struct {
	MasterEngine  *MasterEngine
	CurrentDBName string
}

// NewQueryExecutionEngine creates a new QueryExecutionEngine
func NewQueryExecutionEngine(me *MasterEngine) *QueryExecutionEngine {
	return &QueryExecutionEngine{
		MasterEngine:  me,
		CurrentDBName: "",
	}
}

func Replicate(Engine *QueryExecutionEngine) *QueryExecutionEngine {
	return &QueryExecutionEngine{
		MasterEngine:  Engine.MasterEngine,
		CurrentDBName: Engine.CurrentDBName,
	}
}

// Execute executes a query and returns the result as a string
func (e *QueryExecutionEngine) Execute(query Query) string {
	// Handle commands that don't require a current DB
	switch query.Type {
	case UseDB:
		return e.useDatabase(query.DBName)

	case ShowDBs:
		return e.showDatabases()

	case CurrentDB:
		return e.currentDatabase()

	case DropDB:
		return e.dropDatabase(query.DBName)
	}

	// Check if a database is selected for commands that require one
	if len(e.MasterEngine.DBs) == 0 || e.CurrentDBName == "" {
		return "No database selected. Use 'use <db>' to select or create a database."
	}

	// Get the current database
	db, ok := e.MasterEngine.GetDB(e.CurrentDBName)
	if !ok {
		return "Database not found. Use 'use <db>' to select or create a database."
	}

	// Handle database-specific commands
	switch query.Type {
	case Keys:
		return e.listKeys(db)

	case SetTTL:
		return e.setTTL(db, query.Key, query.TTL)

	case GetTTL:
		return e.getTTL(db, query.Key)

	case Put:
		return e.put(db, query.Key, query.Value)

	case Get:
		return e.get(db, query.Key)
	
	case GetAll:
		return e.getAll(db)

	case Delete:
		return e.delete(db, query.Key)

	case List:
		return e.list(db)

	case RemoveTTL:
		return e.removeTTL(db, query.Key)

	case UpdateTTL:
		return e.updateTTL(db, query.Key, query.TTL)

	default:
		return "Unknown command"
	}
}

// useDatabase selects or creates a database
func (e *QueryExecutionEngine) useDatabase(dbName string) string {
	e.CurrentDBName = dbName
	_, ok := e.MasterEngine.GetDB(dbName)
	if !ok {
		e.MasterEngine.AddDB(dbName)
		return fmt.Sprintf("Database not found. Created new database: %s", dbName)
	}
	return fmt.Sprintf("Using database: %s", dbName)
}

// showDatabases lists all available databases
func (e *QueryExecutionEngine) showDatabases() string {
	if len(e.MasterEngine.DBs) == 0 {
		return "No databases found."
	}

	var result strings.Builder
	result.WriteString("Available databases:\n")
	for key := range e.MasterEngine.DBs {
		result.WriteString(fmt.Sprintf("- %s\n", key))
	}
	return result.String()
}

// currentDatabase returns the name of the currently selected database
func (e *QueryExecutionEngine) currentDatabase() string {
	if e.CurrentDBName == "" {
		return "No database selected. Use 'use <db>' to select a database."
	}
	return fmt.Sprintf("Current database: %s", e.CurrentDBName)
}

// dropDatabase drops a database
func (e *QueryExecutionEngine) dropDatabase(dbName string) string {
	db, exists := e.MasterEngine.GetDB(dbName)
	if !exists {
		return fmt.Sprintf("Database not found: %s", dbName)
	}

	db.DropDb()
	e.MasterEngine.DeleteDB(dbName)

	if e.CurrentDBName == dbName {
		e.CurrentDBName = ""
	}

	return fmt.Sprintf("Database dropped: %s", dbName)
}

// listKeys lists all keys in the database
func (e *QueryExecutionEngine) listKeys(db *Storage.Db) string {
	// Capture the output of db.Keys() in a string
	keys := db.Keys()
	if len(keys) == 0 {
		return "No keys found in database."
	}

	var result strings.Builder
	result.WriteString("Keys in database:\n")
	for _, key := range keys {
		result.WriteString(fmt.Sprintf("- %s\n", key))
	}
	return result.String()
}

// setTTL sets the TTL for a key
func (e *QueryExecutionEngine) setTTL(db *Storage.Db, key string, ttl int) string {
	err := db.Setttl(key, ttl)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return fmt.Sprintf("TTL for '%s' set to: %d seconds", key, ttl)
}

// getTTL gets the TTL for a key
func (e *QueryExecutionEngine) getTTL(db *Storage.Db, key string) string {
	ttl, err := db.Getttl(key)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return fmt.Sprintf("TTL for '%s': %d seconds", key, ttl)
}

// put adds or updates a key-value pair
func (e *QueryExecutionEngine) put(db *Storage.Db, key string, value string) string {
	db.Put(key, value)
	return fmt.Sprintf("Value for '%s' set to: %s", key, value)
}

// get retrieves a value for a key
func (e *QueryExecutionEngine) get(db *Storage.Db, key string) string {
	value, err := db.Get(key)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return fmt.Sprintf("Value for '%s': %v", key, value.Data)
}

func (e *QueryExecutionEngine) getAll(db *Storage.Db) string {
	store, err := db.GetAll()
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	outputStr := ""
	for key, val := range store {
		outputStr += fmt.Sprintf("%s: %v\n", key, val.Data)
	}
	return outputStr
}

// delete removes a key-value pair
func (e *QueryExecutionEngine) delete(db *Storage.Db, key string) string {
	err := db.Delete(key)
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return fmt.Sprintf("Deleted key: %s", key)
}

// list shows all key-value pairs in the that have ttl
func (e *QueryExecutionEngine) list(db *Storage.Db) string {
	// We need to get the key-value pairs from the database
	// Assuming db.List() returns key-value pairs
	entries := db.List()
	if len(entries) == 0 {
		return "No entries found in database."
	}

	return fmt.Sprintf("%v", entries)
}

// removeTTL removes the TTL for a key
func (e *QueryExecutionEngine) removeTTL(db *Storage.Db, key string) string {
	return db.RemoveTTL(key)
}

// updateTTL updates the TTL for a key
func (e *QueryExecutionEngine) updateTTL(db *Storage.Db, key string, ttl int) string {
	return db.Updatettldb(key, ttl)
}
