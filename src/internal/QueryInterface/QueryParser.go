package QueryInterface

import (
	"errors"
	"strconv"
	"strings"
)

// CommandType represents different types of commands
type CommandType int

const (
	UseDB CommandType = iota
	ShowDBs
	CurrentDB
	Keys
	DropDB
	SetTTL
	GetTTL
	Put
	Get
	GetAll
	Delete
	List
	RemoveTTL
	UpdateTTL
	Unknown
)

// Query represents a parsed command
type Query struct {
	Type   CommandType
	DBName string
	Key    string
	Value  string
	TTL    int
}

// ParseQuery parses a string input into a Query struct
func ParseQuery(input string) (Query, error) {
	input = strings.TrimSpace(input)

	// Save the original input for extracting case-sensitive arguments later
	originalInput := input

	// Convert to lowercase only for command parsing
	lowerInput := strings.ToLower(input)

	// Default query with Unknown type
	query := Query{
		Type: Unknown,
	}

	// Split the input into parts using the lowercase version for command detection
	parts := strings.Split(lowerInput, " ")
	if len(parts) == 0 {
		return query, errors.New("empty input")
	}

	// Get the original parts for preserving case in arguments
	originalParts := strings.Split(originalInput, " ")

	// Determine the command type and parse arguments
	switch parts[0] {
	case "use":
		if len(parts) != 2 {
			return query, errors.New("invalid input. Usage: use <db>")
		}
		query.Type = UseDB
		// Use original case for database name
		query.DBName = originalParts[1]

	case "showdbs":
		query.Type = ShowDBs

	case "curdb":
		query.Type = CurrentDB

	case "keys":
		query.Type = Keys

	case "dropdb":
		if len(parts) != 2 {
			return query, errors.New("invalid input. Usage: dropdb <db>")
		}
		query.Type = DropDB
		// Use original case for database name
		query.DBName = originalParts[1]

	case "setttl":
		if len(parts) != 3 {
			return query, errors.New("invalid input. Usage: setttl <key> <ttl>")
		}
		ttl, err := strconv.Atoi(originalParts[2])
		if err != nil {
			return query, errors.New("invalid TTL. Usage: setttl <key> <ttl>")
		}
		query.Type = SetTTL
		query.Key = originalParts[1]
		query.TTL = ttl

	case "getttl":
		if len(parts) != 2 {
			return query, errors.New("invalid input. Usage: getttl <key>")
		}
		query.Type = GetTTL
		query.Key = originalParts[1]

	case "put":
		if len(parts) != 3 {
			return query, errors.New("invalid input. Usage: put <key> <value>")
		}
		query.Type = Put
		query.Key = originalParts[1]
		query.Value = originalParts[2]

	case "get":
		if len(parts) != 2 {
			return query, errors.New("invalid input. Usage: get <key>")
		}
		query.Type = Get
		query.Key = originalParts[1]
	
	case "getall":
		if len(parts) != 1 {
			return query, errors.New("invalid command. Usage: getall")
		}
		query.Type = GetAll

	case "delete":
		if len(parts) != 2 {
			return query, errors.New("invalid input. Usage: delete <key>")
		}
		query.Type = Delete
		query.Key = originalParts[1]

	case "list":
		query.Type = List

	case "rmttl":
		if len(parts) != 2 {
			return query, errors.New("invalid input. Usage: rmttl <key>")
		}
		query.Type = RemoveTTL
		query.Key = originalParts[1]

	case "upttl":
		if len(parts) != 3 {
			return query, errors.New("invalid input. Usage: upttl <key> <ttl>")
		}
		ttl, err := strconv.Atoi(originalParts[2])
		if err != nil {
			return query, errors.New("invalid TTL. Usage: upttl <key> <ttl>")
		}
		query.Type = UpdateTTL
		query.Key = originalParts[1]
		query.TTL = ttl

	default:
		return query, errors.New("unknown command")
	}

	return query, nil
}

// String returns a string representation of a CommandType
func (c CommandType) String() string {
	switch c {
	case UseDB:
		return "UseDB"
	case ShowDBs:
		return "ShowDBs"
	case CurrentDB:
		return "CurrentDB"
	case Keys:
		return "Keys"
	case DropDB:
		return "DropDB"
	case SetTTL:
		return "SetTTL"
	case GetTTL:
		return "GetTTL"
	case Put:
		return "Put"
	case Get:
		return "Get"
	case GetAll:
		return "GetAll"
	case Delete:
		return "Delete"
	case List:
		return "List"
	case RemoveTTL:
		return "RemoveTTL"
	case UpdateTTL:
		return "UpdateTTL"
	default:
		return "Unknown"
	}
}
