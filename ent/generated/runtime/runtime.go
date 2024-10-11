// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"mazza/ent/generated/accountingentry"
	"mazza/ent/generated/company"
	"mazza/ent/generated/customer"
	"mazza/ent/generated/employee"
	"mazza/ent/generated/file"
	"mazza/ent/generated/payable"
	"mazza/ent/generated/product"
	"mazza/ent/generated/project"
	"mazza/ent/generated/projectmilestone"
	"mazza/ent/generated/projecttask"
	"mazza/ent/generated/receivable"
	"mazza/ent/generated/supplier"
	"mazza/ent/generated/token"
	"mazza/ent/generated/treasury"
	"mazza/ent/generated/user"
	"mazza/ent/generated/workshift"
	"mazza/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountingentryMixin := schema.AccountingEntry{}.Mixin()
	accountingentryMixinFields0 := accountingentryMixin[0].Fields()
	_ = accountingentryMixinFields0
	accountingentryFields := schema.AccountingEntry{}.Fields()
	_ = accountingentryFields
	// accountingentryDescCreatedAt is the schema descriptor for createdAt field.
	accountingentryDescCreatedAt := accountingentryMixinFields0[0].Descriptor()
	// accountingentry.DefaultCreatedAt holds the default value on creation for the createdAt field.
	accountingentry.DefaultCreatedAt = accountingentryDescCreatedAt.Default.(func() time.Time)
	// accountingentryDescUpdatedAt is the schema descriptor for updatedAt field.
	accountingentryDescUpdatedAt := accountingentryMixinFields0[1].Descriptor()
	// accountingentry.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	accountingentry.DefaultUpdatedAt = accountingentryDescUpdatedAt.Default.(func() time.Time)
	// accountingentry.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	accountingentry.UpdateDefaultUpdatedAt = accountingentryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// accountingentryDescNumber is the schema descriptor for number field.
	accountingentryDescNumber := accountingentryFields[0].Descriptor()
	// accountingentry.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	accountingentry.NumberValidator = accountingentryDescNumber.Validators[0].(func(int) error)
	// accountingentryDescGroup is the schema descriptor for group field.
	accountingentryDescGroup := accountingentryFields[1].Descriptor()
	// accountingentry.GroupValidator is a validator for the "group" field. It is called by the builders before save.
	accountingentry.GroupValidator = accountingentryDescGroup.Validators[0].(func(int) error)
	// accountingentryDescDate is the schema descriptor for date field.
	accountingentryDescDate := accountingentryFields[2].Descriptor()
	// accountingentry.DefaultDate holds the default value on creation for the date field.
	accountingentry.DefaultDate = accountingentryDescDate.Default.(func() time.Time)
	// accountingentryDescAccount is the schema descriptor for account field.
	accountingentryDescAccount := accountingentryFields[3].Descriptor()
	// accountingentry.AccountValidator is a validator for the "account" field. It is called by the builders before save.
	accountingentry.AccountValidator = accountingentryDescAccount.Validators[0].(func(string) error)
	// accountingentryDescIsReversal is the schema descriptor for isReversal field.
	accountingentryDescIsReversal := accountingentryFields[9].Descriptor()
	// accountingentry.DefaultIsReversal holds the default value on creation for the isReversal field.
	accountingentry.DefaultIsReversal = accountingentryDescIsReversal.Default.(bool)
	// accountingentryDescReversed is the schema descriptor for reversed field.
	accountingentryDescReversed := accountingentryFields[10].Descriptor()
	// accountingentry.DefaultReversed holds the default value on creation for the reversed field.
	accountingentry.DefaultReversed = accountingentryDescReversed.Default.(bool)
	companyMixin := schema.Company{}.Mixin()
	companyMixinFields0 := companyMixin[0].Fields()
	_ = companyMixinFields0
	companyFields := schema.Company{}.Fields()
	_ = companyFields
	// companyDescCreatedAt is the schema descriptor for createdAt field.
	companyDescCreatedAt := companyMixinFields0[0].Descriptor()
	// company.DefaultCreatedAt holds the default value on creation for the createdAt field.
	company.DefaultCreatedAt = companyDescCreatedAt.Default.(func() time.Time)
	// companyDescUpdatedAt is the schema descriptor for updatedAt field.
	companyDescUpdatedAt := companyMixinFields0[1].Descriptor()
	// company.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	company.DefaultUpdatedAt = companyDescUpdatedAt.Default.(func() time.Time)
	// company.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	company.UpdateDefaultUpdatedAt = companyDescUpdatedAt.UpdateDefault.(func() time.Time)
	// companyDescBaseCurrency is the schema descriptor for baseCurrency field.
	companyDescBaseCurrency := companyFields[1].Descriptor()
	// company.DefaultBaseCurrency holds the default value on creation for the baseCurrency field.
	company.DefaultBaseCurrency = companyDescBaseCurrency.Default.(string)
	// companyDescLastInvoiceNumber is the schema descriptor for lastInvoiceNumber field.
	companyDescLastInvoiceNumber := companyFields[9].Descriptor()
	// company.DefaultLastInvoiceNumber holds the default value on creation for the lastInvoiceNumber field.
	company.DefaultLastInvoiceNumber = companyDescLastInvoiceNumber.Default.(int32)
	// company.LastInvoiceNumberValidator is a validator for the "lastInvoiceNumber" field. It is called by the builders before save.
	company.LastInvoiceNumberValidator = companyDescLastInvoiceNumber.Validators[0].(func(int32) error)
	// companyDescNumberOfEmployees is the schema descriptor for numberOfEmployees field.
	companyDescNumberOfEmployees := companyFields[12].Descriptor()
	// company.DefaultNumberOfEmployees holds the default value on creation for the numberOfEmployees field.
	company.DefaultNumberOfEmployees = companyDescNumberOfEmployees.Default.(int32)
	// company.NumberOfEmployeesValidator is a validator for the "numberOfEmployees" field. It is called by the builders before save.
	company.NumberOfEmployeesValidator = companyDescNumberOfEmployees.Validators[0].(func(int32) error)
	// companyDescVatRate is the schema descriptor for vatRate field.
	companyDescVatRate := companyFields[16].Descriptor()
	// company.DefaultVatRate holds the default value on creation for the vatRate field.
	company.DefaultVatRate = companyDescVatRate.Default.(float64)
	// companyDescIncompleteSetup is the schema descriptor for incompleteSetup field.
	companyDescIncompleteSetup := companyFields[18].Descriptor()
	// company.DefaultIncompleteSetup holds the default value on creation for the incompleteSetup field.
	company.DefaultIncompleteSetup = companyDescIncompleteSetup.Default.(bool)
	customerMixin := schema.Customer{}.Mixin()
	customerMixinFields0 := customerMixin[0].Fields()
	_ = customerMixinFields0
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescCreatedAt is the schema descriptor for createdAt field.
	customerDescCreatedAt := customerMixinFields0[0].Descriptor()
	// customer.DefaultCreatedAt holds the default value on creation for the createdAt field.
	customer.DefaultCreatedAt = customerDescCreatedAt.Default.(func() time.Time)
	// customerDescUpdatedAt is the schema descriptor for updatedAt field.
	customerDescUpdatedAt := customerMixinFields0[1].Descriptor()
	// customer.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	customer.DefaultUpdatedAt = customerDescUpdatedAt.Default.(func() time.Time)
	// customer.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	customer.UpdateDefaultUpdatedAt = customerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// customerDescIsDefault is the schema descriptor for isDefault field.
	customerDescIsDefault := customerFields[5].Descriptor()
	// customer.DefaultIsDefault holds the default value on creation for the isDefault field.
	customer.DefaultIsDefault = customerDescIsDefault.Default.(bool)
	// customerDescName is the schema descriptor for name field.
	customerDescName := customerFields[6].Descriptor()
	// customer.NameValidator is a validator for the "name" field. It is called by the builders before save.
	customer.NameValidator = customerDescName.Validators[0].(func(string) error)
	// customerDescTaxId is the schema descriptor for taxId field.
	customerDescTaxId := customerFields[8].Descriptor()
	// customer.TaxIdValidator is a validator for the "taxId" field. It is called by the builders before save.
	customer.TaxIdValidator = customerDescTaxId.Validators[0].(func(string) error)
	employeeMixin := schema.Employee{}.Mixin()
	employeeMixinFields0 := employeeMixin[0].Fields()
	_ = employeeMixinFields0
	employeeFields := schema.Employee{}.Fields()
	_ = employeeFields
	// employeeDescCreatedAt is the schema descriptor for createdAt field.
	employeeDescCreatedAt := employeeMixinFields0[0].Descriptor()
	// employee.DefaultCreatedAt holds the default value on creation for the createdAt field.
	employee.DefaultCreatedAt = employeeDescCreatedAt.Default.(func() time.Time)
	// employeeDescUpdatedAt is the schema descriptor for updatedAt field.
	employeeDescUpdatedAt := employeeMixinFields0[1].Descriptor()
	// employee.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	employee.DefaultUpdatedAt = employeeDescUpdatedAt.Default.(func() time.Time)
	// employee.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	employee.UpdateDefaultUpdatedAt = employeeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// employeeDescName is the schema descriptor for name field.
	employeeDescName := employeeFields[0].Descriptor()
	// employee.NameValidator is a validator for the "name" field. It is called by the builders before save.
	employee.NameValidator = employeeDescName.Validators[0].(func(string) error)
	fileMixin := schema.File{}.Mixin()
	fileMixinFields0 := fileMixin[0].Fields()
	_ = fileMixinFields0
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescCreatedAt is the schema descriptor for createdAt field.
	fileDescCreatedAt := fileMixinFields0[0].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the createdAt field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	// fileDescUpdatedAt is the schema descriptor for updatedAt field.
	fileDescUpdatedAt := fileMixinFields0[1].Descriptor()
	// file.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	file.DefaultUpdatedAt = fileDescUpdatedAt.Default.(func() time.Time)
	// file.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	file.UpdateDefaultUpdatedAt = fileDescUpdatedAt.UpdateDefault.(func() time.Time)
	// fileDescURI is the schema descriptor for uri field.
	fileDescURI := fileFields[3].Descriptor()
	// file.URIValidator is a validator for the "uri" field. It is called by the builders before save.
	file.URIValidator = fileDescURI.Validators[0].(func(string) error)
	payableMixin := schema.Payable{}.Mixin()
	payableMixinFields0 := payableMixin[0].Fields()
	_ = payableMixinFields0
	payableFields := schema.Payable{}.Fields()
	_ = payableFields
	// payableDescCreatedAt is the schema descriptor for createdAt field.
	payableDescCreatedAt := payableMixinFields0[0].Descriptor()
	// payable.DefaultCreatedAt holds the default value on creation for the createdAt field.
	payable.DefaultCreatedAt = payableDescCreatedAt.Default.(func() time.Time)
	// payableDescUpdatedAt is the schema descriptor for updatedAt field.
	payableDescUpdatedAt := payableMixinFields0[1].Descriptor()
	// payable.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	payable.DefaultUpdatedAt = payableDescUpdatedAt.Default.(func() time.Time)
	// payable.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	payable.UpdateDefaultUpdatedAt = payableDescUpdatedAt.UpdateDefault.(func() time.Time)
	// payableDescEntryGroup is the schema descriptor for entryGroup field.
	payableDescEntryGroup := payableFields[0].Descriptor()
	// payable.EntryGroupValidator is a validator for the "entryGroup" field. It is called by the builders before save.
	payable.EntryGroupValidator = payableDescEntryGroup.Validators[0].(func(int) error)
	// payableDescDaysDue is the schema descriptor for daysDue field.
	payableDescDaysDue := payableFields[4].Descriptor()
	// payable.DaysDueValidator is a validator for the "daysDue" field. It is called by the builders before save.
	payable.DaysDueValidator = payableDescDaysDue.Validators[0].(func(int) error)
	productMixin := schema.Product{}.Mixin()
	productMixinFields0 := productMixin[0].Fields()
	_ = productMixinFields0
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescCreatedAt is the schema descriptor for createdAt field.
	productDescCreatedAt := productMixinFields0[0].Descriptor()
	// product.DefaultCreatedAt holds the default value on creation for the createdAt field.
	product.DefaultCreatedAt = productDescCreatedAt.Default.(func() time.Time)
	// productDescUpdatedAt is the schema descriptor for updatedAt field.
	productDescUpdatedAt := productMixinFields0[1].Descriptor()
	// product.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	product.DefaultUpdatedAt = productDescUpdatedAt.Default.(func() time.Time)
	// product.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	product.UpdateDefaultUpdatedAt = productDescUpdatedAt.UpdateDefault.(func() time.Time)
	// productDescIsDefault is the schema descriptor for isDefault field.
	productDescIsDefault := productFields[1].Descriptor()
	// product.DefaultIsDefault holds the default value on creation for the isDefault field.
	product.DefaultIsDefault = productDescIsDefault.Default.(bool)
	// productDescMinimumStock is the schema descriptor for minimumStock field.
	productDescMinimumStock := productFields[2].Descriptor()
	// product.DefaultMinimumStock holds the default value on creation for the minimumStock field.
	product.DefaultMinimumStock = productDescMinimumStock.Default.(int)
	// product.MinimumStockValidator is a validator for the "minimumStock" field. It is called by the builders before save.
	product.MinimumStockValidator = productDescMinimumStock.Validators[0].(func(int) error)
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[3].Descriptor()
	// product.NameValidator is a validator for the "name" field. It is called by the builders before save.
	product.NameValidator = productDescName.Validators[0].(func(string) error)
	// productDescPrice is the schema descriptor for price field.
	productDescPrice := productFields[4].Descriptor()
	// product.DefaultPrice holds the default value on creation for the price field.
	product.DefaultPrice = productDescPrice.Default.(int)
	// product.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	product.PriceValidator = productDescPrice.Validators[0].(func(int) error)
	// productDescSku is the schema descriptor for sku field.
	productDescSku := productFields[5].Descriptor()
	// product.SkuValidator is a validator for the "sku" field. It is called by the builders before save.
	product.SkuValidator = productDescSku.Validators[0].(func(string) error)
	// productDescStock is the schema descriptor for stock field.
	productDescStock := productFields[6].Descriptor()
	// product.DefaultStock holds the default value on creation for the stock field.
	product.DefaultStock = productDescStock.Default.(float64)
	// product.StockValidator is a validator for the "stock" field. It is called by the builders before save.
	product.StockValidator = productDescStock.Validators[0].(func(float64) error)
	// productDescUnitCost is the schema descriptor for unitCost field.
	productDescUnitCost := productFields[8].Descriptor()
	// product.UnitCostValidator is a validator for the "unitCost" field. It is called by the builders before save.
	product.UnitCostValidator = productDescUnitCost.Validators[0].(func(float64) error)
	projectMixin := schema.Project{}.Mixin()
	projectMixinFields0 := projectMixin[0].Fields()
	_ = projectMixinFields0
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescCreatedAt is the schema descriptor for createdAt field.
	projectDescCreatedAt := projectMixinFields0[0].Descriptor()
	// project.DefaultCreatedAt holds the default value on creation for the createdAt field.
	project.DefaultCreatedAt = projectDescCreatedAt.Default.(func() time.Time)
	// projectDescUpdatedAt is the schema descriptor for updatedAt field.
	projectDescUpdatedAt := projectMixinFields0[1].Descriptor()
	// project.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	project.DefaultUpdatedAt = projectDescUpdatedAt.Default.(func() time.Time)
	// project.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	project.UpdateDefaultUpdatedAt = projectDescUpdatedAt.UpdateDefault.(func() time.Time)
	// projectDescName is the schema descriptor for name field.
	projectDescName := projectFields[0].Descriptor()
	// project.NameValidator is a validator for the "name" field. It is called by the builders before save.
	project.NameValidator = projectDescName.Validators[0].(func(string) error)
	// projectDescDescription is the schema descriptor for description field.
	projectDescDescription := projectFields[1].Descriptor()
	// project.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	project.DescriptionValidator = projectDescDescription.Validators[0].(func(string) error)
	// projectDescProgress is the schema descriptor for progress field.
	projectDescProgress := projectFields[4].Descriptor()
	// project.DefaultProgress holds the default value on creation for the progress field.
	project.DefaultProgress = projectDescProgress.Default.(float64)
	// project.ProgressValidator is a validator for the "progress" field. It is called by the builders before save.
	project.ProgressValidator = func() func(float64) error {
		validators := projectDescProgress.Validators
		fns := [...]func(float64) error{
			validators[0].(func(float64) error),
			validators[1].(func(float64) error),
		}
		return func(progress float64) error {
			for _, fn := range fns {
				if err := fn(progress); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	projectmilestoneFields := schema.ProjectMilestone{}.Fields()
	_ = projectmilestoneFields
	// projectmilestoneDescName is the schema descriptor for name field.
	projectmilestoneDescName := projectmilestoneFields[0].Descriptor()
	// projectmilestone.NameValidator is a validator for the "name" field. It is called by the builders before save.
	projectmilestone.NameValidator = projectmilestoneDescName.Validators[0].(func(string) error)
	projecttaskHooks := schema.ProjectTask{}.Hooks()
	projecttask.Hooks[0] = projecttaskHooks[0]
	projecttaskFields := schema.ProjectTask{}.Fields()
	_ = projecttaskFields
	// projecttaskDescCreatedAt is the schema descriptor for createdAt field.
	projecttaskDescCreatedAt := projecttaskFields[0].Descriptor()
	// projecttask.DefaultCreatedAt holds the default value on creation for the createdAt field.
	projecttask.DefaultCreatedAt = projecttaskDescCreatedAt.Default.(func() time.Time)
	// projecttaskDescName is the schema descriptor for name field.
	projecttaskDescName := projecttaskFields[1].Descriptor()
	// projecttask.NameValidator is a validator for the "name" field. It is called by the builders before save.
	projecttask.NameValidator = projecttaskDescName.Validators[0].(func(string) error)
	// projecttaskDescAssigneeName is the schema descriptor for assigneeName field.
	projecttaskDescAssigneeName := projecttaskFields[2].Descriptor()
	// projecttask.AssigneeNameValidator is a validator for the "assigneeName" field. It is called by the builders before save.
	projecttask.AssigneeNameValidator = projecttaskDescAssigneeName.Validators[0].(func(string) error)
	receivableMixin := schema.Receivable{}.Mixin()
	receivableMixinFields0 := receivableMixin[0].Fields()
	_ = receivableMixinFields0
	receivableFields := schema.Receivable{}.Fields()
	_ = receivableFields
	// receivableDescCreatedAt is the schema descriptor for createdAt field.
	receivableDescCreatedAt := receivableMixinFields0[0].Descriptor()
	// receivable.DefaultCreatedAt holds the default value on creation for the createdAt field.
	receivable.DefaultCreatedAt = receivableDescCreatedAt.Default.(func() time.Time)
	// receivableDescUpdatedAt is the schema descriptor for updatedAt field.
	receivableDescUpdatedAt := receivableMixinFields0[1].Descriptor()
	// receivable.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	receivable.DefaultUpdatedAt = receivableDescUpdatedAt.Default.(func() time.Time)
	// receivable.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	receivable.UpdateDefaultUpdatedAt = receivableDescUpdatedAt.UpdateDefault.(func() time.Time)
	// receivableDescEntryGroup is the schema descriptor for entryGroup field.
	receivableDescEntryGroup := receivableFields[0].Descriptor()
	// receivable.EntryGroupValidator is a validator for the "entryGroup" field. It is called by the builders before save.
	receivable.EntryGroupValidator = receivableDescEntryGroup.Validators[0].(func(int) error)
	// receivableDescDaysDue is the schema descriptor for daysDue field.
	receivableDescDaysDue := receivableFields[4].Descriptor()
	// receivable.DaysDueValidator is a validator for the "daysDue" field. It is called by the builders before save.
	receivable.DaysDueValidator = receivableDescDaysDue.Validators[0].(func(int) error)
	supplierMixin := schema.Supplier{}.Mixin()
	supplierMixinFields0 := supplierMixin[0].Fields()
	_ = supplierMixinFields0
	supplierFields := schema.Supplier{}.Fields()
	_ = supplierFields
	// supplierDescCreatedAt is the schema descriptor for createdAt field.
	supplierDescCreatedAt := supplierMixinFields0[0].Descriptor()
	// supplier.DefaultCreatedAt holds the default value on creation for the createdAt field.
	supplier.DefaultCreatedAt = supplierDescCreatedAt.Default.(func() time.Time)
	// supplierDescUpdatedAt is the schema descriptor for updatedAt field.
	supplierDescUpdatedAt := supplierMixinFields0[1].Descriptor()
	// supplier.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	supplier.DefaultUpdatedAt = supplierDescUpdatedAt.Default.(func() time.Time)
	// supplier.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	supplier.UpdateDefaultUpdatedAt = supplierDescUpdatedAt.UpdateDefault.(func() time.Time)
	// supplierDescIsDefault is the schema descriptor for isDefault field.
	supplierDescIsDefault := supplierFields[5].Descriptor()
	// supplier.DefaultIsDefault holds the default value on creation for the isDefault field.
	supplier.DefaultIsDefault = supplierDescIsDefault.Default.(bool)
	// supplierDescName is the schema descriptor for name field.
	supplierDescName := supplierFields[6].Descriptor()
	// supplier.NameValidator is a validator for the "name" field. It is called by the builders before save.
	supplier.NameValidator = supplierDescName.Validators[0].(func(string) error)
	// supplierDescTaxId is the schema descriptor for taxId field.
	supplierDescTaxId := supplierFields[8].Descriptor()
	// supplier.TaxIdValidator is a validator for the "taxId" field. It is called by the builders before save.
	supplier.TaxIdValidator = supplierDescTaxId.Validators[0].(func(string) error)
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescToken is the schema descriptor for token field.
	tokenDescToken := tokenFields[2].Descriptor()
	// token.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	token.TokenValidator = tokenDescToken.Validators[0].(func(string) error)
	treasuryMixin := schema.Treasury{}.Mixin()
	treasuryMixinFields0 := treasuryMixin[0].Fields()
	_ = treasuryMixinFields0
	treasuryFields := schema.Treasury{}.Fields()
	_ = treasuryFields
	// treasuryDescCreatedAt is the schema descriptor for createdAt field.
	treasuryDescCreatedAt := treasuryMixinFields0[0].Descriptor()
	// treasury.DefaultCreatedAt holds the default value on creation for the createdAt field.
	treasury.DefaultCreatedAt = treasuryDescCreatedAt.Default.(func() time.Time)
	// treasuryDescUpdatedAt is the schema descriptor for updatedAt field.
	treasuryDescUpdatedAt := treasuryMixinFields0[1].Descriptor()
	// treasury.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	treasury.DefaultUpdatedAt = treasuryDescUpdatedAt.Default.(func() time.Time)
	// treasury.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	treasury.UpdateDefaultUpdatedAt = treasuryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// treasuryDescIsDefault is the schema descriptor for isDefault field.
	treasuryDescIsDefault := treasuryFields[6].Descriptor()
	// treasury.DefaultIsDefault holds the default value on creation for the isDefault field.
	treasury.DefaultIsDefault = treasuryDescIsDefault.Default.(bool)
	// treasuryDescIsMainAccount is the schema descriptor for isMainAccount field.
	treasuryDescIsMainAccount := treasuryFields[7].Descriptor()
	// treasury.DefaultIsMainAccount holds the default value on creation for the isMainAccount field.
	treasury.DefaultIsMainAccount = treasuryDescIsMainAccount.Default.(bool)
	// treasuryDescName is the schema descriptor for name field.
	treasuryDescName := treasuryFields[8].Descriptor()
	// treasury.NameValidator is a validator for the "name" field. It is called by the builders before save.
	treasury.NameValidator = treasuryDescName.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for createdAt field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the createdAt field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updatedAt field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescDisabled is the schema descriptor for disabled field.
	userDescDisabled := userFields[5].Descriptor()
	// user.DefaultDisabled holds the default value on creation for the disabled field.
	user.DefaultDisabled = userDescDisabled.Default.(bool)
	// userDescNotVerified is the schema descriptor for notVerified field.
	userDescNotVerified := userFields[6].Descriptor()
	// user.DefaultNotVerified holds the default value on creation for the notVerified field.
	user.DefaultNotVerified = userDescNotVerified.Default.(bool)
	workshiftMixin := schema.Workshift{}.Mixin()
	workshiftHooks := schema.Workshift{}.Hooks()
	workshift.Hooks[0] = workshiftHooks[0]
	workshift.Hooks[1] = workshiftHooks[1]
	workshiftMixinFields0 := workshiftMixin[0].Fields()
	_ = workshiftMixinFields0
	workshiftFields := schema.Workshift{}.Fields()
	_ = workshiftFields
	// workshiftDescCreatedAt is the schema descriptor for createdAt field.
	workshiftDescCreatedAt := workshiftMixinFields0[0].Descriptor()
	// workshift.DefaultCreatedAt holds the default value on creation for the createdAt field.
	workshift.DefaultCreatedAt = workshiftDescCreatedAt.Default.(func() time.Time)
	// workshiftDescUpdatedAt is the schema descriptor for updatedAt field.
	workshiftDescUpdatedAt := workshiftMixinFields0[1].Descriptor()
	// workshift.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	workshift.DefaultUpdatedAt = workshiftDescUpdatedAt.Default.(func() time.Time)
	// workshift.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	workshift.UpdateDefaultUpdatedAt = workshiftDescUpdatedAt.UpdateDefault.(func() time.Time)
	// workshiftDescClockIn is the schema descriptor for clockIn field.
	workshiftDescClockIn := workshiftFields[1].Descriptor()
	// workshift.DefaultClockIn holds the default value on creation for the clockIn field.
	workshift.DefaultClockIn = workshiftDescClockIn.Default.(func() time.Time)
}

const (
	Version = "v0.14.1"                                         // Version of ent codegen.
	Sum     = "h1:fUERL506Pqr92EPHJqr8EYxbPioflJo6PudkrEA8a/s=" // Sum of ent codegen.
)
