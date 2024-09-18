// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mazza/ent/company"
	"mazza/ent/customer"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
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
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges             CustomerEdges `json:"edges"`
	company_customers *int
	selectValues      sql.SelectValues
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// Receivables holds the value of the receivables edge.
	Receivables []*Receivable `json:"receivables,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedReceivables map[string][]*Receivable
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) CompanyOrErr() (*Company, error) {
	if e.Company != nil {
		return e.Company, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: company.Label}
	}
	return nil, &NotLoadedError{edge: "company"}
}

// ReceivablesOrErr returns the Receivables value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) ReceivablesOrErr() ([]*Receivable, error) {
	if e.loadedTypes[1] {
		return e.Receivables, nil
	}
	return nil, &NotLoadedError{edge: "receivables"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldIsDefault:
			values[i] = new(sql.NullBool)
		case customer.FieldID:
			values[i] = new(sql.NullInt64)
		case customer.FieldAddress, customer.FieldCity, customer.FieldCountry, customer.FieldDescription, customer.FieldEmail, customer.FieldName, customer.FieldPhone, customer.FieldTaxId:
			values[i] = new(sql.NullString)
		case customer.FieldCreatedAt, customer.FieldUpdatedAt, customer.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case customer.ForeignKeys[0]: // company_customers
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case customer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case customer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case customer.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedAt", values[i])
			} else if value.Valid {
				c.DeletedAt = new(time.Time)
				*c.DeletedAt = value.Time
			}
		case customer.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				c.Address = value.String
			}
		case customer.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				c.City = value.String
			}
		case customer.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				c.Country = value.String
			}
		case customer.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case customer.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case customer.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isDefault", values[i])
			} else if value.Valid {
				c.IsDefault = value.Bool
			}
		case customer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case customer.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = value.String
			}
		case customer.FieldTaxId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field taxId", values[i])
			} else if value.Valid {
				c.TaxId = value.String
			}
		case customer.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_customers", value)
			} else if value.Valid {
				c.company_customers = new(int)
				*c.company_customers = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Customer.
// This includes values selected through modifiers, order, etc.
func (c *Customer) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryCompany queries the "company" edge of the Customer entity.
func (c *Customer) QueryCompany() *CompanyQuery {
	return NewCustomerClient(c.config).QueryCompany(c)
}

// QueryReceivables queries the "receivables" edge of the Customer entity.
func (c *Customer) QueryReceivables() *ReceivableQuery {
	return NewCustomerClient(c.config).QueryReceivables(c)
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return NewCustomerClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := c.DeletedAt; v != nil {
		builder.WriteString("deletedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(c.Address)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(c.City)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(c.Country)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("isDefault=")
	builder.WriteString(fmt.Sprintf("%v", c.IsDefault))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(c.Phone)
	builder.WriteString(", ")
	builder.WriteString("taxId=")
	builder.WriteString(c.TaxId)
	builder.WriteByte(')')
	return builder.String()
}

// NamedReceivables returns the Receivables named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Customer) NamedReceivables(name string) ([]*Receivable, error) {
	if c.Edges.namedReceivables == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedReceivables[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Customer) appendNamedReceivables(name string, edges ...*Receivable) {
	if c.Edges.namedReceivables == nil {
		c.Edges.namedReceivables = make(map[string][]*Receivable)
	}
	if len(edges) == 0 {
		c.Edges.namedReceivables[name] = []*Receivable{}
	} else {
		c.Edges.namedReceivables[name] = append(c.Edges.namedReceivables[name], edges...)
	}
}

// Customers is a parsable slice of Customer.
type Customers []*Customer
