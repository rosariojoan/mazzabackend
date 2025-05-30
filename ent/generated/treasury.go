// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"mazza/ent/generated/company"
	"mazza/ent/generated/treasury"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Treasury is the model entity for the Treasury schema.
type Treasury struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deletedAt" field.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance float64 `json:"balance,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TreasuryQuery when eager-loading is set.
	Edges              TreasuryEdges `json:"edges"`
	company_treasuries *int
	selectValues       sql.SelectValues
}

// TreasuryEdges holds the relations/edges for other nodes in the graph.
type TreasuryEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TreasuryEdges) CompanyOrErr() (*Company, error) {
	if e.Company != nil {
		return e.Company, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: company.Label}
	}
	return nil, &NotLoadedError{edge: "company"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Treasury) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case treasury.FieldBalance:
			values[i] = new(sql.NullFloat64)
		case treasury.FieldID:
			values[i] = new(sql.NullInt64)
		case treasury.FieldCreatedAt, treasury.FieldUpdatedAt, treasury.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case treasury.ForeignKeys[0]: // company_treasuries
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Treasury fields.
func (t *Treasury) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case treasury.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case treasury.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case treasury.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case treasury.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedAt", values[i])
			} else if value.Valid {
				t.DeletedAt = new(time.Time)
				*t.DeletedAt = value.Time
			}
		case treasury.FieldBalance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				t.Balance = value.Float64
			}
		case treasury.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_treasuries", value)
			} else if value.Valid {
				t.company_treasuries = new(int)
				*t.company_treasuries = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Treasury.
// This includes values selected through modifiers, order, etc.
func (t *Treasury) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryCompany queries the "company" edge of the Treasury entity.
func (t *Treasury) QueryCompany() *CompanyQuery {
	return NewTreasuryClient(t.config).QueryCompany(t)
}

// Update returns a builder for updating this Treasury.
// Note that you need to call Treasury.Unwrap() before calling this method if this Treasury
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Treasury) Update() *TreasuryUpdateOne {
	return NewTreasuryClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Treasury entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Treasury) Unwrap() *Treasury {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("generated: Treasury is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Treasury) String() string {
	var builder strings.Builder
	builder.WriteString("Treasury(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := t.DeletedAt; v != nil {
		builder.WriteString("deletedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("balance=")
	builder.WriteString(fmt.Sprintf("%v", t.Balance))
	builder.WriteByte(')')
	return builder.String()
}

// Treasuries is a parsable slice of Treasury.
type Treasuries []*Treasury
