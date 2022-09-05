// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/greboid/puzzad/ent/guess"
)

// Guess is the model entity for the Guess schema.
type Guess struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GuessQuery when eager-loading is set.
	Edges GuessEdges `json:"edges"`
}

// GuessEdges holds the relations/edges for other nodes in the graph.
type GuessEdges struct {
	// Question holds the value of the question edge.
	Question []*Question `json:"question,omitempty"`
	// Team holds the value of the team edge.
	Team []*User `json:"team,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// QuestionOrErr returns the Question value or an error if the edge
// was not loaded in eager-loading.
func (e GuessEdges) QuestionOrErr() ([]*Question, error) {
	if e.loadedTypes[0] {
		return e.Question, nil
	}
	return nil, &NotLoadedError{edge: "question"}
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading.
func (e GuessEdges) TeamOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Guess) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case guess.FieldID:
			values[i] = new(sql.NullInt64)
		case guess.FieldContent:
			values[i] = new(sql.NullString)
		case guess.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Guess", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Guess fields.
func (gu *Guess) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case guess.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gu.ID = int(value.Int64)
		case guess.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				gu.CreateTime = value.Time
			}
		case guess.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				gu.Content = value.String
			}
		}
	}
	return nil
}

// QueryQuestion queries the "question" edge of the Guess entity.
func (gu *Guess) QueryQuestion() *QuestionQuery {
	return (&GuessClient{config: gu.config}).QueryQuestion(gu)
}

// QueryTeam queries the "team" edge of the Guess entity.
func (gu *Guess) QueryTeam() *UserQuery {
	return (&GuessClient{config: gu.config}).QueryTeam(gu)
}

// Update returns a builder for updating this Guess.
// Note that you need to call Guess.Unwrap() before calling this method if this Guess
// was returned from a transaction, and the transaction was committed or rolled back.
func (gu *Guess) Update() *GuessUpdateOne {
	return (&GuessClient{config: gu.config}).UpdateOne(gu)
}

// Unwrap unwraps the Guess entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gu *Guess) Unwrap() *Guess {
	_tx, ok := gu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Guess is not a transactional entity")
	}
	gu.config.driver = _tx.drv
	return gu
}

// String implements the fmt.Stringer.
func (gu *Guess) String() string {
	var builder strings.Builder
	builder.WriteString("Guess(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gu.ID))
	builder.WriteString("create_time=")
	builder.WriteString(gu.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(gu.Content)
	builder.WriteByte(')')
	return builder.String()
}

// Guesses is a parsable slice of Guess.
type Guesses []*Guess

func (gu Guesses) config(cfg config) {
	for _i := range gu {
		gu[_i].config = cfg
	}
}
