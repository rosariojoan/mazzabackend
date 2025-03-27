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
		{Name: "category", Type: field.TypeString, Default: ""},
		{Name: "is_debit", Type: field.TypeBool},
		{Name: "is_reversal", Type: field.TypeBool, Default: false},
		{Name: "reversed", Type: field.TypeBool, Default: false},
		{Name: "company_accounting_entries", Type: field.TypeInt, Nullable: true},
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
				Symbol:     "accounting_entries_users_accountingEntries",
				Columns:    []*schema.Column{AccountingEntriesColumns[17]},
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
		{Name: "industry", Type: field.TypeString, Nullable: true},
		{Name: "last_entry_date", Type: field.TypeTime},
		{Name: "last_invoice_number", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "logo_url", Type: field.TypeString, Nullable: true},
		{Name: "logo_storage_uri", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "number_of_employees", Type: field.TypeInt32, Default: 0},
		{Name: "phone", Type: field.TypeString, Nullable: true},
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
				Columns:    []*schema.Column{CompaniesColumns[24]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "company_country_name",
				Unique:  true,
				Columns: []*schema.Column{CompaniesColumns[8], CompaniesColumns[17]},
			},
		},
	}
	// CompanyDocumentsColumns holds the columns for the "company_documents" table.
	CompanyDocumentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "filename", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "keywords", Type: field.TypeString},
		{Name: "category", Type: field.TypeEnum, Enums: []string{"LEGAL", "CONTRACT", "LICENSE", "TAX", "HR"}},
		{Name: "size", Type: field.TypeInt},
		{Name: "file_type", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PENDING", "APPROVED", "REJECTED", "EXPIRED"}},
		{Name: "url", Type: field.TypeString},
		{Name: "storage_uri", Type: field.TypeString},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true},
		{Name: "expiry_date", Type: field.TypeTime},
		{Name: "company_documents", Type: field.TypeInt},
		{Name: "user_uploaded_documents", Type: field.TypeInt, Nullable: true},
		{Name: "user_approved_documents", Type: field.TypeInt, Nullable: true},
	}
	// CompanyDocumentsTable holds the schema information for the "company_documents" table.
	CompanyDocumentsTable = &schema.Table{
		Name:       "company_documents",
		Columns:    CompanyDocumentsColumns,
		PrimaryKey: []*schema.Column{CompanyDocumentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "company_documents_companies_documents",
				Columns:    []*schema.Column{CompanyDocumentsColumns[15]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "company_documents_users_uploadedDocuments",
				Columns:    []*schema.Column{CompanyDocumentsColumns[16]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "company_documents_users_approvedDocuments",
				Columns:    []*schema.Column{CompanyDocumentsColumns[17]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
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
				Symbol:     "employees_users_employee",
				Columns:    []*schema.Column{EmployeesColumns[10]},
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
		},
	}
	// MemberSignupTokensColumns holds the columns for the "member_signup_tokens" table.
	MemberSignupTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "token", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"SUPERUSER", "ADMIN", "ACCOUNTANT", "AUDITOR", "STAFF"}},
		{Name: "note", Type: field.TypeString},
		{Name: "number_accessed", Type: field.TypeInt, Default: 0},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "already_used", Type: field.TypeBool, Default: false},
		{Name: "company_member_signup_tokens", Type: field.TypeInt},
		{Name: "user_created_member_signup_tokens", Type: field.TypeInt, Nullable: true},
	}
	// MemberSignupTokensTable holds the schema information for the "member_signup_tokens" table.
	MemberSignupTokensTable = &schema.Table{
		Name:       "member_signup_tokens",
		Columns:    MemberSignupTokensColumns,
		PrimaryKey: []*schema.Column{MemberSignupTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "member_signup_tokens_companies_memberSignupTokens",
				Columns:    []*schema.Column{MemberSignupTokensColumns[13]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "member_signup_tokens_users_createdMemberSignupTokens",
				Columns:    []*schema.Column{MemberSignupTokensColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
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
		{Name: "name", Type: field.TypeString, Default: "Diversos"},
		{Name: "outstanding_balance", Type: field.TypeFloat64},
		{Name: "total_transaction", Type: field.TypeFloat64},
		{Name: "due_date", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"paid", "pending", "default"}},
		{Name: "company_payables", Type: field.TypeInt, Nullable: true},
		{Name: "supplier_payables", Type: field.TypeInt, Nullable: true},
	}
	// PayablesTable holds the schema information for the "payables" table.
	PayablesTable = &schema.Table{
		Name:       "payables",
		Columns:    PayablesColumns,
		PrimaryKey: []*schema.Column{PayablesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payables_companies_payables",
				Columns:    []*schema.Column{PayablesColumns[11]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "payables_suppliers_payables",
				Columns:    []*schema.Column{PayablesColumns[12]},
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
		{Name: "stock", Type: field.TypeInt, Default: 0},
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
				Columns:    []*schema.Column{ProductsColumns[5]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
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
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
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
		{Name: "user_created_tasks", Type: field.TypeInt, Nullable: true},
	}
	// ProjectTasksTable holds the schema information for the "project_tasks" table.
	ProjectTasksTable = &schema.Table{
		Name:       "project_tasks",
		Columns:    ProjectTasksColumns,
		PrimaryKey: []*schema.Column{ProjectTasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_tasks_projects_tasks",
				Columns:    []*schema.Column{ProjectTasksColumns[10]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "project_tasks_users_assignedProjectTasks",
				Columns:    []*schema.Column{ProjectTasksColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "project_tasks_users_createdTasks",
				Columns:    []*schema.Column{ProjectTasksColumns[12]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "projecttask_name_project_tasks",
				Unique:  true,
				Columns: []*schema.Column{ProjectTasksColumns[2], ProjectTasksColumns[10]},
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
		{Name: "name", Type: field.TypeString, Default: "Diversos"},
		{Name: "outstanding_balance", Type: field.TypeFloat64},
		{Name: "total_transaction", Type: field.TypeFloat64},
		{Name: "due_date", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"paid", "pending", "default"}},
		{Name: "company_receivables", Type: field.TypeInt, Nullable: true},
		{Name: "customer_receivables", Type: field.TypeInt, Nullable: true},
	}
	// ReceivablesTable holds the schema information for the "receivables" table.
	ReceivablesTable = &schema.Table{
		Name:       "receivables",
		Columns:    ReceivablesColumns,
		PrimaryKey: []*schema.Column{ReceivablesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "receivables_companies_receivables",
				Columns:    []*schema.Column{ReceivablesColumns[11]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "receivables_customers_receivables",
				Columns:    []*schema.Column{ReceivablesColumns[12]},
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
		{Name: "balance", Type: field.TypeFloat64},
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
				Columns:    []*schema.Column{TreasuriesColumns[5]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "firebase_uid", Type: field.TypeString, Unique: true},
		{Name: "fcm_token", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "photo_url", Type: field.TypeString, Nullable: true},
		{Name: "department", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "birthdate", Type: field.TypeTime, Nullable: true},
		{Name: "last_login", Type: field.TypeTime, Nullable: true},
		{Name: "gender", Type: field.TypeEnum, Enums: []string{"male", "female"}},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "user_subordinates", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_users_subordinates",
				Columns:    []*schema.Column{UsersColumns[17]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"SUPERUSER", "ADMIN", "ACCOUNTANT", "AUDITOR", "STAFF"}},
		{Name: "notes", Type: field.TypeString, Size: 255},
		{Name: "company_available_roles", Type: field.TypeInt, Nullable: true},
		{Name: "user_assigned_roles", Type: field.TypeInt, Nullable: true},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_companies_availableRoles",
				Columns:    []*schema.Column{UserRolesColumns[3]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_roles_users_assignedRoles",
				Columns:    []*schema.Column{UserRolesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
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
		{Name: "project_task_work_shifts", Type: field.TypeInt, Nullable: true},
		{Name: "user_approved_work_shifts", Type: field.TypeInt, Nullable: true},
		{Name: "user_work_shifts", Type: field.TypeInt, Nullable: true},
		{Name: "workshift_edit_request", Type: field.TypeInt, Unique: true, Nullable: true},
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
				Symbol:     "workshifts_project_tasks_workShifts",
				Columns:    []*schema.Column{WorkshiftsColumns[13]},
				RefColumns: []*schema.Column{ProjectTasksColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "workshifts_users_approvedWorkShifts",
				Columns:    []*schema.Column{WorkshiftsColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "workshifts_users_workShifts",
				Columns:    []*schema.Column{WorkshiftsColumns[15]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "workshifts_workshifts_editRequest",
				Columns:    []*schema.Column{WorkshiftsColumns[16]},
				RefColumns: []*schema.Column{WorkshiftsColumns[0]},
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
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountingEntriesTable,
		CompaniesTable,
		CompanyDocumentsTable,
		CustomersTable,
		EmployeesTable,
		FilesTable,
		MemberSignupTokensTable,
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
		CompanyUsersTable,
		UserParticipatedProjectTasksTable,
	}
)

func init() {
	AccountingEntriesTable.ForeignKeys[0].RefTable = CompaniesTable
	AccountingEntriesTable.ForeignKeys[1].RefTable = UsersTable
	CompaniesTable.ForeignKeys[0].RefTable = CompaniesTable
	CompanyDocumentsTable.ForeignKeys[0].RefTable = CompaniesTable
	CompanyDocumentsTable.ForeignKeys[1].RefTable = UsersTable
	CompanyDocumentsTable.ForeignKeys[2].RefTable = UsersTable
	CustomersTable.ForeignKeys[0].RefTable = CompaniesTable
	EmployeesTable.ForeignKeys[0].RefTable = CompaniesTable
	EmployeesTable.ForeignKeys[1].RefTable = UsersTable
	FilesTable.ForeignKeys[0].RefTable = CompaniesTable
	MemberSignupTokensTable.ForeignKeys[0].RefTable = CompaniesTable
	MemberSignupTokensTable.ForeignKeys[1].RefTable = UsersTable
	PayablesTable.ForeignKeys[0].RefTable = CompaniesTable
	PayablesTable.ForeignKeys[1].RefTable = SuppliersTable
	ProductsTable.ForeignKeys[0].RefTable = CompaniesTable
	ProjectsTable.ForeignKeys[0].RefTable = CompaniesTable
	ProjectsTable.ForeignKeys[1].RefTable = UsersTable
	ProjectsTable.ForeignKeys[2].RefTable = UsersTable
	ProjectMilestonesTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectTasksTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectTasksTable.ForeignKeys[1].RefTable = UsersTable
	ProjectTasksTable.ForeignKeys[2].RefTable = UsersTable
	ReceivablesTable.ForeignKeys[0].RefTable = CompaniesTable
	ReceivablesTable.ForeignKeys[1].RefTable = CustomersTable
	SuppliersTable.ForeignKeys[0].RefTable = CompaniesTable
	TokensTable.ForeignKeys[0].RefTable = CompaniesTable
	TokensTable.ForeignKeys[1].RefTable = UsersTable
	TreasuriesTable.ForeignKeys[0].RefTable = CompaniesTable
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.ForeignKeys[0].RefTable = CompaniesTable
	UserRolesTable.ForeignKeys[1].RefTable = UsersTable
	WorkshiftsTable.ForeignKeys[0].RefTable = CompaniesTable
	WorkshiftsTable.ForeignKeys[1].RefTable = ProjectTasksTable
	WorkshiftsTable.ForeignKeys[2].RefTable = UsersTable
	WorkshiftsTable.ForeignKeys[3].RefTable = UsersTable
	WorkshiftsTable.ForeignKeys[4].RefTable = WorkshiftsTable
	CompanyUsersTable.ForeignKeys[0].RefTable = CompaniesTable
	CompanyUsersTable.ForeignKeys[1].RefTable = UsersTable
	UserParticipatedProjectTasksTable.ForeignKeys[0].RefTable = UsersTable
	UserParticipatedProjectTasksTable.ForeignKeys[1].RefTable = ProjectTasksTable
	UserParticipatedProjectTasksTable.Annotation = &entsql.Annotation{}
}
