// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"mazza/ent/generated/company"
	"mazza/ent/generated/product"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deletedAt" field.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// IsDefault holds the value of the "isDefault" field.
	IsDefault bool `json:"isDefault,omitempty"`
	// MinimumStock holds the value of the "minimumStock" field.
	MinimumStock int `json:"minimumStock,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// Sku holds the value of the "sku" field.
	Sku string `json:"sku,omitempty"`
	// Stock holds the value of the "stock" field.
	Stock float64 `json:"stock,omitempty"`
	// Category holds the value of the "category" field.
	Category product.Category `json:"category,omitempty"`
	// UnitCost holds the value of the "unitCost" field.
	UnitCost float64 `json:"unitCost,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges            ProductEdges `json:"edges"`
	company_products *int
	selectValues     sql.SelectValues
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Company holds the value of the company edge.
	Company *Company `json:"company,omitempty"`
	// Pictures holds the value of the pictures edge.
	Pictures []*File `json:"pictures,omitempty"`
	// AccountingEntries holds the value of the accountingEntries edge.
	AccountingEntries []*AccountingEntry `json:"accountingEntries,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedPictures          map[string][]*File
	namedAccountingEntries map[string][]*AccountingEntry
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) CompanyOrErr() (*Company, error) {
	if e.Company != nil {
		return e.Company, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: company.Label}
	}
	return nil, &NotLoadedError{edge: "company"}
}

// PicturesOrErr returns the Pictures value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) PicturesOrErr() ([]*File, error) {
	if e.loadedTypes[1] {
		return e.Pictures, nil
	}
	return nil, &NotLoadedError{edge: "pictures"}
}

// AccountingEntriesOrErr returns the AccountingEntries value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) AccountingEntriesOrErr() ([]*AccountingEntry, error) {
	if e.loadedTypes[2] {
		return e.AccountingEntries, nil
	}
	return nil, &NotLoadedError{edge: "accountingEntries"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldIsDefault:
			values[i] = new(sql.NullBool)
		case product.FieldStock, product.FieldUnitCost:
			values[i] = new(sql.NullFloat64)
		case product.FieldID, product.FieldMinimumStock, product.FieldPrice:
			values[i] = new(sql.NullInt64)
		case product.FieldDescription, product.FieldName, product.FieldSku, product.FieldCategory:
			values[i] = new(sql.NullString)
		case product.FieldCreatedAt, product.FieldUpdatedAt, product.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case product.ForeignKeys[0]: // company_products
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case product.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case product.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case product.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedAt", values[i])
			} else if value.Valid {
				pr.DeletedAt = new(time.Time)
				*pr.DeletedAt = value.Time
			}
		case product.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case product.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isDefault", values[i])
			} else if value.Valid {
				pr.IsDefault = value.Bool
			}
		case product.FieldMinimumStock:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field minimumStock", values[i])
			} else if value.Valid {
				pr.MinimumStock = int(value.Int64)
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pr.Price = int(value.Int64)
			}
		case product.FieldSku:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sku", values[i])
			} else if value.Valid {
				pr.Sku = value.String
			}
		case product.FieldStock:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field stock", values[i])
			} else if value.Valid {
				pr.Stock = value.Float64
			}
		case product.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				pr.Category = product.Category(value.String)
			}
		case product.FieldUnitCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field unitCost", values[i])
			} else if value.Valid {
				pr.UnitCost = value.Float64
			}
		case product.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_products", value)
			} else if value.Valid {
				pr.company_products = new(int)
				*pr.company_products = int(value.Int64)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Product.
// This includes values selected through modifiers, order, etc.
func (pr *Product) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryCompany queries the "company" edge of the Product entity.
func (pr *Product) QueryCompany() *CompanyQuery {
	return NewProductClient(pr.config).QueryCompany(pr)
}

// QueryPictures queries the "pictures" edge of the Product entity.
func (pr *Product) QueryPictures() *FileQuery {
	return NewProductClient(pr.config).QueryPictures(pr)
}

// QueryAccountingEntries queries the "accountingEntries" edge of the Product entity.
func (pr *Product) QueryAccountingEntries() *AccountingEntryQuery {
	return NewProductClient(pr.config).QueryAccountingEntries(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return NewProductClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("generated: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := pr.DeletedAt; v != nil {
		builder.WriteString("deletedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("isDefault=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsDefault))
	builder.WriteString(", ")
	builder.WriteString("minimumStock=")
	builder.WriteString(fmt.Sprintf("%v", pr.MinimumStock))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteString(", ")
	builder.WriteString("sku=")
	builder.WriteString(pr.Sku)
	builder.WriteString(", ")
	builder.WriteString("stock=")
	builder.WriteString(fmt.Sprintf("%v", pr.Stock))
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fmt.Sprintf("%v", pr.Category))
	builder.WriteString(", ")
	builder.WriteString("unitCost=")
	builder.WriteString(fmt.Sprintf("%v", pr.UnitCost))
	builder.WriteByte(')')
	return builder.String()
}

// NamedPictures returns the Pictures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedPictures(name string) ([]*File, error) {
	if pr.Edges.namedPictures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedPictures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedPictures(name string, edges ...*File) {
	if pr.Edges.namedPictures == nil {
		pr.Edges.namedPictures = make(map[string][]*File)
	}
	if len(edges) == 0 {
		pr.Edges.namedPictures[name] = []*File{}
	} else {
		pr.Edges.namedPictures[name] = append(pr.Edges.namedPictures[name], edges...)
	}
}

// NamedAccountingEntries returns the AccountingEntries named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Product) NamedAccountingEntries(name string) ([]*AccountingEntry, error) {
	if pr.Edges.namedAccountingEntries == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedAccountingEntries[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Product) appendNamedAccountingEntries(name string, edges ...*AccountingEntry) {
	if pr.Edges.namedAccountingEntries == nil {
		pr.Edges.namedAccountingEntries = make(map[string][]*AccountingEntry)
	}
	if len(edges) == 0 {
		pr.Edges.namedAccountingEntries[name] = []*AccountingEntry{}
	} else {
		pr.Edges.namedAccountingEntries[name] = append(pr.Edges.namedAccountingEntries[name], edges...)
	}
}

// Products is a parsable slice of Product.
type Products []*Product
