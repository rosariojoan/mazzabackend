// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"mazza/ent/generated/company"
	"mazza/ent/generated/customer"
	"mazza/ent/generated/invoice"
	"mazza/ent/generated/receivable"
	"mazza/ent/generated/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InvoiceCreate is the builder for creating a Invoice entity.
type InvoiceCreate struct {
	config
	mutation *InvoiceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (ic *InvoiceCreate) SetCreatedAt(t time.Time) *InvoiceCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCreatedAt(t *time.Time) *InvoiceCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updatedAt" field.
func (ic *InvoiceCreate) SetUpdatedAt(t time.Time) *InvoiceCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableUpdatedAt(t *time.Time) *InvoiceCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetDeletedAt sets the "deletedAt" field.
func (ic *InvoiceCreate) SetDeletedAt(t time.Time) *InvoiceCreate {
	ic.mutation.SetDeletedAt(t)
	return ic
}

// SetNillableDeletedAt sets the "deletedAt" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableDeletedAt(t *time.Time) *InvoiceCreate {
	if t != nil {
		ic.SetDeletedAt(*t)
	}
	return ic
}

// SetCompanyLogo sets the "companyLogo" field.
func (ic *InvoiceCreate) SetCompanyLogo(s string) *InvoiceCreate {
	ic.mutation.SetCompanyLogo(s)
	return ic
}

// SetNillableCompanyLogo sets the "companyLogo" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCompanyLogo(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetCompanyLogo(*s)
	}
	return ic
}

// SetCompanyName sets the "companyName" field.
func (ic *InvoiceCreate) SetCompanyName(s string) *InvoiceCreate {
	ic.mutation.SetCompanyName(s)
	return ic
}

// SetCompanyTaxID sets the "companyTaxID" field.
func (ic *InvoiceCreate) SetCompanyTaxID(s string) *InvoiceCreate {
	ic.mutation.SetCompanyTaxID(s)
	return ic
}

// SetNillableCompanyTaxID sets the "companyTaxID" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCompanyTaxID(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetCompanyTaxID(*s)
	}
	return ic
}

// SetCompanyAddress sets the "companyAddress" field.
func (ic *InvoiceCreate) SetCompanyAddress(s string) *InvoiceCreate {
	ic.mutation.SetCompanyAddress(s)
	return ic
}

// SetCompanyCity sets the "companyCity" field.
func (ic *InvoiceCreate) SetCompanyCity(s string) *InvoiceCreate {
	ic.mutation.SetCompanyCity(s)
	return ic
}

// SetCompanyEmail sets the "companyEmail" field.
func (ic *InvoiceCreate) SetCompanyEmail(s string) *InvoiceCreate {
	ic.mutation.SetCompanyEmail(s)
	return ic
}

// SetNillableCompanyEmail sets the "companyEmail" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCompanyEmail(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetCompanyEmail(*s)
	}
	return ic
}

// SetCompanyPhone sets the "companyPhone" field.
func (ic *InvoiceCreate) SetCompanyPhone(s string) *InvoiceCreate {
	ic.mutation.SetCompanyPhone(s)
	return ic
}

// SetNillableCompanyPhone sets the "companyPhone" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCompanyPhone(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetCompanyPhone(*s)
	}
	return ic
}

// SetNumber sets the "number" field.
func (ic *InvoiceCreate) SetNumber(s string) *InvoiceCreate {
	ic.mutation.SetNumber(s)
	return ic
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableNumber(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetNumber(*s)
	}
	return ic
}

// SetIssueDate sets the "issueDate" field.
func (ic *InvoiceCreate) SetIssueDate(t time.Time) *InvoiceCreate {
	ic.mutation.SetIssueDate(t)
	return ic
}

// SetDueDate sets the "dueDate" field.
func (ic *InvoiceCreate) SetDueDate(t time.Time) *InvoiceCreate {
	ic.mutation.SetDueDate(t)
	return ic
}

// SetPaidAt sets the "paidAt" field.
func (ic *InvoiceCreate) SetPaidAt(t time.Time) *InvoiceCreate {
	ic.mutation.SetPaidAt(t)
	return ic
}

