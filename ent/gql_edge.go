// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (ae *AccountingEntry) Company(ctx context.Context) (*Company, error) {
	result, err := ae.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = ae.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ae *AccountingEntry) User(ctx context.Context) (*User, error) {
	result, err := ae.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = ae.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (cm *CashMovement) Treasury(ctx context.Context) (*Treasury, error) {
	result, err := cm.Edges.TreasuryOrErr()
	if IsNotLoaded(err) {
		result, err = cm.QueryTreasury().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Company) AvailableRoles(ctx context.Context) (result []*UserRole, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedAvailableRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.AvailableRolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryAvailableRoles().All(ctx)
	}
	return result, err
}

func (c *Company) AccountingEntries(ctx context.Context) (result []*AccountingEntry, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedAccountingEntries(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.AccountingEntriesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryAccountingEntries().All(ctx)
	}
	return result, err
}

func (c *Company) Customers(ctx context.Context) (result []*Customer, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedCustomers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.CustomersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryCustomers().All(ctx)
	}
	return result, err
}

func (c *Company) Employees(ctx context.Context) (result []*Employee, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedEmployees(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.EmployeesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryEmployees().All(ctx)
	}
	return result, err
}

func (c *Company) Files(ctx context.Context) (result []*File, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedFiles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.FilesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryFiles().All(ctx)
	}
	return result, err
}

func (c *Company) Products(ctx context.Context) (result []*Product, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedProducts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.ProductsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryProducts().All(ctx)
	}
	return result, err
}

func (c *Company) Suppliers(ctx context.Context) (result []*Supplier, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedSuppliers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.SuppliersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QuerySuppliers().All(ctx)
	}
	return result, err
}

func (c *Company) Tokens(ctx context.Context) (result []*Token, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedTokens(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.TokensOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryTokens().All(ctx)
	}
	return result, err
}

func (c *Company) Treasuries(ctx context.Context) (result []*Treasury, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedTreasuries(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.TreasuriesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryTreasuries().All(ctx)
	}
	return result, err
}

func (c *Company) WorkShifts(ctx context.Context) (result []*Workshift, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedWorkShifts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.WorkShiftsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryWorkShifts().All(ctx)
	}
	return result, err
}

func (c *Company) WorkTasks(ctx context.Context) (result []*Worktask, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedWorkTasks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.WorkTasksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryWorkTasks().All(ctx)
	}
	return result, err
}

func (c *Company) WorkTags(ctx context.Context) (result []*Worktag, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedWorkTags(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.WorkTagsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryWorkTags().All(ctx)
	}
	return result, err
}

func (c *Company) Users(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.UsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryUsers().All(ctx)
	}
	return result, err
}

func (c *Company) DaughterCompanies(ctx context.Context) (result []*Company, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedDaughterCompanies(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.DaughterCompaniesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryDaughterCompanies().All(ctx)
	}
	return result, err
}

func (c *Company) ParentCompany(ctx context.Context) (*Company, error) {
	result, err := c.Edges.ParentCompanyOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryParentCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Customer) Company(ctx context.Context) (*Company, error) {
	result, err := c.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Customer) Receivables(ctx context.Context) (result []*Receivable, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedReceivables(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.ReceivablesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryReceivables().All(ctx)
	}
	return result, err
}

func (e *Employee) Company(ctx context.Context) (*Company, error) {
	result, err := e.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Employee) User(ctx context.Context) (*User, error) {
	result, err := e.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Employee) Employees(ctx context.Context) (result []*Employee, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = e.NamedEmployees(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = e.Edges.EmployeesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = e.QueryEmployees().All(ctx)
	}
	return result, err
}

func (e *Employee) Supervisor(ctx context.Context) (*Employee, error) {
	result, err := e.Edges.SupervisorOrErr()
	if IsNotLoaded(err) {
		result, err = e.QuerySupervisor().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Employee) WorkShifts(ctx context.Context) (result []*Workshift, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = e.NamedWorkShifts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = e.Edges.WorkShiftsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = e.QueryWorkShifts().All(ctx)
	}
	return result, err
}

func (e *Employee) ApprovedWorkShifts(ctx context.Context) (result []*Workshift, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = e.NamedApprovedWorkShifts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = e.Edges.ApprovedWorkShiftsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = e.QueryApprovedWorkShifts().All(ctx)
	}
	return result, err
}

func (e *Employee) AssignedTasks(ctx context.Context) (result []*Worktask, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = e.NamedAssignedTasks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = e.Edges.AssignedTasksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = e.QueryAssignedTasks().All(ctx)
	}
	return result, err
}

