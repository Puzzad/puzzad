// Code generated by ent, DO NOT EDIT.

package guess

import (
	"time"
)

const (
	// Label holds the string label denoting the guess type in the database.
	Label = "guess"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldSubmitted holds the string denoting the submitted field in the database.
	FieldSubmitted = "submitted"
	// EdgeQuestion holds the string denoting the question edge name in mutations.
	EdgeQuestion = "question"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// Table holds the table name of the guess in the database.
	Table = "guesses"
	// QuestionTable is the table that holds the question relation/edge.
	QuestionTable = "questions"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "questions"
	// QuestionColumn is the table column denoting the question relation/edge.
	QuestionColumn = "guess_question"
	// TeamTable is the table that holds the team relation/edge.
	TeamTable = "teams"
	// TeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamInverseTable = "teams"
	// TeamColumn is the table column denoting the team relation/edge.
	TeamColumn = "guess_team"
)

// Columns holds all SQL columns for guess fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldContent,
	FieldSubmitted,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultSubmitted holds the default value on creation for the "submitted" field.
	DefaultSubmitted time.Time
)