// SetNillablePaidAt sets the "paidAt" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillablePaidAt(t *time.Time) *InvoiceCreate {
	if t != nil {
		ic.SetPaidAt(*t)
	}
	return ic
}

// SetStatus sets the "status" field.
func (ic *InvoiceCreate) SetStatus(i invoice.Status) *InvoiceCreate {
	ic.mutation.SetStatus(i)
	return ic
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableStatus(i *invoice.Status) *InvoiceCreate {
	if i != nil {
		ic.SetStatus(*i)
	}
	return ic
}

// SetCustomerName sets the "customerName" field.
func (ic *InvoiceCreate) SetCustomerName(s string) *InvoiceCreate {
	ic.mutation.SetCustomerName(s)
	return ic
}

// SetNillableCustomerName sets the "customerName" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCustomerName(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetCustomerName(*s)
	}
	return ic
}

// SetCustomerTaxID sets the "customerTaxID" field.
func (ic *InvoiceCreate) SetCustomerTaxID(s string) *InvoiceCreate {
	ic.mutation.SetCustomerTaxID(s)
	return ic
}

// SetNillableCustomerTaxID sets the "customerTaxID" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCustomerTaxID(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetCustomerTaxID(*s)
	}
	return ic
}

// SetCustomerAddress sets the "customerAddress" field.
func (ic *InvoiceCreate) SetCustomerAddress(s string) *InvoiceCreate {
	ic.mutation.SetCustomerAddress(s)
	return ic
}

// SetNillableCustomerAddress sets the "customerAddress" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCustomerAddress(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetCustomerAddress(*s)
	}
	return ic
}

// SetCustomerCity sets the "customerCity" field.
func (ic *InvoiceCreate) SetCustomerCity(s string) *InvoiceCreate {
	ic.mutation.SetCustomerCity(s)
	return ic
}

// SetNillableCustomerCity sets the "customerCity" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCustomerCity(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetCustomerCity(*s)
	}
	return ic
}

// SetCustomerEmail sets the "customerEmail" field.
func (ic *InvoiceCreate) SetCustomerEmail(s string) *InvoiceCreate {
	ic.mutation.SetCustomerEmail(s)
	return ic
}

// SetNillableCustomerEmail sets the "customerEmail" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCustomerEmail(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetCustomerEmail(*s)
	}
	return ic
}

// SetCustomerPhone sets the "customerPhone" field.
func (ic *InvoiceCreate) SetCustomerPhone(s string) *InvoiceCreate {
	ic.mutation.SetCustomerPhone(s)
	return ic
}

// SetNillableCustomerPhone sets the "customerPhone" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableCustomerPhone(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetCustomerPhone(*s)
	}
	return ic
}

// SetItems sets the "items" field.
func (ic *InvoiceCreate) SetItems(s string) *InvoiceCreate {
	ic.mutation.SetItems(s)
	return ic
}

// SetSubtotal sets the "subtotal" field.
func (ic *InvoiceCreate) SetSubtotal(f float64) *InvoiceCreate {
	ic.mutation.SetSubtotal(f)
	return ic
}

// SetTax sets the "tax" field.
func (ic *InvoiceCreate) SetTax(f float64) *InvoiceCreate {
	ic.mutation.SetTax(f)
	return ic
}

// SetTotal sets the "total" field.
func (ic *InvoiceCreate) SetTotal(f float64) *InvoiceCreate {
	ic.mutation.SetTotal(f)
	return ic
}

// SetNotes sets the "notes" field.
func (ic *InvoiceCreate) SetNotes(s string) *InvoiceCreate {
	ic.mutation.SetNotes(s)
	return ic
}

// SetNillableNotes sets the "notes" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableNotes(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetNotes(*s)
	}
	return ic
}

// SetPaymentMethod sets the "paymentMethod" field.
func (ic *InvoiceCreate) SetPaymentMethod(s string) *InvoiceCreate {
	ic.mutation.SetPaymentMethod(s)
	return ic
}

// SetNillablePaymentMethod sets the "paymentMethod" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillablePaymentMethod(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetPaymentMethod(*s)
	}
	return ic
}

// SetBankName sets the "bankName" field.
func (ic *InvoiceCreate) SetBankName(s string) *InvoiceCreate {
	ic.mutation.SetBankName(s)
	return ic
}

