// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mazza/ent/company"
	"mazza/ent/supplier"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Supplier is the model entity for the Supplier schema.
type Supplier struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deletedAt" field.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// IsDefault holds the value of the "isDefault" field.
	IsDefault bool `json:"isDefault,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// TaxId holds the value of the "taxId" field.
	TaxId string `json:"taxId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SupplierQuery when eager-loading is set.
	Edges             SupplierEdges `json:"edges"`
	company_suppliers *int
	selectValues      sql.SelectValues
}

// SupplierEdges holds the relations/edges for other nodes in the graph.
type SupplierEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// Payables holds the value of the payables edge.
	Payables []*Payable `json:"payables,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedPayables map[string][]*Payable
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SupplierEdges) CompanyOrErr() (*Company, error) {
	if e.Company != nil {
		return e.Company, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: company.Label}
	}
	return nil, &NotLoadedError{edge: "company"}
}

// PayablesOrErr returns the Payables value or an error if the edge
// was not loaded in eager-loading.
func (e SupplierEdges) PayablesOrErr() ([]*Payable, error) {
	if e.loadedTypes[1] {
		return e.Payables, nil
	}
	return nil, &NotLoadedError{edge: "payables"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Supplier) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case supplier.FieldIsDefault:
			values[i] = new(sql.NullBool)
		case supplier.FieldID:
			values[i] = new(sql.NullInt64)
		case supplier.FieldAddress, supplier.FieldCity, supplier.FieldCountry, supplier.FieldDescription, supplier.FieldEmail, supplier.FieldName, supplier.FieldPhone, supplier.FieldTaxId:
			values[i] = new(sql.NullString)
		case supplier.FieldCreatedAt, supplier.FieldUpdatedAt, supplier.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case supplier.ForeignKeys[0]: // company_suppliers
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Supplier fields.
func (s *Supplier) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case supplier.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case supplier.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case supplier.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case supplier.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedAt", values[i])
			} else if value.Valid {
				s.DeletedAt = new(time.Time)
				*s.DeletedAt = value.Time
			}
		case supplier.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				s.Address = value.String
			}
		case supplier.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				s.City = value.String
			}
		case supplier.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				s.Country = value.String
			}
		case supplier.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case supplier.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				s.Email = value.String
			}
		case supplier.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isDefault", values[i])
			} else if value.Valid {
				s.IsDefault = value.Bool
			}
		case supplier.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case supplier.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				s.Phone = value.String
			}
		case supplier.FieldTaxId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field taxId", values[i])
			} else if value.Valid {
				s.TaxId = value.String
			}
		case supplier.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_suppliers", value)
			} else if value.Valid {
				s.company_suppliers = new(int)
				*s.company_suppliers = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Supplier.
// This includes values selected through modifiers, order, etc.
func (s *Supplier) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryCompany queries the "company" edge of the Supplier entity.
func (s *Supplier) QueryCompany() *CompanyQuery {
	return NewSupplierClient(s.config).QueryCompany(s)
}

// QueryPayables queries the "payables" edge of the Supplier entity.
func (s *Supplier) QueryPayables() *PayableQuery {
	return NewSupplierClient(s.config).QueryPayables(s)
}

// Update returns a builder for updating this Supplier.
// Note that you need to call Supplier.Unwrap() before calling this method if this Supplier
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Supplier) Update() *SupplierUpdateOne {
	return NewSupplierClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Supplier entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Supplier) Unwrap() *Supplier {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Supplier is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Supplier) String() string {
	var builder strings.Builder
	builder.WriteString("Supplier(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := s.DeletedAt; v != nil {
		builder.WriteString("deletedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(s.Address)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(s.City)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(s.Country)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(s.Email)
	builder.WriteString(", ")
	builder.WriteString("isDefault=")
	builder.WriteString(fmt.Sprintf("%v", s.IsDefault))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(s.Phone)
	builder.WriteString(", ")
	builder.WriteString("taxId=")
	builder.WriteString(s.TaxId)
	builder.WriteByte(')')
	return builder.String()
}

// NamedPayables returns the Payables named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Supplier) NamedPayables(name string) ([]*Payable, error) {
	if s.Edges.namedPayables == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedPayables[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Supplier) appendNamedPayables(name string, edges ...*Payable) {
	if s.Edges.namedPayables == nil {
		s.Edges.namedPayables = make(map[string][]*Payable)
	}
	if len(edges) == 0 {
		s.Edges.namedPayables[name] = []*Payable{}
	} else {
		s.Edges.namedPayables[name] = append(s.Edges.namedPayables[name], edges...)
	}
}

// Suppliers is a parsable slice of Supplier.
type Suppliers []*Supplier
