// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Teeth/app/ent/mealplan"
	"github.com/Teeth/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Mealplan is the model entity for the Mealplan schema.
type Mealplan struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MealplanName holds the value of the "mealplan_name" field.
	MealplanName string `json:"mealplan_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MealplanQuery when eager-loading is set.
	Edges    MealplanEdges `json:"edges"`
	owner_id *int
}

// MealplanEdges holds the relations/edges for other nodes in the graph.
type MealplanEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User
	// Eatinghistory holds the value of the eatinghistory edge.
	Eatinghistory []*Eatinghistory
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MealplanEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// EatinghistoryOrErr returns the Eatinghistory value or an error if the edge
// was not loaded in eager-loading.
func (e MealplanEdges) EatinghistoryOrErr() ([]*Eatinghistory, error) {
	if e.loadedTypes[1] {
		return e.Eatinghistory, nil
	}
	return nil, &NotLoadedError{edge: "eatinghistory"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mealplan) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // mealplan_name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Mealplan) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // owner_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mealplan fields.
func (m *Mealplan) assignValues(values ...interface{}) error {
	if m, n := len(values), len(mealplan.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field mealplan_name", values[0])
	} else if value.Valid {
		m.MealplanName = value.String
	}
	values = values[1:]
	if len(values) == len(mealplan.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field owner_id", value)
		} else if value.Valid {
			m.owner_id = new(int)
			*m.owner_id = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Mealplan.
func (m *Mealplan) QueryOwner() *UserQuery {
	return (&MealplanClient{config: m.config}).QueryOwner(m)
}

// QueryEatinghistory queries the eatinghistory edge of the Mealplan.
func (m *Mealplan) QueryEatinghistory() *EatinghistoryQuery {
	return (&MealplanClient{config: m.config}).QueryEatinghistory(m)
}

// Update returns a builder for updating this Mealplan.
// Note that, you need to call Mealplan.Unwrap() before calling this method, if this Mealplan
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mealplan) Update() *MealplanUpdateOne {
	return (&MealplanClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Mealplan) Unwrap() *Mealplan {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Mealplan is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mealplan) String() string {
	var builder strings.Builder
	builder.WriteString("Mealplan(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", mealplan_name=")
	builder.WriteString(m.MealplanName)
	builder.WriteByte(')')
	return builder.String()
}

// Mealplans is a parsable slice of Mealplan.
type Mealplans []*Mealplan

func (m Mealplans) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