// SetNillableBankName sets the "bankName" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableBankName(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetBankName(*s)
	}
	return ic
}

// SetBankAgency sets the "bankAgency" field.
func (ic *InvoiceCreate) SetBankAgency(s string) *InvoiceCreate {
	ic.mutation.SetBankAgency(s)
	return ic
}

// SetNillableBankAgency sets the "bankAgency" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableBankAgency(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetBankAgency(*s)
	}
	return ic
}

// SetBankAccountNumber sets the "bankAccountNumber" field.
func (ic *InvoiceCreate) SetBankAccountNumber(s string) *InvoiceCreate {
	ic.mutation.SetBankAccountNumber(s)
	return ic
}

// SetNillableBankAccountNumber sets the "bankAccountNumber" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableBankAccountNumber(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetBankAccountNumber(*s)
	}
	return ic
}

// SetBankAccountName sets the "bankAccountName" field.
func (ic *InvoiceCreate) SetBankAccountName(s string) *InvoiceCreate {
	ic.mutation.SetBankAccountName(s)
	return ic
}

// SetNillableBankAccountName sets the "bankAccountName" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableBankAccountName(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetBankAccountName(*s)
	}
	return ic
}

// SetStorageURI sets the "storageURI" field.
func (ic *InvoiceCreate) SetStorageURI(s string) *InvoiceCreate {
	ic.mutation.SetStorageURI(s)
	return ic
}

// SetNillableStorageURI sets the "storageURI" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableStorageURI(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetStorageURI(*s)
	}
	return ic
}

// SetURL sets the "URL" field.
func (ic *InvoiceCreate) SetURL(s string) *InvoiceCreate {
	ic.mutation.SetURL(s)
	return ic
}

// SetNillableURL sets the "URL" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableURL(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetURL(*s)
	}
	return ic
}

// SetFilename sets the "filename" field.
func (ic *InvoiceCreate) SetFilename(s string) *InvoiceCreate {
	ic.mutation.SetFilename(s)
	return ic
}

// SetNillableFilename sets the "filename" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableFilename(s *string) *InvoiceCreate {
	if s != nil {
		ic.SetFilename(*s)
	}
	return ic
}

