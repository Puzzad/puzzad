// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/greboid/puzzad/ent/team"
)

// Team is the model entity for the Team schema.
type Team struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Status holds the value of the "status" field.
	Status team.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeamQuery when eager-loading is set.
	Edges      TeamEdges `json:"edges"`
	guess_team *int
}

// TeamEdges holds the relations/edges for other nodes in the graph.
type TeamEdges struct {
	// Adventures holds the value of the adventures edge.
	Adventures []*Adventure `json:"adventures,omitempty"`
	// Progress holds the value of the progress edge.
	Progress []*Progress `json:"progress,omitempty"`
	// Access holds the value of the access edge.
	Access []*Access `json:"access,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AdventuresOrErr returns the Adventures value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) AdventuresOrErr() ([]*Adventure, error) {
	if e.loadedTypes[0] {
		return e.Adventures, nil
	}
	return nil, &NotLoadedError{edge: "adventures"}
}

// ProgressOrErr returns the Progress value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) ProgressOrErr() ([]*Progress, error) {
	if e.loadedTypes[1] {
		return e.Progress, nil
	}
	return nil, &NotLoadedError{edge: "progress"}
}

// AccessOrErr returns the Access value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) AccessOrErr() ([]*Access, error) {
	if e.loadedTypes[2] {
		return e.Access, nil
	}
	return nil, &NotLoadedError{edge: "access"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Team) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case team.FieldID:
			values[i] = new(sql.NullInt64)
		case team.FieldName, team.FieldCode, team.FieldEmail, team.FieldStatus:
			values[i] = new(sql.NullString)
		case team.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case team.ForeignKeys[0]: // guess_team
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Team", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Team fields.
func (t *Team) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case team.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case team.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				t.CreateTime = value.Time
			}
		case team.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case team.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				t.Code = value.String
			}
		case team.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				t.Email = value.String
			}
		case team.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = team.Status(value.String)
			}
		case team.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field guess_team", value)
			} else if value.Valid {
				t.guess_team = new(int)
				*t.guess_team = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAdventures queries the "adventures" edge of the Team entity.
func (t *Team) QueryAdventures() *AdventureQuery {
	return (&TeamClient{config: t.config}).QueryAdventures(t)
}

// QueryProgress queries the "progress" edge of the Team entity.
func (t *Team) QueryProgress() *ProgressQuery {
	return (&TeamClient{config: t.config}).QueryProgress(t)
}

// QueryAccess queries the "access" edge of the Team entity.
func (t *Team) QueryAccess() *AccessQuery {
	return (&TeamClient{config: t.config}).QueryAccess(t)
}

// Update returns a builder for updating this Team.
// Note that you need to call Team.Unwrap() before calling this method if this Team
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Team) Update() *TeamUpdateOne {
	return (&TeamClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Team entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Team) Unwrap() *Team {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Team is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Team) String() string {
	var builder strings.Builder
	builder.WriteString("Team(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("create_time=")
	builder.WriteString(t.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(t.Code)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(t.Email)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Teams is a parsable slice of Team.
type Teams []*Team

func (t Teams) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
