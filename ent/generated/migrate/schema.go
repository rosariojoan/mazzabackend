// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountingEntriesColumns holds the columns for the "accounting_entries" table.
	AccountingEntriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "number", Type: field.TypeInt},
		{Name: "group", Type: field.TypeInt},
		{Name: "date", Type: field.TypeTime},
		{Name: "account", Type: field.TypeString},
		{Name: "label", Type: field.TypeString},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "description", Type: field.TypeString},
		{Name: "account_type", Type: field.TypeEnum, Enums: []string{"ASSET", "LIABILITY", "EQUITY", "REVENUE", "EXPENSE", "TAX_EXPENSE", "INCOME", "DIVIDEND_EXPENSE", "CONTRA_ASSET", "CONTRA_LIABILITY", "CONTRA_REVENUE", "CONTRA_EXPENSE"}},
		{Name: "is_debit", Type: field.TypeBool},
		{Name: "is_reversal", Type: field.TypeBool, Default: false},
		{Name: "reversed", Type: field.TypeBool, Default: false},
		{Name: "quantity", Type: field.TypeInt, Nullable: true},
		{Name: "company_accounting_entries", Type: field.TypeInt, Nullable: true},
		{Name: "product_accounting_entries", Type: field.TypeInt, Nullable: true},
		{Name: "treasury_accounting_entries", Type: field.TypeInt, Nullable: true},
		{Name: "user_accounting_entries", Type: field.TypeInt, Nullable: true},
	}
	// AccountingEntriesTable holds the schema information for the "accounting_entries" table.
	AccountingEntriesTable = &schema.Table{
		Name:       "accounting_entries",
		Columns:    AccountingEntriesColumns,
		PrimaryKey: []*schema.Column{AccountingEntriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounting_entries_companies_accountingEntries",
				Columns:    []*schema.Column{AccountingEntriesColumns[16]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "accounting_entries_products_accountingEntries",
				Columns:    []*schema.Column{AccountingEntriesColumns[17]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "accounting_entries_treasuries_accountingEntries",
				Columns:    []*schema.Column{AccountingEntriesColumns[18]},
				RefColumns: []*schema.Column{TreasuriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "accounting_entries_users_accountingEntries",
				Columns:    []*schema.Column{AccountingEntriesColumns[19]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "accountingentry_number_company_accounting_entries",
				Unique:  true,
				Columns: []*schema.Column{AccountingEntriesColumns[4], AccountingEntriesColumns[16]},
			},
		},
	}
	// CompaniesColumns holds the columns for the "companies" table.
	CompaniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "base_currency", Type: field.TypeString, Default: "mzn"},
		{Name: "ceo_name", Type: field.TypeString, Nullable: true},
		{Name: "city", Type: field.TypeString},
		{Name: "country", Type: field.TypeString},
		{Name: "established_at", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "last_entry_date", Type: field.TypeTime},
		{Name: "last_invoice_number", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "logo", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "number_of_employees", Type: field.TypeInt32, Default: 0},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "sector", Type: field.TypeString, Nullable: true},
		{Name: "tax_id", Type: field.TypeString, Unique: true},
		{Name: "vat_rate", Type: field.TypeFloat64, Default: 0.16},
		{Name: "website", Type: field.TypeString, Nullable: true},
		{Name: "incomplete_setup", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "company_daughter_companies", Type: field.TypeInt, Nullable: true},
	}
	// CompaniesTable holds the schema information for the "companies" table.
	CompaniesTable = &schema.Table{
		Name:       "companies",
		Columns:    CompaniesColumns,
		PrimaryKey: []*schema.Column{CompaniesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "companies_companies_daughterCompanies",
				Columns:    []*schema.Column{CompaniesColumns[23]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "company_country_name",
				Unique:  true,
				Columns: []*schema.Column{CompaniesColumns[8], CompaniesColumns[15]},
			},
		},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "address", Type: field.TypeString},
		{Name: "city", Type: field.TypeString},
		{Name: "country", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "is_default", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "name", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "tax_id", Type: field.TypeString},
		{Name: "company_customers", Type: field.TypeInt, Nullable: true},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "customers_companies_customers",
				Columns:    []*schema.Column{CustomersColumns[13]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "customer_name_company_customers",
				Unique:  true,
				Columns: []*schema.Column{CustomersColumns[10], CustomersColumns[13]},
			},
			{
				Name:    "customer_tax_id_company_customers",
				Unique:  true,
				Columns: []*schema.Column{CustomersColumns[12], CustomersColumns[13]},
			},
		},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "gender", Type: field.TypeEnum, Enums: []string{"male", "female"}},
		{Name: "position", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString},
		{Name: "company_employees", Type: field.TypeInt, Nullable: true},
		{Name: "employee_subordinates", Type: field.TypeInt, Nullable: true},
		{Name: "user_employee", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:       "employees",
		Columns:    EmployeesColumns,
		PrimaryKey: []*schema.Column{EmployeesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "employees_companies_employees",
				Columns:    []*schema.Column{EmployeesColumns[9]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "employees_employees_subordinates",
				Columns:    []*schema.Column{EmployeesColumns[10]},
				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "employees_users_employee",
				Columns:    []*schema.Column{EmployeesColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "category", Type: field.TypeEnum, Enums: []string{"INVOICE", "SALESQUOTATION"}},
		{Name: "extension", Type: field.TypeString},
		{Name: "size", Type: field.TypeString},
		{Name: "uri", Type: field.TypeString},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "company_files", Type: field.TypeInt, Nullable: true},
		{Name: "product_pictures", Type: field.TypeInt, Nullable: true},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "files_companies_files",
				Columns:    []*schema.Column{FilesColumns[10]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "files_products_pictures",
				Columns:    []*schema.Column{FilesColumns[11]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PayablesColumns holds the columns for the "payables" table.
	PayablesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "entry_group", Type: field.TypeInt},
		{Name: "date", Type: field.TypeTime},
		{Name: "outstanding_balance", Type: field.TypeFloat64},
		{Name: "total_transaction", Type: field.TypeFloat64},
		{Name: "days_due", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PAID", "UNPAID", "DOUBTFUL", "DEFAULT"}},
		{Name: "supplier_payables", Type: field.TypeInt, Nullable: true},
	}
	// PayablesTable holds the schema information for the "payables" table.
	PayablesTable = &schema.Table{
		Name:       "payables",
		Columns:    PayablesColumns,
		PrimaryKey: []*schema.Column{PayablesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payables_suppliers_payables",
				Columns:    []*schema.Column{PayablesColumns[10]},
				RefColumns: []*schema.Column{SuppliersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "description", Type: field.TypeString},
		{Name: "is_default", Type: field.TypeBool, Default: false},
		{Name: "minimum_stock", Type: field.TypeInt, Default: 0},
		{Name: "name", Type: field.TypeString},
		{Name: "price", Type: field.TypeInt, Default: 0},
		{Name: "sku", Type: field.TypeString},
		{Name: "stock", Type: field.TypeFloat64, Default: 0},
		{Name: "category", Type: field.TypeEnum, Enums: []string{"MERCHANDISE", "OTHER", "SERVICE"}},
		{Name: "unit_cost", Type: field.TypeFloat64},
		{Name: "company_products", Type: field.TypeInt, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_companies_products",
				Columns:    []*schema.Column{ProductsColumns[13]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "product_name_company_products",
				Unique:  true,
				Columns: []*schema.Column{ProductsColumns[7], ProductsColumns[13]},
			},
			{
				Name:    "product_sku_company_products",
				Unique:  true,
				Columns: []*schema.Column{ProductsColumns[9], ProductsColumns[13]},
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime},
		{Name: "progress", Type: field.TypeFloat64, Default: 0},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"notStarted", "inProgress", "completed"}, Default: "notStarted"},
		{Name: "company_projects", Type: field.TypeInt, Nullable: true},
		{Name: "user_created_projects", Type: field.TypeInt, Nullable: true},
		{Name: "user_leadered_projects", Type: field.TypeInt, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "projects_companies_projects",
				Columns:    []*schema.Column{ProjectsColumns[10]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "projects_users_createdProjects",
				Columns:    []*schema.Column{ProjectsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "projects_users_leaderedProjects",
				Columns:    []*schema.Column{ProjectsColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "project_name_company_projects",
				Unique:  true,
				Columns: []*schema.Column{ProjectsColumns[4], ProjectsColumns[10]},
			},
		},
	}
	// ProjectMilestonesColumns holds the columns for the "project_milestones" table.
	ProjectMilestonesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "due_date", Type: field.TypeTime},
		{Name: "project_milestones", Type: field.TypeInt},
	}
	// ProjectMilestonesTable holds the schema information for the "project_milestones" table.
	ProjectMilestonesTable = &schema.Table{
		Name:       "project_milestones",
		Columns:    ProjectMilestonesColumns,
		PrimaryKey: []*schema.Column{ProjectMilestonesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_milestones_projects_milestones",
				Columns:    []*schema.Column{ProjectMilestonesColumns[3]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "projectmilestone_name_project_milestones",
				Unique:  true,
				Columns: []*schema.Column{ProjectMilestonesColumns[1], ProjectMilestonesColumns[3]},
			},
		},
	}
	// ProjectTasksColumns holds the columns for the "project_tasks" table.
	ProjectTasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "assignee_name", Type: field.TypeString},
		{Name: "location", Type: field.TypeString, Nullable: true},
		{Name: "due_date", Type: field.TypeTime},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"notStarted", "inProgress", "completed"}},
		{Name: "project_tasks", Type: field.TypeInt},
		{Name: "user_assigned_project_tasks", Type: field.TypeInt, Nullable: true},
	}
	// ProjectTasksTable holds the schema information for the "project_tasks" table.
	ProjectTasksTable = &schema.Table{
		Name:       "project_tasks",
		Columns:    ProjectTasksColumns,
		PrimaryKey: []*schema.Column{ProjectTasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_tasks_projects_tasks",
				Columns:    []*schema.Column{ProjectTasksColumns[9]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "project_tasks_users_assignedProjectTasks",
				Columns:    []*schema.Column{ProjectTasksColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "projecttask_name_project_tasks",
				Unique:  true,
				Columns: []*schema.Column{ProjectTasksColumns[1], ProjectTasksColumns[9]},
			},
		},
	}
	// ReceivablesColumns holds the columns for the "receivables" table.
	ReceivablesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "entry_group", Type: field.TypeInt},
		{Name: "date", Type: field.TypeTime},
		{Name: "outstanding_balance", Type: field.TypeFloat64},
		{Name: "total_transaction", Type: field.TypeFloat64},
		{Name: "days_due", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"paid", "unpaid", "doubtful", "default"}},
		{Name: "customer_receivables", Type: field.TypeInt, Nullable: true},
	}
	// ReceivablesTable holds the schema information for the "receivables" table.
	ReceivablesTable = &schema.Table{
		Name:       "receivables",
		Columns:    ReceivablesColumns,
		PrimaryKey: []*schema.Column{ReceivablesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "receivables_customers_receivables",
				Columns:    []*schema.Column{ReceivablesColumns[10]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SuppliersColumns holds the columns for the "suppliers" table.
	SuppliersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "address", Type: field.TypeString},
		{Name: "city", Type: field.TypeString},
		{Name: "country", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "is_default", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "name", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "tax_id", Type: field.TypeString},
		{Name: "company_suppliers", Type: field.TypeInt, Nullable: true},
	}
	// SuppliersTable holds the schema information for the "suppliers" table.
	SuppliersTable = &schema.Table{
		Name:       "suppliers",
		Columns:    SuppliersColumns,
		PrimaryKey: []*schema.Column{SuppliersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "suppliers_companies_suppliers",
				Columns:    []*schema.Column{SuppliersColumns[13]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "supplier_name_company_suppliers",
				Unique:  true,
				Columns: []*schema.Column{SuppliersColumns[10], SuppliersColumns[13]},
			},
			{
				Name:    "supplier_tax_id_company_suppliers",
				Unique:  true,
				Columns: []*schema.Column{SuppliersColumns[12], SuppliersColumns[13]},
			},
		},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "expiry", Type: field.TypeTime},
		{Name: "category", Type: field.TypeEnum, Enums: []string{"passwordReset", "invitation"}},
		{Name: "token", Type: field.TypeString},
		{Name: "company_tokens", Type: field.TypeInt, Nullable: true},
		{Name: "user_tokens", Type: field.TypeInt, Nullable: true},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tokens_companies_tokens",
				Columns:    []*schema.Column{TokensColumns[4]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tokens_users_tokens",
				Columns:    []*schema.Column{TokensColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TreasuriesColumns holds the columns for the "treasuries" table.
	TreasuriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "account_number", Type: field.TypeString, Nullable: true},
		{Name: "balance", Type: field.TypeFloat64},
		{Name: "bank_name", Type: field.TypeString, Nullable: true},
		{Name: "currency", Type: field.TypeEnum, Enums: []string{"mzn"}},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "iban", Type: field.TypeString, Nullable: true},
		{Name: "is_default", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "is_main_account", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "name", Type: field.TypeString},
		{Name: "category", Type: field.TypeEnum, Enums: []string{"DEPOSIT", "CASH"}},
		{Name: "swift_code", Type: field.TypeString, Nullable: true},
		{Name: "company_treasuries", Type: field.TypeInt, Nullable: true},
	}
	// TreasuriesTable holds the schema information for the "treasuries" table.
	TreasuriesTable = &schema.Table{
		Name:       "treasuries",
		Columns:    TreasuriesColumns,
		PrimaryKey: []*schema.Column{TreasuriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "treasuries_companies_treasuries",
				Columns:    []*schema.Column{TreasuriesColumns[15]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "treasury_name_company_treasuries",
				Unique:  true,
				Columns: []*schema.Column{TreasuriesColumns[12], TreasuriesColumns[15]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "fcm_token", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "not_verified", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"superuser", "admin", "accountant", "auditor", "staff"}},
		{Name: "company_available_roles", Type: field.TypeInt, Nullable: true},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_companies_availableRoles",
				Columns:    []*schema.Column{UserRolesColumns[2]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// WorkshiftsColumns holds the columns for the "workshifts" table.
	WorkshiftsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "approved_at", Type: field.TypeTime, Nullable: true},
		{Name: "clock_in", Type: field.TypeTime},
		{Name: "clock_out", Type: field.TypeTime, Nullable: true},
		{Name: "clock_in_location", Type: field.TypeString},
		{Name: "clock_out_location", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"APPROVED", "PENDING"}, Default: "PENDING"},
		{Name: "company_work_shifts", Type: field.TypeInt, Nullable: true},
		{Name: "employee_work_shifts", Type: field.TypeInt, Nullable: true},
		{Name: "employee_approved_work_shifts", Type: field.TypeInt, Nullable: true},
		{Name: "workshift_edit_request", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "worktask_work_shifts", Type: field.TypeInt, Nullable: true},
	}
	// WorkshiftsTable holds the schema information for the "workshifts" table.
	WorkshiftsTable = &schema.Table{
		Name:       "workshifts",
		Columns:    WorkshiftsColumns,
		PrimaryKey: []*schema.Column{WorkshiftsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workshifts_companies_workShifts",
				Columns:    []*schema.Column{WorkshiftsColumns[12]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "workshifts_employees_workShifts",
				Columns:    []*schema.Column{WorkshiftsColumns[13]},
				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "workshifts_employees_approvedWorkShifts",
				Columns:    []*schema.Column{WorkshiftsColumns[14]},
				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "workshifts_workshifts_editRequest",
				Columns:    []*schema.Column{WorkshiftsColumns[15]},
				RefColumns: []*schema.Column{WorkshiftsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "workshifts_worktasks_workShifts",
				Columns:    []*schema.Column{WorkshiftsColumns[16]},
				RefColumns: []*schema.Column{WorktasksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WorktagsColumns holds the columns for the "worktags" table.
	WorktagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "color", Type: field.TypeString},
		{Name: "company_work_tags", Type: field.TypeInt, Nullable: true},
	}
	// WorktagsTable holds the schema information for the "worktags" table.
	WorktagsTable = &schema.Table{
		Name:       "worktags",
		Columns:    WorktagsColumns,
		PrimaryKey: []*schema.Column{WorktagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "worktags_companies_workTags",
				Columns:    []*schema.Column{WorktagsColumns[6]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "worktag_name_company_work_tags",
				Unique:  true,
				Columns: []*schema.Column{WorktagsColumns[4], WorktagsColumns[6]},
			},
		},
	}
	// WorktasksColumns holds the columns for the "worktasks" table.
	WorktasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"DRAFT", "ASSIGNED", "DONE"}},
		{Name: "subtasks", Type: field.TypeJSON, Nullable: true},
		{Name: "title", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime, Nullable: true},
		{Name: "company_work_tasks", Type: field.TypeInt, Nullable: true},
		{Name: "user_created_tasks", Type: field.TypeInt, Nullable: true},
	}
	// WorktasksTable holds the schema information for the "worktasks" table.
	WorktasksTable = &schema.Table{
		Name:       "worktasks",
		Columns:    WorktasksColumns,
		PrimaryKey: []*schema.Column{WorktasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "worktasks_companies_workTasks",
				Columns:    []*schema.Column{WorktasksColumns[10]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "worktasks_users_createdTasks",
				Columns:    []*schema.Column{WorktasksColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CompanyUsersColumns holds the columns for the "company_users" table.
	CompanyUsersColumns = []*schema.Column{
		{Name: "company_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// CompanyUsersTable holds the schema information for the "company_users" table.
	CompanyUsersTable = &schema.Table{
		Name:       "company_users",
		Columns:    CompanyUsersColumns,
		PrimaryKey: []*schema.Column{CompanyUsersColumns[0], CompanyUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "company_users_company_id",
				Columns:    []*schema.Column{CompanyUsersColumns[0]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "company_users_user_id",
				Columns:    []*schema.Column{CompanyUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// EmployeeAssignedTasksColumns holds the columns for the "employee_assignedTasks" table.
	EmployeeAssignedTasksColumns = []*schema.Column{
		{Name: "employee_id", Type: field.TypeInt},
		{Name: "worktask_id", Type: field.TypeInt},
	}
	// EmployeeAssignedTasksTable holds the schema information for the "employee_assignedTasks" table.
	EmployeeAssignedTasksTable = &schema.Table{
		Name:       "employee_assignedTasks",
		Columns:    EmployeeAssignedTasksColumns,
		PrimaryKey: []*schema.Column{EmployeeAssignedTasksColumns[0], EmployeeAssignedTasksColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "employee_assignedTasks_employee_id",
				Columns:    []*schema.Column{EmployeeAssignedTasksColumns[0]},
				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "employee_assignedTasks_worktask_id",
				Columns:    []*schema.Column{EmployeeAssignedTasksColumns[1]},
				RefColumns: []*schema.Column{WorktasksColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserAssignedRolesColumns holds the columns for the "user_assignedRoles" table.
	UserAssignedRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "user_role_id", Type: field.TypeInt},
	}
	// UserAssignedRolesTable holds the schema information for the "user_assignedRoles" table.
	UserAssignedRolesTable = &schema.Table{
		Name:       "user_assignedRoles",
		Columns:    UserAssignedRolesColumns,
		PrimaryKey: []*schema.Column{UserAssignedRolesColumns[0], UserAssignedRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_assignedRoles_user_id",
				Columns:    []*schema.Column{UserAssignedRolesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_assignedRoles_user_role_id",
				Columns:    []*schema.Column{UserAssignedRolesColumns[1]},
				RefColumns: []*schema.Column{UserRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserParticipatedProjectTasksColumns holds the columns for the "user_participatedProjectTasks" table.
	UserParticipatedProjectTasksColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "project_task_id", Type: field.TypeInt},
	}
	// UserParticipatedProjectTasksTable holds the schema information for the "user_participatedProjectTasks" table.
	UserParticipatedProjectTasksTable = &schema.Table{
		Name:       "user_participatedProjectTasks",
		Columns:    UserParticipatedProjectTasksColumns,
		PrimaryKey: []*schema.Column{UserParticipatedProjectTasksColumns[0], UserParticipatedProjectTasksColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_participatedProjectTasks_user_id",
				Columns:    []*schema.Column{UserParticipatedProjectTasksColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_participatedProjectTasks_project_task_id",
				Columns:    []*schema.Column{UserParticipatedProjectTasksColumns[1]},
				RefColumns: []*schema.Column{ProjectTasksColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// WorktaskWorkTagsColumns holds the columns for the "worktask_workTags" table.
	WorktaskWorkTagsColumns = []*schema.Column{
		{Name: "worktask_id", Type: field.TypeInt},
		{Name: "worktag_id", Type: field.TypeInt},
	}
	// WorktaskWorkTagsTable holds the schema information for the "worktask_workTags" table.
	WorktaskWorkTagsTable = &schema.Table{
		Name:       "worktask_workTags",
		Columns:    WorktaskWorkTagsColumns,
		PrimaryKey: []*schema.Column{WorktaskWorkTagsColumns[0], WorktaskWorkTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "worktask_workTags_worktask_id",
				Columns:    []*schema.Column{WorktaskWorkTagsColumns[0]},
				RefColumns: []*schema.Column{WorktasksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "worktask_workTags_worktag_id",
				Columns:    []*schema.Column{WorktaskWorkTagsColumns[1]},
				RefColumns: []*schema.Column{WorktagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountingEntriesTable,
		CompaniesTable,
		CustomersTable,
		EmployeesTable,
		FilesTable,
		PayablesTable,
		ProductsTable,
		ProjectsTable,
		ProjectMilestonesTable,
		ProjectTasksTable,
		ReceivablesTable,
		SuppliersTable,
		TokensTable,
		TreasuriesTable,
		UsersTable,
		UserRolesTable,
		WorkshiftsTable,
		WorktagsTable,
		WorktasksTable,
		CompanyUsersTable,
		EmployeeAssignedTasksTable,
		UserAssignedRolesTable,
		UserParticipatedProjectTasksTable,
		WorktaskWorkTagsTable,
	}
)

func init() {
	AccountingEntriesTable.ForeignKeys[0].RefTable = CompaniesTable
	AccountingEntriesTable.ForeignKeys[1].RefTable = ProductsTable
	AccountingEntriesTable.ForeignKeys[2].RefTable = TreasuriesTable
	AccountingEntriesTable.ForeignKeys[3].RefTable = UsersTable
	CompaniesTable.ForeignKeys[0].RefTable = CompaniesTable
	CustomersTable.ForeignKeys[0].RefTable = CompaniesTable
	EmployeesTable.ForeignKeys[0].RefTable = CompaniesTable
	EmployeesTable.ForeignKeys[1].RefTable = EmployeesTable
	EmployeesTable.ForeignKeys[2].RefTable = UsersTable
	FilesTable.ForeignKeys[0].RefTable = CompaniesTable
	FilesTable.ForeignKeys[1].RefTable = ProductsTable
	PayablesTable.ForeignKeys[0].RefTable = SuppliersTable
	ProductsTable.ForeignKeys[0].RefTable = CompaniesTable
	ProjectsTable.ForeignKeys[0].RefTable = CompaniesTable
	ProjectsTable.ForeignKeys[1].RefTable = UsersTable
	ProjectsTable.ForeignKeys[2].RefTable = UsersTable
	ProjectMilestonesTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectTasksTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectTasksTable.ForeignKeys[1].RefTable = UsersTable
	ReceivablesTable.ForeignKeys[0].RefTable = CustomersTable
	SuppliersTable.ForeignKeys[0].RefTable = CompaniesTable
	TokensTable.ForeignKeys[0].RefTable = CompaniesTable
	TokensTable.ForeignKeys[1].RefTable = UsersTable
	TreasuriesTable.ForeignKeys[0].RefTable = CompaniesTable
	UserRolesTable.ForeignKeys[0].RefTable = CompaniesTable
	WorkshiftsTable.ForeignKeys[0].RefTable = CompaniesTable
	WorkshiftsTable.ForeignKeys[1].RefTable = EmployeesTable
	WorkshiftsTable.ForeignKeys[2].RefTable = EmployeesTable
	WorkshiftsTable.ForeignKeys[3].RefTable = WorkshiftsTable
	WorkshiftsTable.ForeignKeys[4].RefTable = WorktasksTable
	WorktagsTable.ForeignKeys[0].RefTable = CompaniesTable
	WorktasksTable.ForeignKeys[0].RefTable = CompaniesTable
	WorktasksTable.ForeignKeys[1].RefTable = UsersTable
	CompanyUsersTable.ForeignKeys[0].RefTable = CompaniesTable
	CompanyUsersTable.ForeignKeys[1].RefTable = UsersTable
	EmployeeAssignedTasksTable.ForeignKeys[0].RefTable = EmployeesTable
	EmployeeAssignedTasksTable.ForeignKeys[1].RefTable = WorktasksTable
	EmployeeAssignedTasksTable.Annotation = &entsql.Annotation{}
	UserAssignedRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserAssignedRolesTable.ForeignKeys[1].RefTable = UserRolesTable
	UserAssignedRolesTable.Annotation = &entsql.Annotation{}
	UserParticipatedProjectTasksTable.ForeignKeys[0].RefTable = UsersTable
	UserParticipatedProjectTasksTable.ForeignKeys[1].RefTable = ProjectTasksTable
	UserParticipatedProjectTasksTable.Annotation = &entsql.Annotation{}
	WorktaskWorkTagsTable.ForeignKeys[0].RefTable = WorktasksTable
	WorktaskWorkTagsTable.ForeignKeys[1].RefTable = WorktagsTable
	WorktaskWorkTagsTable.Annotation = &entsql.Annotation{}
}