// SetSize sets the "size" field.
func (ic *InvoiceCreate) SetSize(f float64) *InvoiceCreate {
	ic.mutation.SetSize(f)
	return ic
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (ic *InvoiceCreate) SetNillableSize(f *float64) *InvoiceCreate {
	if f != nil {
		ic.SetSize(*f)
	}
	return ic
}

// SetKeywords sets the "keywords" field.
func (ic *InvoiceCreate) SetKeywords(s string) *InvoiceCreate {
	ic.mutation.SetKeywords(s)
	return ic
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (ic *InvoiceCreate) SetCompanyID(id int) *InvoiceCreate {
	ic.mutation.SetCompanyID(id)
	return ic
}

// SetCompany sets the "company" edge to the Company entity.
func (ic *InvoiceCreate) SetCompany(c *Company) *InvoiceCreate {
	return ic.SetCompanyID(c.ID)
}

// SetIssuedByID sets the "issuedBy" edge to the User entity by ID.
func (ic *InvoiceCreate) SetIssuedByID(id int) *InvoiceCreate {
	ic.mutation.SetIssuedByID(id)
	return ic
}

// SetNillableIssuedByID sets the "issuedBy" edge to the User entity by ID if the given value is not nil.
func (ic *InvoiceCreate) SetNillableIssuedByID(id *int) *InvoiceCreate {
	if id != nil {
		ic = ic.SetIssuedByID(*id)
	}
	return ic
}

// SetIssuedBy sets the "issuedBy" edge to the User entity.
func (ic *InvoiceCreate) SetIssuedBy(u *User) *InvoiceCreate {
	return ic.SetIssuedByID(u.ID)
}

// SetClientID sets the "client" edge to the Customer entity by ID.
func (ic *InvoiceCreate) SetClientID(id int) *InvoiceCreate {
	ic.mutation.SetClientID(id)
	return ic
}

// SetNillableClientID sets the "client" edge to the Customer entity by ID if the given value is not nil.
func (ic *InvoiceCreate) SetNillableClientID(id *int) *InvoiceCreate {
	if id != nil {
		ic = ic.SetClientID(*id)
	}
	return ic
}

// SetClient sets the "client" edge to the Customer entity.
func (ic *InvoiceCreate) SetClient(c *Customer) *InvoiceCreate {
	return ic.SetClientID(c.ID)
}

// SetReceivableID sets the "receivable" edge to the Receivable entity by ID.
func (ic *InvoiceCreate) SetReceivableID(id int) *InvoiceCreate {
	ic.mutation.SetReceivableID(id)
	return ic
}

// SetNillableReceivableID sets the "receivable" edge to the Receivable entity by ID if the given value is not nil.
func (ic *InvoiceCreate) SetNillableReceivableID(id *int) *InvoiceCreate {
	if id != nil {
		ic = ic.SetReceivableID(*id)
	}
	return ic
}

// SetReceivable sets the "receivable" edge to the Receivable entity.
func (ic *InvoiceCreate) SetReceivable(r *Receivable) *InvoiceCreate {
	return ic.SetReceivableID(r.ID)
}

// Mutation returns the InvoiceMutation object of the builder.
func (ic *InvoiceCreate) Mutation() *InvoiceMutation {
	return ic.mutation
}

// Save creates the Invoice in the database.
func (ic *InvoiceCreate) Save(ctx context.Context) (*Invoice, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InvoiceCreate) SaveX(ctx context.Context) *Invoice {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InvoiceCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InvoiceCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InvoiceCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := invoice.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := invoice.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.Status(); !ok {
		v := invoice.DefaultStatus
		ic.mutation.SetStatus(v)
	}
	if _, ok := ic.mutation.CustomerName(); !ok {
		v := invoice.DefaultCustomerName
		ic.mutation.SetCustomerName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InvoiceCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`generated: missing required field "Invoice.createdAt"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`generated: missing required field "Invoice.updatedAt"`)}
	}
	if _, ok := ic.mutation.CompanyName(); !ok {
		return &ValidationError{Name: "companyName", err: errors.New(`generated: missing required field "Invoice.companyName"`)}
	}
	if _, ok := ic.mutation.CompanyAddress(); !ok {
		return &ValidationError{Name: "companyAddress", err: errors.New(`generated: missing required field "Invoice.companyAddress"`)}
	}
	if _, ok := ic.mutation.CompanyCity(); !ok {
		return &ValidationError{Name: "companyCity", err: errors.New(`generated: missing required field "Invoice.companyCity"`)}
	}
	if _, ok := ic.mutation.IssueDate(); !ok {
		return &ValidationError{Name: "issueDate", err: errors.New(`generated: missing required field "Invoice.issueDate"`)}
	}
	if _, ok := ic.mutation.DueDate(); !ok {
		return &ValidationError{Name: "dueDate", err: errors.New(`generated: missing required field "Invoice.dueDate"`)}
	}
	if _, ok := ic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`generated: missing required field "Invoice.status"`)}
	}
	if v, ok := ic.mutation.Status(); ok {
		if err := invoice.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Invoice.status": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Items(); !ok {
		return &ValidationError{Name: "items", err: errors.New(`generated: missing required field "Invoice.items"`)}
	}
	if v, ok := ic.mutation.Items(); ok {
		if err := invoice.ItemsValidator(v); err != nil {
			return &ValidationError{Name: "items", err: fmt.Errorf(`generated: validator failed for field "Invoice.items": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Subtotal(); !ok {
		return &ValidationError{Name: "subtotal", err: errors.New(`generated: missing required field "Invoice.subtotal"`)}
	}
	if v, ok := ic.mutation.Subtotal(); ok {
		if err := invoice.SubtotalValidator(v); err != nil {
			return &ValidationError{Name: "subtotal", err: fmt.Errorf(`generated: validator failed for field "Invoice.subtotal": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Tax(); !ok {
		return &ValidationError{Name: "tax", err: errors.New(`generated: missing required field "Invoice.tax"`)}
	}
	if v, ok := ic.mutation.Tax(); ok {
		if err := invoice.TaxValidator(v); err != nil {
			return &ValidationError{Name: "tax", err: fmt.Errorf(`generated: validator failed for field "Invoice.tax": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Total(); !ok {
		return &ValidationError{Name: "total", err: errors.New(`generated: missing required field "Invoice.total"`)}
	}
	if v, ok := ic.mutation.Total(); ok {
		if err := invoice.TotalValidator(v); err != nil {
			return &ValidationError{Name: "total", err: fmt.Errorf(`generated: validator failed for field "Invoice.total": %w`, err)}
		}
	}
	if v, ok := ic.mutation.Size(); ok {
		if err := invoice.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`generated: validator failed for field "Invoice.size": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Keywords(); !ok {
		return &ValidationError{Name: "keywords", err: errors.New(`generated: missing required field "Invoice.keywords"`)}
	}
	if v, ok := ic.mutation.Keywords(); ok {
		if err := invoice.KeywordsValidator(v); err != nil {
			return &ValidationError{Name: "keywords", err: fmt.Errorf(`generated: validator failed for field "Invoice.keywords": %w`, err)}
		}
	}
	if len(ic.mutation.CompanyIDs()) == 0 {
		return &ValidationError{Name: "company", err: errors.New(`generated: missing required edge "Invoice.company"`)}
	}
	return nil
}

func (ic *InvoiceCreate) sqlSave(ctx context.Context) (*Invoice, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *InvoiceCreate) createSpec() (*Invoice, *sqlgraph.CreateSpec) {
	var (
		_node = &Invoice{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(invoice.Table, sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(invoice.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(invoice.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.DeletedAt(); ok {
		_spec.SetField(invoice.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := ic.mutation.CompanyLogo(); ok {
		_spec.SetField(invoice.FieldCompanyLogo, field.TypeString, value)
		_node.CompanyLogo = &value
	}
	if value, ok := ic.mutation.CompanyName(); ok {
		_spec.SetField(invoice.FieldCompanyName, field.TypeString, value)
		_node.CompanyName = value
	}
	if value, ok := ic.mutation.CompanyTaxID(); ok {
		_spec.SetField(invoice.FieldCompanyTaxID, field.TypeString, value)
		_node.CompanyTaxID = &value
	}
	if value, ok := ic.mutation.CompanyAddress(); ok {
		_spec.SetField(invoice.FieldCompanyAddress, field.TypeString, value)
		_node.CompanyAddress = value
	}
	if value, ok := ic.mutation.CompanyCity(); ok {
		_spec.SetField(invoice.FieldCompanyCity, field.TypeString, value)
		_node.CompanyCity = value
	}
	if value, ok := ic.mutation.CompanyEmail(); ok {
		_spec.SetField(invoice.FieldCompanyEmail, field.TypeString, value)
		_node.CompanyEmail = &value
	}
	if value, ok := ic.mutation.CompanyPhone(); ok {
		_spec.SetField(invoice.FieldCompanyPhone, field.TypeString, value)
		_node.CompanyPhone = &value
	}
	if value, ok := ic.mutation.Number(); ok {
		_spec.SetField(invoice.FieldNumber, field.TypeString, value)
		_node.Number = &value
	}
	if value, ok := ic.mutation.IssueDate(); ok {
		_spec.SetField(invoice.FieldIssueDate, field.TypeTime, value)
		_node.IssueDate = value
	}
	if value, ok := ic.mutation.DueDate(); ok {
		_spec.SetField(invoice.FieldDueDate, field.TypeTime, value)
		_node.DueDate = value
	}
	if value, ok := ic.mutation.PaidAt(); ok {
		_spec.SetField(invoice.FieldPaidAt, field.TypeTime, value)
		_node.PaidAt = &value
	}
	if value, ok := ic.mutation.Status(); ok {
		_spec.SetField(invoice.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := ic.mutation.CustomerName(); ok {
		_spec.SetField(invoice.FieldCustomerName, field.TypeString, value)
		_node.CustomerName = value
	}
	if value, ok := ic.mutation.CustomerTaxID(); ok {
		_spec.SetField(invoice.FieldCustomerTaxID, field.TypeString, value)
		_node.CustomerTaxID = &value
	}
	if value, ok := ic.mutation.CustomerAddress(); ok {
		_spec.SetField(invoice.FieldCustomerAddress, field.TypeString, value)
		_node.CustomerAddress = &value
	}
	if value, ok := ic.mutation.CustomerCity(); ok {
		_spec.SetField(invoice.FieldCustomerCity, field.TypeString, value)
		_node.CustomerCity = &value
	}
	if value, ok := ic.mutation.CustomerEmail(); ok {
		_spec.SetField(invoice.FieldCustomerEmail, field.TypeString, value)
		_node.CustomerEmail = &value
	}
	if value, ok := ic.mutation.CustomerPhone(); ok {
		_spec.SetField(invoice.FieldCustomerPhone, field.TypeString, value)
		_node.CustomerPhone = &value
	}
	if value, ok := ic.mutation.Items(); ok {
		_spec.SetField(invoice.FieldItems, field.TypeString, value)
		_node.Items = value
	}
	if value, ok := ic.mutation.Subtotal(); ok {
		_spec.SetField(invoice.FieldSubtotal, field.TypeFloat64, value)
		_node.Subtotal = value
	}
	if value, ok := ic.mutation.Tax(); ok {
		_spec.SetField(invoice.FieldTax, field.TypeFloat64, value)
		_node.Tax = value
	}
	if value, ok := ic.mutation.Total(); ok {
		_spec.SetField(invoice.FieldTotal, field.TypeFloat64, value)
		_node.Total = value
	}
	if value, ok := ic.mutation.Notes(); ok {
		_spec.SetField(invoice.FieldNotes, field.TypeString, value)
		_node.Notes = &value
	}
	if value, ok := ic.mutation.PaymentMethod(); ok {
		_spec.SetField(invoice.FieldPaymentMethod, field.TypeString, value)
		_node.PaymentMethod = &value
	}
	if value, ok := ic.mutation.BankName(); ok {
		_spec.SetField(invoice.FieldBankName, field.TypeString, value)
		_node.BankName = &value
	}
	if value, ok := ic.mutation.BankAgency(); ok {
		_spec.SetField(invoice.FieldBankAgency, field.TypeString, value)
		_node.BankAgency = &value
	}
	if value, ok := ic.mutation.BankAccountNumber(); ok {
		_spec.SetField(invoice.FieldBankAccountNumber, field.TypeString, value)
		_node.BankAccountNumber = &value
	}
	if value, ok := ic.mutation.BankAccountName(); ok {
		_spec.SetField(invoice.FieldBankAccountName, field.TypeString, value)
		_node.BankAccountName = &value
	}
	if value, ok := ic.mutation.StorageURI(); ok {
		_spec.SetField(invoice.FieldStorageURI, field.TypeString, value)
		_node.StorageURI = &value
	}
	if value, ok := ic.mutation.URL(); ok {
		_spec.SetField(invoice.FieldURL, field.TypeString, value)
		_node.URL = &value
	}
	if value, ok := ic.mutation.Filename(); ok {
		_spec.SetField(invoice.FieldFilename, field.TypeString, value)
		_node.Filename = &value
	}
	if value, ok := ic.mutation.Size(); ok {
		_spec.SetField(invoice.FieldSize, field.TypeFloat64, value)
		_node.Size = &value
	}
	if value, ok := ic.mutation.Keywords(); ok {
		_spec.SetField(invoice.FieldKeywords, field.TypeString, value)
		_node.Keywords = value
	}
	if nodes := ic.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invoice.CompanyTable,
			Columns: []string{invoice.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_invoices = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.IssuedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invoice.IssuedByTable,
			Columns: []string{invoice.IssuedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_issued_invoices = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ClientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invoice.ClientTable,
			Columns: []string{invoice.ClientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.customer_invoices = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ReceivableIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   invoice.ReceivableTable,
			Columns: []string{invoice.ReceivableColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(receivable.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InvoiceCreateBulk is the builder for creating many Invoice entities in bulk.
type InvoiceCreateBulk struct {
	config
	err      error
	builders []*InvoiceCreate
}

// Save creates the Invoice entities in the database.
func (icb *InvoiceCreateBulk) Save(ctx context.Context) ([]*Invoice, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Invoice, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvoiceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InvoiceCreateBulk) SaveX(ctx context.Context) []*Invoice {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InvoiceCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InvoiceCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
