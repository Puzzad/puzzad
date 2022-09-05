// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/greboid/puzzad/ent/adventure"
)

// Adventure is the model entity for the Adventure schema.
type Adventure struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdventureQuery when eager-loading is set.
	Edges AdventureEdges `json:"edges"`
}

// AdventureEdges holds the relations/edges for other nodes in the graph.
type AdventureEdges struct {
	// Game holds the value of the game edge.
	Game []*Game `json:"game,omitempty"`
	// Puzzles holds the value of the puzzles edge.
	Puzzles []*Puzzle `json:"puzzles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GameOrErr returns the Game value or an error if the edge
// was not loaded in eager-loading.
func (e AdventureEdges) GameOrErr() ([]*Game, error) {
	if e.loadedTypes[0] {
		return e.Game, nil
	}
	return nil, &NotLoadedError{edge: "game"}
}

// PuzzlesOrErr returns the Puzzles value or an error if the edge
// was not loaded in eager-loading.
func (e AdventureEdges) PuzzlesOrErr() ([]*Puzzle, error) {
	if e.loadedTypes[1] {
		return e.Puzzles, nil
	}
	return nil, &NotLoadedError{edge: "puzzles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Adventure) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case adventure.FieldID:
			values[i] = new(sql.NullInt64)
		case adventure.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Adventure", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Adventure fields.
func (a *Adventure) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case adventure.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case adventure.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		}
	}
	return nil
}

// QueryGame queries the "game" edge of the Adventure entity.
func (a *Adventure) QueryGame() *GameQuery {
	return (&AdventureClient{config: a.config}).QueryGame(a)
}

// QueryPuzzles queries the "puzzles" edge of the Adventure entity.
func (a *Adventure) QueryPuzzles() *PuzzleQuery {
	return (&AdventureClient{config: a.config}).QueryPuzzles(a)
}

// Update returns a builder for updating this Adventure.
// Note that you need to call Adventure.Unwrap() before calling this method if this Adventure
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Adventure) Update() *AdventureUpdateOne {
	return (&AdventureClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Adventure entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Adventure) Unwrap() *Adventure {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Adventure is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Adventure) String() string {
	var builder strings.Builder
	builder.WriteString("Adventure(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Adventures is a parsable slice of Adventure.
type Adventures []*Adventure

func (a Adventures) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
