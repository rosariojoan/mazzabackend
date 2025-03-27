// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"mazza/ent/generated/company"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Company is the model entity for the Company schema.
type Company struct {
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
	Address *string `json:"address,omitempty"`
	// BaseCurrency holds the value of the "baseCurrency" field.
	BaseCurrency string `json:"baseCurrency,omitempty"`
	// CeoName holds the value of the "ceoName" field.
	CeoName *string `json:"ceoName,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// EstablishedAt holds the value of the "establishedAt" field.
	EstablishedAt time.Time `json:"establishedAt,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Email holds the value of the "email" field.
	Email *string `json:"email,omitempty"`
	// Industry holds the value of the "industry" field.
	Industry *string `json:"industry,omitempty"`
	// LastEntryDate holds the value of the "lastEntryDate" field.
	LastEntryDate *time.Time `json:"lastEntryDate,omitempty"`
	// LastInvoiceNumber holds the value of the "lastInvoiceNumber" field.
	LastInvoiceNumber int32 `json:"lastInvoiceNumber,omitempty"`
	// LogoURL holds the value of the "logoURL" field.
	LogoURL *string `json:"logoURL,omitempty"`
	// LogoStorageURI holds the value of the "logoStorageURI" field.
	LogoStorageURI *string `json:"-"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// NumberOfEmployees holds the value of the "numberOfEmployees" field.
	NumberOfEmployees int32 `json:"numberOfEmployees,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone *string `json:"phone,omitempty"`
	// TaxId holds the value of the "taxId" field.
	TaxId *string `json:"taxId,omitempty"`
	// VatRate holds the value of the "vatRate" field.
	VatRate float64 `json:"vatRate,omitempty"`
	// Website holds the value of the "website" field.
	Website *string `json:"website,omitempty"`
	// IncompleteSetup holds the value of the "incompleteSetup" field.
	IncompleteSetup bool `json:"incompleteSetup,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompanyQuery when eager-loading is set.
	Edges                      CompanyEdges `json:"edges"`
	company_daughter_companies *int
	selectValues               sql.SelectValues
}

// CompanyEdges holds the relations/edges for other nodes in the graph.
type CompanyEdges struct {
	// AvailableRoles holds the value of the availableRoles edge.
	AvailableRoles []*UserRole `json:"availableRoles,omitempty"`
	// AccountingEntries holds the value of the accountingEntries edge.
	AccountingEntries []*AccountingEntry `json:"accountingEntries,omitempty"`
	// Customers holds the value of the customers edge.
	Customers []*Customer `json:"customers,omitempty"`
	// Documents holds the value of the documents edge.
	Documents []*CompanyDocument `json:"documents,omitempty"`
	// Employees holds the value of the employees edge.
	Employees []*Employee `json:"employees,omitempty"`
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// MemberSignupTokens holds the value of the memberSignupTokens edge.
	MemberSignupTokens []*MemberSignupToken `json:"memberSignupTokens,omitempty"`
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// Payables holds the value of the payables edge.
	Payables []*Payable `json:"payables,omitempty"`
	// Receivables holds the value of the receivables edge.
	Receivables []*Receivable `json:"receivables,omitempty"`
	// Suppliers holds the value of the suppliers edge.
	Suppliers []*Supplier `json:"suppliers,omitempty"`
	// Tokens holds the value of the tokens edge.
	Tokens []*Token `json:"tokens,omitempty"`
	// Treasuries holds the value of the treasuries edge.
	Treasuries []*Treasury `json:"treasuries,omitempty"`
	// WorkShifts holds the value of the workShifts edge.
	WorkShifts []*Workshift `json:"workShifts,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// DaughterCompanies holds the value of the daughterCompanies edge.
	DaughterCompanies []*Company `json:"daughterCompanies,omitempty"`
	// ParentCompany holds the value of the parentCompany edge.
	ParentCompany *Company `json:"parentCompany,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [18]bool
	// totalCount holds the count of the edges above.
	totalCount [18]map[string]int

	namedAvailableRoles     map[string][]*UserRole
	namedAccountingEntries  map[string][]*AccountingEntry
	namedCustomers          map[string][]*Customer
	namedDocuments          map[string][]*CompanyDocument
	namedEmployees          map[string][]*Employee
	namedFiles              map[string][]*File
	namedMemberSignupTokens map[string][]*MemberSignupToken
	namedProducts           map[string][]*Product
	namedProjects           map[string][]*Project
	namedPayables           map[string][]*Payable
	namedReceivables        map[string][]*Receivable
	namedSuppliers          map[string][]*Supplier
	namedTokens             map[string][]*Token
	namedTreasuries         map[string][]*Treasury
	namedWorkShifts         map[string][]*Workshift
	namedUsers              map[string][]*User
	namedDaughterCompanies  map[string][]*Company
}

// AvailableRolesOrErr returns the AvailableRoles value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) AvailableRolesOrErr() ([]*UserRole, error) {
	if e.loadedTypes[0] {
		return e.AvailableRoles, nil
	}
	return nil, &NotLoadedError{edge: "availableRoles"}
}

// AccountingEntriesOrErr returns the AccountingEntries value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) AccountingEntriesOrErr() ([]*AccountingEntry, error) {
	if e.loadedTypes[1] {
		return e.AccountingEntries, nil
	}
	return nil, &NotLoadedError{edge: "accountingEntries"}
}

// CustomersOrErr returns the Customers value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) CustomersOrErr() ([]*Customer, error) {
	if e.loadedTypes[2] {
		return e.Customers, nil
	}
	return nil, &NotLoadedError{edge: "customers"}
}

// DocumentsOrErr returns the Documents value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) DocumentsOrErr() ([]*CompanyDocument, error) {
	if e.loadedTypes[3] {
		return e.Documents, nil
	}
	return nil, &NotLoadedError{edge: "documents"}
}

// EmployeesOrErr returns the Employees value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) EmployeesOrErr() ([]*Employee, error) {
	if e.loadedTypes[4] {
		return e.Employees, nil
	}
	return nil, &NotLoadedError{edge: "employees"}
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[5] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// MemberSignupTokensOrErr returns the MemberSignupTokens value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) MemberSignupTokensOrErr() ([]*MemberSignupToken, error) {
	if e.loadedTypes[6] {
		return e.MemberSignupTokens, nil
	}
	return nil, &NotLoadedError{edge: "memberSignupTokens"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[7] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[8] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// PayablesOrErr returns the Payables value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) PayablesOrErr() ([]*Payable, error) {
	if e.loadedTypes[9] {
		return e.Payables, nil
	}
	return nil, &NotLoadedError{edge: "payables"}
}

// ReceivablesOrErr returns the Receivables value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) ReceivablesOrErr() ([]*Receivable, error) {
	if e.loadedTypes[10] {
		return e.Receivables, nil
	}
	return nil, &NotLoadedError{edge: "receivables"}
}

// SuppliersOrErr returns the Suppliers value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) SuppliersOrErr() ([]*Supplier, error) {
	if e.loadedTypes[11] {
		return e.Suppliers, nil
	}
	return nil, &NotLoadedError{edge: "suppliers"}
}

// TokensOrErr returns the Tokens value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) TokensOrErr() ([]*Token, error) {
	if e.loadedTypes[12] {
		return e.Tokens, nil
	}
	return nil, &NotLoadedError{edge: "tokens"}
}

// TreasuriesOrErr returns the Treasuries value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) TreasuriesOrErr() ([]*Treasury, error) {
	if e.loadedTypes[13] {
		return e.Treasuries, nil
	}
	return nil, &NotLoadedError{edge: "treasuries"}
}

// WorkShiftsOrErr returns the WorkShifts value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) WorkShiftsOrErr() ([]*Workshift, error) {
	if e.loadedTypes[14] {
		return e.WorkShifts, nil
	}
	return nil, &NotLoadedError{edge: "workShifts"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[15] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// DaughterCompaniesOrErr returns the DaughterCompanies value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) DaughterCompaniesOrErr() ([]*Company, error) {
	if e.loadedTypes[16] {
		return e.DaughterCompanies, nil
	}
	return nil, &NotLoadedError{edge: "daughterCompanies"}
}

// ParentCompanyOrErr returns the ParentCompany value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompanyEdges) ParentCompanyOrErr() (*Company, error) {
	if e.ParentCompany != nil {
		return e.ParentCompany, nil
	} else if e.loadedTypes[17] {
		return nil, &NotFoundError{label: company.Label}
	}
	return nil, &NotLoadedError{edge: "parentCompany"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Company) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case company.FieldIncompleteSetup:
			values[i] = new(sql.NullBool)
		case company.FieldVatRate:
			values[i] = new(sql.NullFloat64)
		case company.FieldID, company.FieldLastInvoiceNumber, company.FieldNumberOfEmployees:
			values[i] = new(sql.NullInt64)
		case company.FieldAddress, company.FieldBaseCurrency, company.FieldCeoName, company.FieldCity, company.FieldCountry, company.FieldDescription, company.FieldEmail, company.FieldIndustry, company.FieldLogoURL, company.FieldLogoStorageURI, company.FieldName, company.FieldPhone, company.FieldTaxId, company.FieldWebsite:
			values[i] = new(sql.NullString)
		case company.FieldCreatedAt, company.FieldUpdatedAt, company.FieldDeletedAt, company.FieldEstablishedAt, company.FieldLastEntryDate:
			values[i] = new(sql.NullTime)
		case company.ForeignKeys[0]: // company_daughter_companies
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Company fields.
func (c *Company) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case company.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case company.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case company.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case company.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedAt", values[i])
			} else if value.Valid {
				c.DeletedAt = new(time.Time)
				*c.DeletedAt = value.Time
			}
		case company.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				c.Address = new(string)
				*c.Address = value.String
			}
		case company.FieldBaseCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field baseCurrency", values[i])
			} else if value.Valid {
				c.BaseCurrency = value.String
			}
		case company.FieldCeoName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ceoName", values[i])
			} else if value.Valid {
				c.CeoName = new(string)
				*c.CeoName = value.String
			}
		case company.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				c.City = value.String
			}
		case company.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				c.Country = value.String
			}
		case company.FieldEstablishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field establishedAt", values[i])
			} else if value.Valid {
				c.EstablishedAt = value.Time
			}
		case company.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = new(string)
				*c.Description = value.String
			}
		case company.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = new(string)
				*c.Email = value.String
			}
		case company.FieldIndustry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field industry", values[i])
			} else if value.Valid {
				c.Industry = new(string)
				*c.Industry = value.String
			}
		case company.FieldLastEntryDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lastEntryDate", values[i])
			} else if value.Valid {
				c.LastEntryDate = new(time.Time)
				*c.LastEntryDate = value.Time
			}
		case company.FieldLastInvoiceNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field lastInvoiceNumber", values[i])
			} else if value.Valid {
				c.LastInvoiceNumber = int32(value.Int64)
			}
		case company.FieldLogoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logoURL", values[i])
			} else if value.Valid {
				c.LogoURL = new(string)
				*c.LogoURL = value.String
			}
		case company.FieldLogoStorageURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field logoStorageURI", values[i])
			} else if value.Valid {
				c.LogoStorageURI = new(string)
				*c.LogoStorageURI = value.String
			}
		case company.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case company.FieldNumberOfEmployees:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field numberOfEmployees", values[i])
			} else if value.Valid {
				c.NumberOfEmployees = int32(value.Int64)
			}
		case company.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				c.Phone = new(string)
				*c.Phone = value.String
			}
		case company.FieldTaxId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field taxId", values[i])
			} else if value.Valid {
				c.TaxId = new(string)
				*c.TaxId = value.String
			}
		case company.FieldVatRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field vatRate", values[i])
			} else if value.Valid {
				c.VatRate = value.Float64
			}
		case company.FieldWebsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website", values[i])
			} else if value.Valid {
				c.Website = new(string)
				*c.Website = value.String
			}
		case company.FieldIncompleteSetup:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field incompleteSetup", values[i])
			} else if value.Valid {
				c.IncompleteSetup = value.Bool
			}
		case company.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field company_daughter_companies", value)
			} else if value.Valid {
				c.company_daughter_companies = new(int)
				*c.company_daughter_companies = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Company.
// This includes values selected through modifiers, order, etc.
func (c *Company) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryAvailableRoles queries the "availableRoles" edge of the Company entity.
func (c *Company) QueryAvailableRoles() *UserRoleQuery {
	return NewCompanyClient(c.config).QueryAvailableRoles(c)
}

// QueryAccountingEntries queries the "accountingEntries" edge of the Company entity.
func (c *Company) QueryAccountingEntries() *AccountingEntryQuery {
	return NewCompanyClient(c.config).QueryAccountingEntries(c)
}

// QueryCustomers queries the "customers" edge of the Company entity.
func (c *Company) QueryCustomers() *CustomerQuery {
	return NewCompanyClient(c.config).QueryCustomers(c)
}

// QueryDocuments queries the "documents" edge of the Company entity.
func (c *Company) QueryDocuments() *CompanyDocumentQuery {
	return NewCompanyClient(c.config).QueryDocuments(c)
}

// QueryEmployees queries the "employees" edge of the Company entity.
func (c *Company) QueryEmployees() *EmployeeQuery {
	return NewCompanyClient(c.config).QueryEmployees(c)
}

// QueryFiles queries the "files" edge of the Company entity.
func (c *Company) QueryFiles() *FileQuery {
	return NewCompanyClient(c.config).QueryFiles(c)
}

// QueryMemberSignupTokens queries the "memberSignupTokens" edge of the Company entity.
func (c *Company) QueryMemberSignupTokens() *MemberSignupTokenQuery {
	return NewCompanyClient(c.config).QueryMemberSignupTokens(c)
}

// QueryProducts queries the "products" edge of the Company entity.
func (c *Company) QueryProducts() *ProductQuery {
	return NewCompanyClient(c.config).QueryProducts(c)
}

// QueryProjects queries the "projects" edge of the Company entity.
func (c *Company) QueryProjects() *ProjectQuery {
	return NewCompanyClient(c.config).QueryProjects(c)
}

// QueryPayables queries the "payables" edge of the Company entity.
func (c *Company) QueryPayables() *PayableQuery {
	return NewCompanyClient(c.config).QueryPayables(c)
}

// QueryReceivables queries the "receivables" edge of the Company entity.
func (c *Company) QueryReceivables() *ReceivableQuery {
	return NewCompanyClient(c.config).QueryReceivables(c)
}

// QuerySuppliers queries the "suppliers" edge of the Company entity.
func (c *Company) QuerySuppliers() *SupplierQuery {
	return NewCompanyClient(c.config).QuerySuppliers(c)
}

// QueryTokens queries the "tokens" edge of the Company entity.
func (c *Company) QueryTokens() *TokenQuery {
	return NewCompanyClient(c.config).QueryTokens(c)
}

// QueryTreasuries queries the "treasuries" edge of the Company entity.
func (c *Company) QueryTreasuries() *TreasuryQuery {
	return NewCompanyClient(c.config).QueryTreasuries(c)
}

// QueryWorkShifts queries the "workShifts" edge of the Company entity.
func (c *Company) QueryWorkShifts() *WorkshiftQuery {
	return NewCompanyClient(c.config).QueryWorkShifts(c)
}

// QueryUsers queries the "users" edge of the Company entity.
func (c *Company) QueryUsers() *UserQuery {
	return NewCompanyClient(c.config).QueryUsers(c)
}

// QueryDaughterCompanies queries the "daughterCompanies" edge of the Company entity.
func (c *Company) QueryDaughterCompanies() *CompanyQuery {
	return NewCompanyClient(c.config).QueryDaughterCompanies(c)
}

// QueryParentCompany queries the "parentCompany" edge of the Company entity.
func (c *Company) QueryParentCompany() *CompanyQuery {
	return NewCompanyClient(c.config).QueryParentCompany(c)
}

// Update returns a builder for updating this Company.
// Note that you need to call Company.Unwrap() before calling this method if this Company
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Company) Update() *CompanyUpdateOne {
	return NewCompanyClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Company entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Company) Unwrap() *Company {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("generated: Company is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Company) String() string {
	var builder strings.Builder
	builder.WriteString("Company(")
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
	if v := c.Address; v != nil {
		builder.WriteString("address=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("baseCurrency=")
	builder.WriteString(c.BaseCurrency)
	builder.WriteString(", ")
	if v := c.CeoName; v != nil {
		builder.WriteString("ceoName=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(c.City)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(c.Country)
	builder.WriteString(", ")
	builder.WriteString("establishedAt=")
	builder.WriteString(c.EstablishedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := c.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.Email; v != nil {
		builder.WriteString("email=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.Industry; v != nil {
		builder.WriteString("industry=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.LastEntryDate; v != nil {
		builder.WriteString("lastEntryDate=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("lastInvoiceNumber=")
	builder.WriteString(fmt.Sprintf("%v", c.LastInvoiceNumber))
	builder.WriteString(", ")
	if v := c.LogoURL; v != nil {
		builder.WriteString("logoURL=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("logoStorageURI=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("numberOfEmployees=")
	builder.WriteString(fmt.Sprintf("%v", c.NumberOfEmployees))
	builder.WriteString(", ")
	if v := c.Phone; v != nil {
		builder.WriteString("phone=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := c.TaxId; v != nil {
		builder.WriteString("taxId=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("vatRate=")
	builder.WriteString(fmt.Sprintf("%v", c.VatRate))
	builder.WriteString(", ")
	if v := c.Website; v != nil {
		builder.WriteString("website=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("incompleteSetup=")
	builder.WriteString(fmt.Sprintf("%v", c.IncompleteSetup))
	builder.WriteByte(')')
	return builder.String()
}

// NamedAvailableRoles returns the AvailableRoles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedAvailableRoles(name string) ([]*UserRole, error) {
	if c.Edges.namedAvailableRoles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedAvailableRoles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedAvailableRoles(name string, edges ...*UserRole) {
	if c.Edges.namedAvailableRoles == nil {
		c.Edges.namedAvailableRoles = make(map[string][]*UserRole)
	}
	if len(edges) == 0 {
		c.Edges.namedAvailableRoles[name] = []*UserRole{}
	} else {
		c.Edges.namedAvailableRoles[name] = append(c.Edges.namedAvailableRoles[name], edges...)
	}
}

// NamedAccountingEntries returns the AccountingEntries named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedAccountingEntries(name string) ([]*AccountingEntry, error) {
	if c.Edges.namedAccountingEntries == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedAccountingEntries[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedAccountingEntries(name string, edges ...*AccountingEntry) {
	if c.Edges.namedAccountingEntries == nil {
		c.Edges.namedAccountingEntries = make(map[string][]*AccountingEntry)
	}
	if len(edges) == 0 {
		c.Edges.namedAccountingEntries[name] = []*AccountingEntry{}
	} else {
		c.Edges.namedAccountingEntries[name] = append(c.Edges.namedAccountingEntries[name], edges...)
	}
}

// NamedCustomers returns the Customers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedCustomers(name string) ([]*Customer, error) {
	if c.Edges.namedCustomers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCustomers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedCustomers(name string, edges ...*Customer) {
	if c.Edges.namedCustomers == nil {
		c.Edges.namedCustomers = make(map[string][]*Customer)
	}
	if len(edges) == 0 {
		c.Edges.namedCustomers[name] = []*Customer{}
	} else {
		c.Edges.namedCustomers[name] = append(c.Edges.namedCustomers[name], edges...)
	}
}

// NamedDocuments returns the Documents named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedDocuments(name string) ([]*CompanyDocument, error) {
	if c.Edges.namedDocuments == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedDocuments[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedDocuments(name string, edges ...*CompanyDocument) {
	if c.Edges.namedDocuments == nil {
		c.Edges.namedDocuments = make(map[string][]*CompanyDocument)
	}
	if len(edges) == 0 {
		c.Edges.namedDocuments[name] = []*CompanyDocument{}
	} else {
		c.Edges.namedDocuments[name] = append(c.Edges.namedDocuments[name], edges...)
	}
}

// NamedEmployees returns the Employees named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedEmployees(name string) ([]*Employee, error) {
	if c.Edges.namedEmployees == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedEmployees[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedEmployees(name string, edges ...*Employee) {
	if c.Edges.namedEmployees == nil {
		c.Edges.namedEmployees = make(map[string][]*Employee)
	}
	if len(edges) == 0 {
		c.Edges.namedEmployees[name] = []*Employee{}
	} else {
		c.Edges.namedEmployees[name] = append(c.Edges.namedEmployees[name], edges...)
	}
}

// NamedFiles returns the Files named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedFiles(name string) ([]*File, error) {
	if c.Edges.namedFiles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedFiles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedFiles(name string, edges ...*File) {
	if c.Edges.namedFiles == nil {
		c.Edges.namedFiles = make(map[string][]*File)
	}
	if len(edges) == 0 {
		c.Edges.namedFiles[name] = []*File{}
	} else {
		c.Edges.namedFiles[name] = append(c.Edges.namedFiles[name], edges...)
	}
}

// NamedMemberSignupTokens returns the MemberSignupTokens named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedMemberSignupTokens(name string) ([]*MemberSignupToken, error) {
	if c.Edges.namedMemberSignupTokens == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedMemberSignupTokens[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedMemberSignupTokens(name string, edges ...*MemberSignupToken) {
	if c.Edges.namedMemberSignupTokens == nil {
		c.Edges.namedMemberSignupTokens = make(map[string][]*MemberSignupToken)
	}
	if len(edges) == 0 {
		c.Edges.namedMemberSignupTokens[name] = []*MemberSignupToken{}
	} else {
		c.Edges.namedMemberSignupTokens[name] = append(c.Edges.namedMemberSignupTokens[name], edges...)
	}
}

// NamedProducts returns the Products named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedProducts(name string) ([]*Product, error) {
	if c.Edges.namedProducts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedProducts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedProducts(name string, edges ...*Product) {
	if c.Edges.namedProducts == nil {
		c.Edges.namedProducts = make(map[string][]*Product)
	}
	if len(edges) == 0 {
		c.Edges.namedProducts[name] = []*Product{}
	} else {
		c.Edges.namedProducts[name] = append(c.Edges.namedProducts[name], edges...)
	}
}

// NamedProjects returns the Projects named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedProjects(name string) ([]*Project, error) {
	if c.Edges.namedProjects == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedProjects[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedProjects(name string, edges ...*Project) {
	if c.Edges.namedProjects == nil {
		c.Edges.namedProjects = make(map[string][]*Project)
	}
	if len(edges) == 0 {
		c.Edges.namedProjects[name] = []*Project{}
	} else {
		c.Edges.namedProjects[name] = append(c.Edges.namedProjects[name], edges...)
	}
}

// NamedPayables returns the Payables named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedPayables(name string) ([]*Payable, error) {
	if c.Edges.namedPayables == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedPayables[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedPayables(name string, edges ...*Payable) {
	if c.Edges.namedPayables == nil {
		c.Edges.namedPayables = make(map[string][]*Payable)
	}
	if len(edges) == 0 {
		c.Edges.namedPayables[name] = []*Payable{}
	} else {
		c.Edges.namedPayables[name] = append(c.Edges.namedPayables[name], edges...)
	}
}

// NamedReceivables returns the Receivables named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedReceivables(name string) ([]*Receivable, error) {
	if c.Edges.namedReceivables == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedReceivables[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedReceivables(name string, edges ...*Receivable) {
	if c.Edges.namedReceivables == nil {
		c.Edges.namedReceivables = make(map[string][]*Receivable)
	}
	if len(edges) == 0 {
		c.Edges.namedReceivables[name] = []*Receivable{}
	} else {
		c.Edges.namedReceivables[name] = append(c.Edges.namedReceivables[name], edges...)
	}
}

// NamedSuppliers returns the Suppliers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedSuppliers(name string) ([]*Supplier, error) {
	if c.Edges.namedSuppliers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedSuppliers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedSuppliers(name string, edges ...*Supplier) {
	if c.Edges.namedSuppliers == nil {
		c.Edges.namedSuppliers = make(map[string][]*Supplier)
	}
	if len(edges) == 0 {
		c.Edges.namedSuppliers[name] = []*Supplier{}
	} else {
		c.Edges.namedSuppliers[name] = append(c.Edges.namedSuppliers[name], edges...)
	}
}

// NamedTokens returns the Tokens named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedTokens(name string) ([]*Token, error) {
	if c.Edges.namedTokens == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedTokens[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedTokens(name string, edges ...*Token) {
	if c.Edges.namedTokens == nil {
		c.Edges.namedTokens = make(map[string][]*Token)
	}
	if len(edges) == 0 {
		c.Edges.namedTokens[name] = []*Token{}
	} else {
		c.Edges.namedTokens[name] = append(c.Edges.namedTokens[name], edges...)
	}
}

// NamedTreasuries returns the Treasuries named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedTreasuries(name string) ([]*Treasury, error) {
	if c.Edges.namedTreasuries == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedTreasuries[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedTreasuries(name string, edges ...*Treasury) {
	if c.Edges.namedTreasuries == nil {
		c.Edges.namedTreasuries = make(map[string][]*Treasury)
	}
	if len(edges) == 0 {
		c.Edges.namedTreasuries[name] = []*Treasury{}
	} else {
		c.Edges.namedTreasuries[name] = append(c.Edges.namedTreasuries[name], edges...)
	}
}

// NamedWorkShifts returns the WorkShifts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedWorkShifts(name string) ([]*Workshift, error) {
	if c.Edges.namedWorkShifts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedWorkShifts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedWorkShifts(name string, edges ...*Workshift) {
	if c.Edges.namedWorkShifts == nil {
		c.Edges.namedWorkShifts = make(map[string][]*Workshift)
	}
	if len(edges) == 0 {
		c.Edges.namedWorkShifts[name] = []*Workshift{}
	} else {
		c.Edges.namedWorkShifts[name] = append(c.Edges.namedWorkShifts[name], edges...)
	}
}

// NamedUsers returns the Users named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedUsers(name string) ([]*User, error) {
	if c.Edges.namedUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedUsers(name string, edges ...*User) {
	if c.Edges.namedUsers == nil {
		c.Edges.namedUsers = make(map[string][]*User)
	}
	if len(edges) == 0 {
		c.Edges.namedUsers[name] = []*User{}
	} else {
		c.Edges.namedUsers[name] = append(c.Edges.namedUsers[name], edges...)
	}
}

// NamedDaughterCompanies returns the DaughterCompanies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Company) NamedDaughterCompanies(name string) ([]*Company, error) {
	if c.Edges.namedDaughterCompanies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedDaughterCompanies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Company) appendNamedDaughterCompanies(name string, edges ...*Company) {
	if c.Edges.namedDaughterCompanies == nil {
		c.Edges.namedDaughterCompanies = make(map[string][]*Company)
	}
	if len(edges) == 0 {
		c.Edges.namedDaughterCompanies[name] = []*Company{}
	} else {
		c.Edges.namedDaughterCompanies[name] = append(c.Edges.namedDaughterCompanies[name], edges...)
	}
}

// Companies is a parsable slice of Company.
type Companies []*Company
