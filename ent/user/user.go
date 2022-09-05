// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldVerifyCode holds the string denoting the verifycode field in the database.
	FieldVerifyCode = "verify_code"
	// FieldVerifyExpiry holds the string denoting the verifyexpiry field in the database.
	FieldVerifyExpiry = "verify_expiry"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPasshash holds the string denoting the passhash field in the database.
	FieldPasshash = "passhash"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeGame holds the string denoting the game edge name in mutations.
	EdgeGame = "game"
	// Table holds the table name of the user in the database.
	Table = "users"
	// GameTable is the table that holds the game relation/edge.
	GameTable = "games"
	// GameInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GameInverseTable = "games"
	// GameColumn is the table column denoting the game relation/edge.
	GameColumn = "user_game"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldVerifyCode,
	FieldVerifyExpiry,
	FieldEmail,
	FieldPasshash,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"guess_team",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultVerifyCode holds the default value on creation for the "verifyCode" field.
	DefaultVerifyCode func() string
	// DefaultVerifyExpiry holds the default value on creation for the "verifyExpiry" field.
	DefaultVerifyExpiry func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// StatusUnverified is the default value of the Status enum.
const DefaultStatus = StatusUnverified

// Status values.
const (
	StatusUnverified Status = "Unverified"
	StatusVerified   Status = "Verified"
	StatusDisabled   Status = "Disabled"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusUnverified, StatusVerified, StatusDisabled:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}