func (f *File) Company(ctx context.Context) (*Company, error) {
	result, err := f.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = f.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (f *File) Product(ctx context.Context) (*Product, error) {
	result, err := f.Edges.ProductOrErr()
	if IsNotLoaded(err) {
		result, err = f.QueryProduct().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pa *Payable) Supplier(ctx context.Context) (*Supplier, error) {
	result, err := pa.Edges.SupplierOrErr()
	if IsNotLoaded(err) {
		result, err = pa.QuerySupplier().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Product) Company(ctx context.Context) (*Company, error) {
	result, err := pr.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = pr.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pr *Product) Pictures(ctx context.Context) (result []*File, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pr.NamedPictures(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pr.Edges.PicturesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pr.QueryPictures().All(ctx)
	}
	return result, err
}

func (pr *Product) ProductMovements(ctx context.Context) (result []*ProductMovement, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pr.NamedProductMovements(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pr.Edges.ProductMovementsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pr.QueryProductMovements().All(ctx)
	}
	return result, err
}

func (pm *ProductMovement) Product(ctx context.Context) (*Product, error) {
	result, err := pm.Edges.ProductOrErr()
	if IsNotLoaded(err) {
		result, err = pm.QueryProduct().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (r *Receivable) Customer(ctx context.Context) (*Customer, error) {
	result, err := r.Edges.CustomerOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryCustomer().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Supplier) Company(ctx context.Context) (*Company, error) {
	result, err := s.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Supplier) Payables(ctx context.Context) (result []*Payable, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = s.NamedPayables(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = s.Edges.PayablesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = s.QueryPayables().All(ctx)
	}
	return result, err
}

func (t *Token) Company(ctx context.Context) (*Company, error) {
	result, err := t.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Token) User(ctx context.Context) (*User, error) {
	result, err := t.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Treasury) Company(ctx context.Context) (*Company, error) {
	result, err := t.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Treasury) CashMovements(ctx context.Context) (result []*CashMovement, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedCashMovements(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.CashMovementsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryCashMovements().All(ctx)
	}
	return result, err
}

func (u *User) AccountingEntries(ctx context.Context) (result []*AccountingEntry, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedAccountingEntries(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.AccountingEntriesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryAccountingEntries().All(ctx)
	}
	return result, err
}

func (u *User) Company(ctx context.Context) (result []*Company, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedCompany(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.CompanyOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryCompany().All(ctx)
	}
	return result, err
}

func (u *User) AssignedRoles(ctx context.Context) (result []*UserRole, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedAssignedRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.AssignedRolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryAssignedRoles().All(ctx)
	}
	return result, err
}

func (u *User) CreatedTasks(ctx context.Context) (result []*Worktask, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedCreatedTasks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.CreatedTasksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryCreatedTasks().All(ctx)
	}
	return result, err
}

func (u *User) Employee(ctx context.Context) (*Employee, error) {
	result, err := u.Edges.EmployeeOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryEmployee().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Tokens(ctx context.Context) (result []*Token, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedTokens(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.TokensOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryTokens().All(ctx)
	}
	return result, err
}

func (ur *UserRole) Company(ctx context.Context) (*Company, error) {
	result, err := ur.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = ur.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ur *UserRole) User(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ur.NamedUser(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ur.Edges.UserOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ur.QueryUser().All(ctx)
	}
	return result, err
}

func (w *Workshift) Company(ctx context.Context) (*Company, error) {
	result, err := w.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Workshift) Employee(ctx context.Context) (*Employee, error) {
	result, err := w.Edges.EmployeeOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryEmployee().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Workshift) ApprovedBy(ctx context.Context) (*Employee, error) {
	result, err := w.Edges.ApprovedByOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryApprovedBy().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Workshift) WorkTask(ctx context.Context) (*Worktask, error) {
	result, err := w.Edges.WorkTaskOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryWorkTask().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Workshift) EditRequest(ctx context.Context) (*Workshift, error) {
	result, err := w.Edges.EditRequestOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryEditRequest().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Workshift) WorkShift(ctx context.Context) (*Workshift, error) {
	result, err := w.Edges.WorkShiftOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryWorkShift().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Worktag) Company(ctx context.Context) (*Company, error) {
	result, err := w.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Worktag) WorkTasks(ctx context.Context) (result []*Worktask, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = w.NamedWorkTasks(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = w.Edges.WorkTasksOrErr()
	}
	if IsNotLoaded(err) {
		result, err = w.QueryWorkTasks().All(ctx)
	}
	return result, err
}

func (w *Worktask) Company(ctx context.Context) (*Company, error) {
	result, err := w.Edges.CompanyOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryCompany().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Worktask) CreatedBy(ctx context.Context) (*User, error) {
	result, err := w.Edges.CreatedByOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryCreatedBy().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (w *Worktask) AssignedTo(ctx context.Context) (result []*Employee, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = w.NamedAssignedTo(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = w.Edges.AssignedToOrErr()
	}
	if IsNotLoaded(err) {
		result, err = w.QueryAssignedTo().All(ctx)
	}
	return result, err
}

func (w *Worktask) WorkShifts(ctx context.Context) (result []*Workshift, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = w.NamedWorkShifts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = w.Edges.WorkShiftsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = w.QueryWorkShifts().All(ctx)
	}
	return result, err
}

func (w *Worktask) WorkTags(ctx context.Context) (result []*Worktag, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = w.NamedWorkTags(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = w.Edges.WorkTagsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = w.QueryWorkTags().All(ctx)
	}
	return result, err
}